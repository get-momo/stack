package stripe

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/formancehq/payments/cmd/connectors/internal/connectors/currency"
	"github.com/formancehq/payments/cmd/connectors/internal/connectors/stripe/client"
	"github.com/formancehq/payments/cmd/connectors/internal/ingestion"
	"github.com/formancehq/payments/cmd/connectors/internal/metrics"
	"github.com/formancehq/payments/cmd/connectors/internal/storage"
	"github.com/formancehq/payments/cmd/connectors/internal/task"
	"github.com/formancehq/payments/internal/models"
	"github.com/formancehq/stack/libs/go-libs/contextutil"
	"github.com/formancehq/stack/libs/go-libs/logging"
	"github.com/stripe/stripe-go/v72"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

var (
	initiateTransferAttrs = metric.WithAttributes(append(connectorAttrs, attribute.String(metrics.ObjectAttributeKey, "initiate_transfer"))...)
	initiatePayoutAttrs   = metric.WithAttributes(append(connectorAttrs, attribute.String(metrics.ObjectAttributeKey, "initiate_payout"))...)
)

func initiatePaymentTask(transferID string, stripeClient *client.DefaultClient) task.Task {
	return func(
		ctx context.Context,
		logger logging.Logger,
		connectorID models.ConnectorID,
		ingester ingestion.Ingester,
		scheduler task.Scheduler,
		storageReader storage.Reader,
		metricsRegistry metrics.MetricsRegistry,
	) error {
		logger.Info("initiate payment for transfer-initiation %s", transferID)

		transferInitiationID := models.MustTransferInitiationIDFromString(transferID)
		transfer, err := getTransfer(ctx, storageReader, transferInitiationID, true)
		if err != nil {
			return err
		}

		attrs := metric.WithAttributes(connectorAttrs...)
		var paymentID *models.PaymentID
		defer func() {
			if err != nil {
				ctx, cancel := contextutil.Detached(ctx)
				defer cancel()
				metricsRegistry.ConnectorObjectsErrors().Add(ctx, 1, attrs)
				if err := ingester.UpdateTransferInitiationPaymentsStatus(ctx, transfer, paymentID, models.TransferInitiationStatusFailed, err.Error(), transfer.Attempts, time.Now()); err != nil {
					logger.Error("failed to update transfer initiation status: %v", err)
				}
			}
		}()

		err = ingester.UpdateTransferInitiationPaymentsStatus(ctx, transfer, paymentID, models.TransferInitiationStatusProcessing, "", transfer.Attempts, time.Now())
		if err != nil {
			return err
		}

		attrs = initiateTransferAttrs
		if transfer.Type == models.TransferInitiationTypePayout {
			attrs = initiatePayoutAttrs
		}

		logger.Info("initiate payment between", transfer.SourceAccountID, " and %s", transfer.DestinationAccountID)

		now := time.Now()
		defer func() {
			metricsRegistry.ConnectorObjectsLatency().Record(ctx, time.Since(now).Milliseconds(), attrs)
		}()

		if transfer.SourceAccount != nil {
			if transfer.SourceAccount.Type == models.AccountTypeExternal {
				err = errors.New("payin not implemented: source account must be an internal account")
				return err
			}
		}

		var curr string
		curr, _, err = currency.GetCurrencyAndPrecisionFromAsset(supportedCurrenciesWithDecimal, transfer.Asset)
		if err != nil {
			return err
		}

		stripeClient := client.Client(stripeClient)
		// If source account is nil, or equal to root (which is a special
		// account we create for stripe for the balance platform), we don't need
		// to set the stripe account.
		if transfer.SourceAccount != nil && transfer.SourceAccount.Reference != rootAccountReference {
			stripeClient = stripeClient.ForAccount(transfer.SourceAccountID.Reference)
		}

		var connectorPaymentID string
		var paymentType models.PaymentType
		ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
		defer cancel()
		switch transfer.DestinationAccount.Type {
		case models.AccountTypeInternal:
			// Transfer between internal accounts
			var resp *stripe.Transfer
			resp, err = stripeClient.CreateTransfer(ctx, &client.CreateTransferRequest{
				IdempotencyKey: fmt.Sprintf("%s_%d", transfer.ID.Reference, transfer.Attempts),
				Amount:         transfer.Amount.Int64(),
				Currency:       curr,
				Destination:    transfer.DestinationAccountID.Reference,
				Description:    transfer.Description,
			})
			if err != nil {
				return err
			}

			connectorPaymentID = resp.BalanceTransaction.ID
			paymentType = models.PaymentTypeTransfer
		case models.AccountTypeExternal:
			// Payout to an external account
			var resp *stripe.Payout
			resp, err = stripeClient.CreatePayout(ctx, &client.CreatePayoutRequest{
				IdempotencyKey: fmt.Sprintf("%s_%d", transfer.ID.Reference, transfer.Attempts),
				Amount:         transfer.Amount.Int64(),
				Currency:       curr,
				Destination:    transfer.DestinationAccountID.Reference,
				Description:    transfer.Description,
			})
			if err != nil {
				return err
			}

			connectorPaymentID = resp.BalanceTransaction.ID
			paymentType = models.PaymentTypePayOut
		}
		metricsRegistry.ConnectorObjects().Add(ctx, 1, attrs)

		paymentID = &models.PaymentID{
			PaymentReference: models.PaymentReference{
				Reference: connectorPaymentID,
				Type:      paymentType,
			},
			ConnectorID: connectorID,
		}
		err = ingester.AddTransferInitiationPaymentID(ctx, transfer, paymentID, time.Now())
		if err != nil {
			return err
		}

		taskDescriptor, err := models.EncodeTaskDescriptor(TaskDescriptor{
			Name:       "Update transfer initiation status",
			Key:        taskNameUpdatePaymentStatus,
			TransferID: transfer.ID.String(),
			PaymentID:  paymentID.String(),
			Attempt:    1,
		})
		if err != nil {
			return err
		}

		ctx, _ = contextutil.DetachedWithTimeout(ctx, 10*time.Second)
		err = scheduler.Schedule(ctx, taskDescriptor, models.TaskSchedulerOptions{
			ScheduleOption: models.OPTIONS_RUN_NOW,
			RestartOption:  models.OPTIONS_RESTART_IF_NOT_ACTIVE,
		})
		if err != nil && !errors.Is(err, task.ErrAlreadyScheduled) {
			return err
		}

		return nil
	}
}

var (
	updateTransferAttrs = metric.WithAttributes(append(connectorAttrs, attribute.String(metrics.ObjectAttributeKey, "update_transfer"))...)
	updatePayoutAttrs   = metric.WithAttributes(append(connectorAttrs, attribute.String(metrics.ObjectAttributeKey, "update_payout"))...)
)

func updatePaymentStatusTask(
	transferID string,
	pID string,
	attempt int,
	stripeClient *client.DefaultClient,
) task.Task {
	return func(
		ctx context.Context,
		logger logging.Logger,
		ingester ingestion.Ingester,
		scheduler task.Scheduler,
		storageReader storage.Reader,
		metricsRegistry metrics.MetricsRegistry,
	) error {
		paymentID := models.MustPaymentIDFromString(pID)
		transferInitiationID := models.MustTransferInitiationIDFromString(transferID)
		transfer, err := getTransfer(ctx, storageReader, transferInitiationID, false)
		if err != nil {
			return err
		}
		logger.Info("attempt: ", attempt, " fetching status of ", paymentID)

		attrs := updateTransferAttrs
		if transfer.Type == models.TransferInitiationTypePayout {
			attrs = updatePayoutAttrs
		}

		now := time.Now()
		defer func() {
			metricsRegistry.ConnectorObjectsLatency().Record(ctx, time.Since(now).Milliseconds(), attrs)
		}()

		defer func() {
			if err != nil {
				metricsRegistry.ConnectorObjectsErrors().Add(ctx, 1, attrs)
			}
		}()

		var status stripe.PayoutFailureCode
		var resultMessage string
		switch transfer.Type {
		case models.TransferInitiationTypeTransfer:
			// Nothing to do
			err = ingester.UpdateTransferInitiationPaymentsStatus(ctx, transfer, paymentID, models.TransferInitiationStatusProcessed, "", transfer.Attempts, time.Now())
			if err != nil {
				return err
			}

			return nil

		case models.TransferInitiationTypePayout:
			var resp *stripe.Payout
			resp, err = stripeClient.GetPayout(ctx, paymentID.Reference)
			if err != nil {
				return err
			}

			status = resp.FailureCode
			resultMessage = resp.FailureMessage
		}

		if status == "" {
			err = ingester.UpdateTransferInitiationPaymentsStatus(ctx, transfer, paymentID, models.TransferInitiationStatusProcessed, "", transfer.Attempts, time.Now())
			if err != nil {
				return err
			}

			return nil
		}

		err = ingester.UpdateTransferInitiationPaymentsStatus(ctx, transfer, paymentID, models.TransferInitiationStatusFailed, resultMessage, transfer.Attempts, time.Now())
		if err != nil {
			return err
		}

		return nil

	}
}

func getTransfer(
	ctx context.Context,
	reader storage.Reader,
	transferID models.TransferInitiationID,
	expand bool,
) (*models.TransferInitiation, error) {
	transfer, err := reader.ReadTransferInitiation(ctx, transferID)
	if err != nil {
		return nil, err
	}

	if expand {
		if transfer.SourceAccountID != nil {
			sourceAccount, err := reader.GetAccount(ctx, transfer.SourceAccountID.String())
			if err != nil {
				return nil, err
			}
			transfer.SourceAccount = sourceAccount
		}

		destinationAccount, err := reader.GetAccount(ctx, transfer.DestinationAccountID.String())
		if err != nil {
			return nil, err
		}
		transfer.DestinationAccount = destinationAccount
	}

	return transfer, nil
}

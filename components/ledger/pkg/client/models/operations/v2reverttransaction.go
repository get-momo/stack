// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/ledger/client/internal/utils"
	"github.com/formancehq/stack/ledger/client/models/components"
	"math/big"
)

type V2RevertTransactionRequest struct {
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Transaction ID.
	ID *big.Int `pathParam:"style=simple,explode=false,name=id"`
	// Force revert
	Force *bool `queryParam:"style=form,explode=true,name=force"`
	// Revert transaction at effective date of the original tx
	AtEffectiveDate *bool `queryParam:"style=form,explode=true,name=atEffectiveDate"`
}

func (v V2RevertTransactionRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2RevertTransactionRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2RevertTransactionRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2RevertTransactionRequest) GetID() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.ID
}

func (o *V2RevertTransactionRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *V2RevertTransactionRequest) GetAtEffectiveDate() *bool {
	if o == nil {
		return nil
	}
	return o.AtEffectiveDate
}

type V2RevertTransactionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	V2RevertTransactionResponse *components.V2RevertTransactionResponse
}

func (o *V2RevertTransactionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V2RevertTransactionResponse) GetV2RevertTransactionResponse() *components.V2RevertTransactionResponse {
	if o == nil {
		return nil
	}
	return o.V2RevertTransactionResponse
}

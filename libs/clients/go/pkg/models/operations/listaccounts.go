// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"net/http"
)

// ListAccountsBalanceOperator - Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
type ListAccountsBalanceOperator string

const (
	ListAccountsBalanceOperatorGte ListAccountsBalanceOperator = "gte"
	ListAccountsBalanceOperatorLte ListAccountsBalanceOperator = "lte"
	ListAccountsBalanceOperatorGt  ListAccountsBalanceOperator = "gt"
	ListAccountsBalanceOperatorLt  ListAccountsBalanceOperator = "lt"
	ListAccountsBalanceOperatorE   ListAccountsBalanceOperator = "e"
	ListAccountsBalanceOperatorNe  ListAccountsBalanceOperator = "ne"
)

func (e ListAccountsBalanceOperator) ToPointer() *ListAccountsBalanceOperator {
	return &e
}

func (e *ListAccountsBalanceOperator) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "gte":
		fallthrough
	case "lte":
		fallthrough
	case "gt":
		fallthrough
	case "lt":
		fallthrough
	case "e":
		fallthrough
	case "ne":
		*e = ListAccountsBalanceOperator(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ListAccountsBalanceOperator: %v", v)
	}
}

// ListAccountsMetadata - Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
type ListAccountsMetadata struct {
}

type ListAccountsRequest struct {
	// Filter accounts by address pattern (regular expression placed between ^ and $).
	Address *string `queryParam:"style=form,explode=true,name=address"`
	// Pagination cursor, will return accounts after given address, in descending order.
	After *string `queryParam:"style=form,explode=true,name=after"`
	// Filter accounts by their balance (default operator is gte)
	Balance *int64 `queryParam:"style=form,explode=true,name=balance"`
	// Operator used for the filtering of balances can be greater than/equal, less than/equal, greater than, less than, equal or not.
	//
	BalanceOperator *ListAccountsBalanceOperator `queryParam:"style=form,explode=true,name=balanceOperator"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
	// Name of the ledger.
	Ledger string `pathParam:"style=simple,explode=false,name=ledger"`
	// Filter accounts by metadata key value pairs. Nested objects can be used as seen in the example below.
	Metadata *ListAccountsMetadata `queryParam:"style=deepObject,explode=true,name=metadata"`
	// The maximum number of results to return per page.
	//
	PageSize *int64 `default:"15" queryParam:"style=form,explode=true,name=pageSize"`
}

func (l ListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsRequest) GetAddress() *string {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *ListAccountsRequest) GetAfter() *string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *ListAccountsRequest) GetBalance() *int64 {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *ListAccountsRequest) GetBalanceOperator() *ListAccountsBalanceOperator {
	if o == nil {
		return nil
	}
	return o.BalanceOperator
}

func (o *ListAccountsRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *ListAccountsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *ListAccountsRequest) GetMetadata() *ListAccountsMetadata {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *ListAccountsRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

type ListAccountsResponse struct {
	// OK
	AccountsCursorResponse *shared.AccountsCursorResponse
	// HTTP response content type for this operation
	ContentType string
	// Error
	ErrorResponse *shared.ErrorResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *ListAccountsResponse) GetAccountsCursorResponse() *shared.AccountsCursorResponse {
	if o == nil {
		return nil
	}
	return o.AccountsCursorResponse
}

func (o *ListAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountsResponse) GetErrorResponse() *shared.ErrorResponse {
	if o == nil {
		return nil
	}
	return o.ErrorResponse
}

func (o *ListAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
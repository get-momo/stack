// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/stack/ledger/client/models/components"
)

type V2ListLedgersRequest struct {
	// The maximum number of results to return per page.
	//
	PageSize *int64 `queryParam:"style=form,explode=true,name=pageSize"`
	// Parameter used in pagination requests. Maximum page size is set to 15.
	// Set to the value of next for the next page of results.
	// Set to the value of previous for the previous page of results.
	// No other parameters can be set when this parameter is set.
	//
	Cursor *string `queryParam:"style=form,explode=true,name=cursor"`
}

func (o *V2ListLedgersRequest) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V2ListLedgersRequest) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

type V2ListLedgersResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// OK
	V2LedgerListResponse *components.V2LedgerListResponse
}

func (o *V2ListLedgersResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *V2ListLedgersResponse) GetV2LedgerListResponse() *components.V2LedgerListResponse {
	if o == nil {
		return nil
	}
	return o.V2LedgerListResponse
}

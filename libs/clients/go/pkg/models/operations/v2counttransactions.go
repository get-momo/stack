// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"net/http"
	"time"
)

type V2CountTransactionsRequest struct {
	RequestBody map[string]interface{} `request:"mediaType=application/json"`
	// Name of the ledger.
	Ledger string     `pathParam:"style=simple,explode=false,name=ledger"`
	Pit    *time.Time `queryParam:"style=form,explode=true,name=pit"`
}

func (v V2CountTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V2CountTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V2CountTransactionsRequest) GetRequestBody() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.RequestBody
}

func (o *V2CountTransactionsRequest) GetLedger() string {
	if o == nil {
		return ""
	}
	return o.Ledger
}

func (o *V2CountTransactionsRequest) GetPit() *time.Time {
	if o == nil {
		return nil
	}
	return o.Pit
}

type V2CountTransactionsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	Headers     map[string][]string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Error
	V2ErrorResponse *shared.V2ErrorResponse
}

func (o *V2CountTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2CountTransactionsResponse) GetHeaders() map[string][]string {
	if o == nil {
		return nil
	}
	return o.Headers
}

func (o *V2CountTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2CountTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2CountTransactionsResponse) GetV2ErrorResponse() *shared.V2ErrorResponse {
	if o == nil {
		return nil
	}
	return o.V2ErrorResponse
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"net/http"
)

type V2GetInstanceStageHistoryRequest struct {
	// The instance id
	InstanceID string `pathParam:"style=simple,explode=false,name=instanceID"`
	// The stage number
	Number int64 `pathParam:"style=simple,explode=false,name=number"`
}

func (o *V2GetInstanceStageHistoryRequest) GetInstanceID() string {
	if o == nil {
		return ""
	}
	return o.InstanceID
}

func (o *V2GetInstanceStageHistoryRequest) GetNumber() int64 {
	if o == nil {
		return 0
	}
	return o.Number
}

type V2GetInstanceStageHistoryResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// General error
	V2Error *shared.V2Error
	// The workflow instance stage history
	V2GetWorkflowInstanceHistoryStageResponse *shared.V2GetWorkflowInstanceHistoryStageResponse
}

func (o *V2GetInstanceStageHistoryResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *V2GetInstanceStageHistoryResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *V2GetInstanceStageHistoryResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *V2GetInstanceStageHistoryResponse) GetV2Error() *shared.V2Error {
	if o == nil {
		return nil
	}
	return o.V2Error
}

func (o *V2GetInstanceStageHistoryResponse) GetV2GetWorkflowInstanceHistoryStageResponse() *shared.V2GetWorkflowInstanceHistoryStageResponse {
	if o == nil {
		return nil
	}
	return o.V2GetWorkflowInstanceHistoryStageResponse
}
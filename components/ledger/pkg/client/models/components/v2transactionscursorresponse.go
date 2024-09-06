// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package components

type V2TransactionsCursorResponseCursor struct {
	PageSize int64                   `json:"pageSize"`
	HasMore  bool                    `json:"hasMore"`
	Previous *string                 `json:"previous,omitempty"`
	Next     *string                 `json:"next,omitempty"`
	Data     []V2ExpandedTransaction `json:"data"`
}

func (o *V2TransactionsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *V2TransactionsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *V2TransactionsCursorResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

func (o *V2TransactionsCursorResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *V2TransactionsCursorResponseCursor) GetData() []V2ExpandedTransaction {
	if o == nil {
		return []V2ExpandedTransaction{}
	}
	return o.Data
}

type V2TransactionsCursorResponse struct {
	Cursor V2TransactionsCursorResponseCursor `json:"cursor"`
}

func (o *V2TransactionsCursorResponse) GetCursor() V2TransactionsCursorResponseCursor {
	if o == nil {
		return V2TransactionsCursorResponseCursor{}
	}
	return o.Cursor
}

// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StageSend struct {
	Amount      *Monetary             `json:"amount,omitempty"`
	Destination *StageSendDestination `json:"destination,omitempty"`
	Metadata    map[string]string     `json:"metadata,omitempty"`
	Source      *StageSendSource      `json:"source,omitempty"`
}

func (o *StageSend) GetAmount() *Monetary {
	if o == nil {
		return nil
	}
	return o.Amount
}

func (o *StageSend) GetDestination() *StageSendDestination {
	if o == nil {
		return nil
	}
	return o.Destination
}

func (o *StageSend) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *StageSend) GetSource() *StageSendSource {
	if o == nil {
		return nil
	}
	return o.Source
}
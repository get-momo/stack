// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type GetBalanceResponse struct {
	Data BalanceWithAssets `json:"data"`
}

func (o *GetBalanceResponse) GetData() BalanceWithAssets {
	if o == nil {
		return BalanceWithAssets{}
	}
	return o.Data
}
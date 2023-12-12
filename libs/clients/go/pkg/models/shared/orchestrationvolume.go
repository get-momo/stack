// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/pkg/utils"
	"math/big"
)

type OrchestrationVolume struct {
	Balance *big.Int `json:"balance,omitempty"`
	Input   *big.Int `json:"input"`
	Output  *big.Int `json:"output"`
}

func (o OrchestrationVolume) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(o, "", false)
}

func (o *OrchestrationVolume) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &o, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *OrchestrationVolume) GetBalance() *big.Int {
	if o == nil {
		return nil
	}
	return o.Balance
}

func (o *OrchestrationVolume) GetInput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Input
}

func (o *OrchestrationVolume) GetOutput() *big.Int {
	if o == nil {
		return big.NewInt(0)
	}
	return o.Output
}

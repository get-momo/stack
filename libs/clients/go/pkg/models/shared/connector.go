// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type Connector string

const (
	ConnectorStripe        Connector = "STRIPE"
	ConnectorDummyPay      Connector = "DUMMY-PAY"
	ConnectorWise          Connector = "WISE"
	ConnectorModulr        Connector = "MODULR"
	ConnectorCurrencyCloud Connector = "CURRENCY-CLOUD"
	ConnectorBankingCircle Connector = "BANKING-CIRCLE"
	ConnectorMangopay      Connector = "MANGOPAY"
	ConnectorMoneycorp     Connector = "MONEYCORP"
	ConnectorAtlar         Connector = "ATLAR"
)

func (e Connector) ToPointer() *Connector {
	return &e
}

func (e *Connector) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "STRIPE":
		fallthrough
	case "DUMMY-PAY":
		fallthrough
	case "WISE":
		fallthrough
	case "MODULR":
		fallthrough
	case "CURRENCY-CLOUD":
		fallthrough
	case "BANKING-CIRCLE":
		fallthrough
	case "MANGOPAY":
		fallthrough
	case "MONEYCORP":
		fallthrough
	case "ATLAR":
		*e = Connector(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Connector: %v", v)
	}
}

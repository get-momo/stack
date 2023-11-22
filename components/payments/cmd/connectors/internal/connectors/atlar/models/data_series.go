// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataSeries data series
//
// swagger:model DataSeries
type DataSeries struct {

	// A list of data points for an x- and y-axis.
	DataPoints []*DataPoint `json:"dataPoints"`

	// The grouping factor. Could e.g. be date, hour or account ID. Each statistic will support different grouping options.
	// Example: DATE
	GroupBy string `json:"groupBy,omitempty"`
}

// Validate validates this data series
func (m *DataSeries) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDataPoints(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSeries) validateDataPoints(formats strfmt.Registry) error {
	if swag.IsZero(m.DataPoints) { // not required
		return nil
	}

	for i := 0; i < len(m.DataPoints); i++ {
		if swag.IsZero(m.DataPoints[i]) { // not required
			continue
		}

		if m.DataPoints[i] != nil {
			if err := m.DataPoints[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this data series based on the context it is used
func (m *DataSeries) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDataPoints(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DataSeries) contextValidateDataPoints(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.DataPoints); i++ {

		if m.DataPoints[i] != nil {

			if swag.IsZero(m.DataPoints[i]) { // not required
				return nil
			}

			if err := m.DataPoints[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dataPoints" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DataSeries) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataSeries) UnmarshalBinary(b []byte) error {
	var res DataSeries
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
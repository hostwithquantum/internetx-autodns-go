// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Vat vat
//
// swagger:model Vat
type Vat struct {

	// Calculated vat amount
	Amount float64 `json:"amount,omitempty"`

	// Vat rate
	Rate *VatRate `json:"rate,omitempty"`
}

// Validate validates this vat
func (m *Vat) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Vat) validateRate(formats strfmt.Registry) error {

	if swag.IsZero(m.Rate) { // not required
		return nil
	}

	if m.Rate != nil {
		if err := m.Rate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("rate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Vat) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Vat) UnmarshalBinary(b []byte) error {
	var res Vat
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
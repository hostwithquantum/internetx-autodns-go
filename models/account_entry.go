// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AccountEntry account entry
//
// swagger:model AccountEntry
type AccountEntry struct {

	// Booking amount
	// Minimum: 1
	Amount float64 `json:"amount,omitempty"`

	// Used currency of the amount
	Currency string `json:"currency,omitempty"`

	// Name of the booking
	Label string `json:"label,omitempty"`

	// Net amount
	NetAmount float64 `json:"netAmount,omitempty"`

	// Value added tax sum
	VatAmount float64 `json:"vatAmount,omitempty"`

	// Value Added Tax
	Vats []*Vat `json:"vats"`
}

// Validate validates this account entry
func (m *AccountEntry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountEntry) validateAmount(formats strfmt.Registry) error {

	if swag.IsZero(m.Amount) { // not required
		return nil
	}

	if err := validate.Minimum("amount", "body", float64(m.Amount), 1, false); err != nil {
		return err
	}

	return nil
}

func (m *AccountEntry) validateVats(formats strfmt.Registry) error {

	if swag.IsZero(m.Vats) { // not required
		return nil
	}

	for i := 0; i < len(m.Vats); i++ {
		if swag.IsZero(m.Vats[i]) { // not required
			continue
		}

		if m.Vats[i] != nil {
			if err := m.Vats[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("vats" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountEntry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountEntry) UnmarshalBinary(b []byte) error {
	var res AccountEntry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
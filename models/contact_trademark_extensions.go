// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactTrademarkExtensions contact trademark extensions
//
// swagger:model ContactTrademarkExtensions
type ContactTrademarkExtensions struct {

	// Application date of the trade mark (yyyy-MM-dd).
	// Example: 2022-07-29
	// Format: date
	Appdate strfmt.Date `json:"appdate,omitempty"`

	// Country in which the trade mark is valid (ISO 3166 Country Code).
	// Example: DE
	Country string `json:"country,omitempty"`

	// Trademark name.
	Name string `json:"name,omitempty"`

	// Trademark number.
	Number string `json:"number,omitempty"`

	// The trademark office where the trademark is registered at.
	Office string `json:"office,omitempty"`

	// Registration date.
	// Format: date
	Regdate strfmt.Date `json:"regdate,omitempty"`
}

// Validate validates this contact trademark extensions
func (m *ContactTrademarkExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAppdate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegdate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactTrademarkExtensions) validateAppdate(formats strfmt.Registry) error {
	if swag.IsZero(m.Appdate) { // not required
		return nil
	}

	if err := validate.FormatOf("appdate", "body", "date", m.Appdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactTrademarkExtensions) validateRegdate(formats strfmt.Registry) error {
	if swag.IsZero(m.Regdate) { // not required
		return nil
	}

	if err := validate.FormatOf("regdate", "body", "date", m.Regdate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this contact trademark extensions based on context it is used
func (m *ContactTrademarkExtensions) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactTrademarkExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactTrademarkExtensions) UnmarshalBinary(b []byte) error {
	var res ContactTrademarkExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

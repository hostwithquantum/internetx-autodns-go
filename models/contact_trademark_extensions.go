// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ContactTrademarkExtensions contact trademark extensions
//
// swagger:model ContactTrademarkExtensions
type ContactTrademarkExtensions struct {

	// The application date.
	// Format: date-time
	Appdate strfmt.DateTime `json:"appdate,omitempty"`

	// The country.
	Country string `json:"country,omitempty"`

	// The name.
	Name string `json:"name,omitempty"`

	// The number.
	Number string `json:"number,omitempty"`

	// The office.
	Office string `json:"office,omitempty"`

	// The registration date.
	// Format: date-time
	Regdate strfmt.DateTime `json:"regdate,omitempty"`
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

	if err := validate.FormatOf("appdate", "body", "date-time", m.Appdate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContactTrademarkExtensions) validateRegdate(formats strfmt.Registry) error {

	if swag.IsZero(m.Regdate) { // not required
		return nil
	}

	if err := validate.FormatOf("regdate", "body", "date-time", m.Regdate.String(), formats); err != nil {
		return err
	}

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

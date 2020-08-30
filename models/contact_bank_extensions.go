// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactBankExtensions contact bank extensions
//
// swagger:model ContactBankExtensions
type ContactBankExtensions struct {

	// The FTLD token.
	FtldToken string `json:"ftldToken,omitempty"`
}

// Validate validates this contact bank extensions
func (m *ContactBankExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactBankExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactBankExtensions) UnmarshalBinary(b []byte) error {
	var res ContactBankExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
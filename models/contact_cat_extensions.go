// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactCatExtensions contact cat extensions
//
// swagger:model ContactCatExtensions
type ContactCatExtensions struct {

	// The statement of intended use for the domain name.
	IntendedUse string `json:"intendedUse,omitempty"`
}

// Validate validates this contact cat extensions
func (m *ContactCatExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactCatExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactCatExtensions) UnmarshalBinary(b []byte) error {
	var res ContactCatExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
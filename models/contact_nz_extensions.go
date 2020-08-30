// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactNzExtensions contact nz extensions
//
// swagger:model ContactNzExtensions
type ContactNzExtensions struct {

	// The irpo.
	Irpo bool `json:"irpo,omitempty"`
}

// Validate validates this contact nz extensions
func (m *ContactNzExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactNzExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactNzExtensions) UnmarshalBinary(b []byte) error {
	var res ContactNzExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

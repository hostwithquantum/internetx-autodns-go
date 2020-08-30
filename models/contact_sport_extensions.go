// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactSportExtensions contact sport extensions
//
// swagger:model ContactSportExtensions
type ContactSportExtensions struct {

	// The intended use.
	IntendedUse string `json:"intendedUse,omitempty"`
}

// Validate validates this contact sport extensions
func (m *ContactSportExtensions) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ContactSportExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactSportExtensions) UnmarshalBinary(b []byte) error {
	var res ContactSportExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

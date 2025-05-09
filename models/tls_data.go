// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TLSData Tls data
//
// swagger:model TlsData
type TLSData struct {

	// Whether the version is supported or not
	Supported bool `json:"supported,omitempty"`

	// The tls version
	Version string `json:"version,omitempty"`
}

// Validate validates this Tls data
func (m *TLSData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this Tls data based on context it is used
func (m *TLSData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TLSData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TLSData) UnmarshalBinary(b []byte) error {
	var res TLSData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

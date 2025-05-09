// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateInstallCheckNotice certificate install check notice
//
// swagger:model CertificateInstallCheckNotice
type CertificateInstallCheckNotice struct {

	// The label of the notice
	Label string `json:"label,omitempty"`

	// The type of the notice
	Type string `json:"type,omitempty"`

	// The value of the notice
	Value string `json:"value,omitempty"`
}

// Validate validates this certificate install check notice
func (m *CertificateInstallCheckNotice) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certificate install check notice based on context it is used
func (m *CertificateInstallCheckNotice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateInstallCheckNotice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateInstallCheckNotice) UnmarshalBinary(b []byte) error {
	var res CertificateInstallCheckNotice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

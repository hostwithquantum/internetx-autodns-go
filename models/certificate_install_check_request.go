// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CertificateInstallCheckRequest certificate install check request
//
// swagger:model CertificateInstallCheckRequest
type CertificateInstallCheckRequest struct {

	// The hostname to check
	Hostname string `json:"hostname,omitempty"`
}

// Validate validates this certificate install check request
func (m *CertificateInstallCheckRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this certificate install check request based on context it is used
func (m *CertificateInstallCheckRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateInstallCheckRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateInstallCheckRequest) UnmarshalBinary(b []byte) error {
	var res CertificateInstallCheckRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

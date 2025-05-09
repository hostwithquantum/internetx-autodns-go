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

// CertificateHistory certificate history
//
// swagger:model CertificateHistory
type CertificateHistory struct {

	// The unique certificate order number
	OrderID string `json:"orderId,omitempty"`

	// The revoked date of the certificate.
	// Format: date-time
	Revoked strfmt.DateTime `json:"revoked,omitempty"`

	// The serialnumber of the certificate.
	SerialNumber string `json:"serialNumber,omitempty"`
}

// Validate validates this certificate history
func (m *CertificateHistory) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRevoked(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CertificateHistory) validateRevoked(formats strfmt.Registry) error {
	if swag.IsZero(m.Revoked) { // not required
		return nil
	}

	if err := validate.FormatOf("revoked", "body", "date-time", m.Revoked.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this certificate history based on context it is used
func (m *CertificateHistory) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CertificateHistory) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertificateHistory) UnmarshalBinary(b []byte) error {
	var res CertificateHistory
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

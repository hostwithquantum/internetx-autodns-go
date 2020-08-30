// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactCaExtensions contact ca extensions
//
// swagger:model ContactCaExtensions
type ContactCaExtensions struct {

	// The agreement version.
	AgreementVersion float64 `json:"agreementVersion,omitempty"`

	// The cira cpr.
	Cpr CiraCprConstants `json:"cpr,omitempty"`

	// The official representative.
	OfficialRepresentative string `json:"officialRepresentative,omitempty"`

	// The originating ip.
	OriginatingIP string `json:"originatingIp,omitempty"`

	// The trustee percentage.
	TrusteePercentage float64 `json:"trusteePercentage,omitempty"`
}

// Validate validates this contact ca extensions
func (m *ContactCaExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCpr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactCaExtensions) validateCpr(formats strfmt.Registry) error {

	if swag.IsZero(m.Cpr) { // not required
		return nil
	}

	if err := m.Cpr.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("cpr")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactCaExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactCaExtensions) UnmarshalBinary(b []byte) error {
	var res ContactCaExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

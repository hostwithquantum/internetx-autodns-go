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

// DomainTmchClaimNoticeExtensions domain tmch claim notice extensions
//
// swagger:model DomainTmchClaimNoticeExtensions
type DomainTmchClaimNoticeExtensions struct {

	// The confirm ip
	ConfirmIP string `json:"confirmIp,omitempty"`

	// The date of confirmation
	// Format: date-time
	Confirmed strfmt.DateTime `json:"confirmed,omitempty"`

	// The external refernce
	ExternalReference string `json:"externalReference,omitempty"`

	// Date after the claim expires
	// Format: date-time
	NotAfter strfmt.DateTime `json:"notAfter,omitempty"`
}

// Validate validates this domain tmch claim notice extensions
func (m *DomainTmchClaimNoticeExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfirmed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNotAfter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainTmchClaimNoticeExtensions) validateConfirmed(formats strfmt.Registry) error {

	if swag.IsZero(m.Confirmed) { // not required
		return nil
	}

	if err := validate.FormatOf("confirmed", "body", "date-time", m.Confirmed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainTmchClaimNoticeExtensions) validateNotAfter(formats strfmt.Registry) error {

	if swag.IsZero(m.NotAfter) { // not required
		return nil
	}

	if err := validate.FormatOf("notAfter", "body", "date-time", m.NotAfter.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainTmchClaimNoticeExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainTmchClaimNoticeExtensions) UnmarshalBinary(b []byte) error {
	var res DomainTmchClaimNoticeExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

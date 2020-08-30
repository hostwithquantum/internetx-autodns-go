// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactAuExtensions contact au extensions
//
// swagger:model ContactAuExtensions
type ContactAuExtensions struct {

	// The corresponding ID number for the eligibility name if applicable, if in doubt just use the same details as your Registrant ID.
	EligibilityID string `json:"eligibilityId,omitempty"`

	// The type of Eligibility ID.
	EligibilityIDType AuEligibilityIDTypeConstants `json:"eligibilityIdType,omitempty"`

	// 	This field is usually optional and whether it is required is determined by your policy reason for the domain registration, if in doubt just use the same details as your Registrant Name.
	EligibilityName string `json:"eligibilityName,omitempty"`

	// The type of entity that the registrant is.
	EligibilityType AuEligibilityTypeConstants `json:"eligibilityType,omitempty"`

	// Reason for eligibility.
	PolicyReason int32 `json:"policyReason,omitempty"`

	// Australian Business Number or Registered Business Number.
	RegistrantID string `json:"registrantId,omitempty"`

	// Concerns the Registrant ID which you have selected in the Registrant ID form field.
	RegistrantIDType AuRegistrantIDTypeConstants `json:"registrantIdType,omitempty"`

	// The registrants name.
	RegistrantName string `json:"registrantName,omitempty"`
}

// Validate validates this contact au extensions
func (m *ContactAuExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEligibilityIDType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEligibilityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistrantIDType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactAuExtensions) validateEligibilityIDType(formats strfmt.Registry) error {

	if swag.IsZero(m.EligibilityIDType) { // not required
		return nil
	}

	if err := m.EligibilityIDType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligibilityIdType")
		}
		return err
	}

	return nil
}

func (m *ContactAuExtensions) validateEligibilityType(formats strfmt.Registry) error {

	if swag.IsZero(m.EligibilityType) { // not required
		return nil
	}

	if err := m.EligibilityType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("eligibilityType")
		}
		return err
	}

	return nil
}

func (m *ContactAuExtensions) validateRegistrantIDType(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistrantIDType) { // not required
		return nil
	}

	if err := m.RegistrantIDType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("registrantIdType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactAuExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactAuExtensions) UnmarshalBinary(b []byte) error {
	var res ContactAuExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

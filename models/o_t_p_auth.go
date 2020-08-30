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

// OTPAuth o t p auth
//
// swagger:model OTPAuth
type OTPAuth struct {

	// Crypto algorithm
	// Required: true
	Algorithm CryptoFormatConstants `json:"algorithm"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The length of the token
	// Required: true
	Digits *int32 `json:"digits"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// The type of protocol
	// Required: true
	ProtocolType ProtocolTypeConstants `json:"protocolType"`

	// The generated qr code png.
	QrCode *BasicDocument `json:"qrCode,omitempty"`

	// Base32 encoded shared secret.
	Secret string `json:"secret,omitempty"`

	// Timeout in seconds
	// Required: true
	// Maximum: 90
	// Minimum: 5
	Timeout *int32 `json:"timeout"`

	// The generated support 'tokens'.
	// Unique: true
	Tokens []string `json:"tokens"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating user of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this o t p auth
func (m *OTPAuth) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlgorithm(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigits(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocolType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQrCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeout(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTokens(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdater(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OTPAuth) validateAlgorithm(formats strfmt.Registry) error {

	if err := m.Algorithm.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("algorithm")
		}
		return err
	}

	return nil
}

func (m *OTPAuth) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OTPAuth) validateDigits(formats strfmt.Registry) error {

	if err := validate.Required("digits", "body", m.Digits); err != nil {
		return err
	}

	return nil
}

func (m *OTPAuth) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *OTPAuth) validateProtocolType(formats strfmt.Registry) error {

	if err := m.ProtocolType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protocolType")
		}
		return err
	}

	return nil
}

func (m *OTPAuth) validateQrCode(formats strfmt.Registry) error {

	if swag.IsZero(m.QrCode) { // not required
		return nil
	}

	if m.QrCode != nil {
		if err := m.QrCode.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qrCode")
			}
			return err
		}
	}

	return nil
}

func (m *OTPAuth) validateTimeout(formats strfmt.Registry) error {

	if err := validate.Required("timeout", "body", m.Timeout); err != nil {
		return err
	}

	if err := validate.MinimumInt("timeout", "body", int64(*m.Timeout), 5, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("timeout", "body", int64(*m.Timeout), 90, false); err != nil {
		return err
	}

	return nil
}

func (m *OTPAuth) validateTokens(formats strfmt.Registry) error {

	if swag.IsZero(m.Tokens) { // not required
		return nil
	}

	if err := validate.UniqueItems("tokens", "body", m.Tokens); err != nil {
		return err
	}

	return nil
}

func (m *OTPAuth) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *OTPAuth) validateUpdater(formats strfmt.Registry) error {

	if swag.IsZero(m.Updater) { // not required
		return nil
	}

	if m.Updater != nil {
		if err := m.Updater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OTPAuth) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OTPAuth) UnmarshalBinary(b []byte) error {
	var res OTPAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
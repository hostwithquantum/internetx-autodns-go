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

// LoginData login data
//
// swagger:model LoginData
type LoginData struct {

	// The context.
	// Required: true
	Context *int32 `json:"context"`

	// The password.
	Password string `json:"password,omitempty"`

	// The one time password in case of 2fa authentication.
	Token string `json:"token,omitempty"`

	// The user name.
	// Required: true
	// Pattern: ^[^_].*
	User *string `json:"user"`
}

// Validate validates this login data
func (m *LoginData) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LoginData) validateContext(formats strfmt.Registry) error {

	if err := validate.Required("context", "body", m.Context); err != nil {
		return err
	}

	return nil
}

func (m *LoginData) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if err := validate.Pattern("user", "body", string(*m.User), `^[^_].*`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LoginData) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LoginData) UnmarshalBinary(b []byte) error {
	var res LoginData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

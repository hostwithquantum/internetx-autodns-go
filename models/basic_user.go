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

// BasicUser basic user
//
// swagger:model BasicUser
type BasicUser struct {

	// The context. A separated section.
	Context int32 `json:"context,omitempty"`

	// The date on which the password has changed
	// Format: date-time
	PasswordChanged strfmt.DateTime `json:"passwordChanged,omitempty"`

	// Is the current password expired
	PasswordExpired bool `json:"passwordExpired,omitempty"`

	// The date on which the password will expire
	// Format: date-time
	PasswordExpires strfmt.DateTime `json:"passwordExpires,omitempty"`

	// The user name.
	User string `json:"user,omitempty"`
}

// Validate validates this basic user
func (m *BasicUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePasswordChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePasswordExpires(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BasicUser) validatePasswordChanged(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordChanged) { // not required
		return nil
	}

	if err := validate.FormatOf("passwordChanged", "body", "date-time", m.PasswordChanged.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BasicUser) validatePasswordExpires(formats strfmt.Registry) error {
	if swag.IsZero(m.PasswordExpires) { // not required
		return nil
	}

	if err := validate.FormatOf("passwordExpires", "body", "date-time", m.PasswordExpires.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this basic user based on context it is used
func (m *BasicUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BasicUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BasicUser) UnmarshalBinary(b []byte) error {
	var res BasicUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

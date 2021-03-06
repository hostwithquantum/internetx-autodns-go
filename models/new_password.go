// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPassword new password
//
// swagger:model NewPassword
type NewPassword struct {

	// The context of the user
	Context int32 `json:"context,omitempty"`

	// The disguised email to which the TAN is to be sent.
	Email string `json:"email,omitempty"`

	// The selected method
	Method string `json:"method,omitempty"`

	// The disguised mobilephone number to which the TAN is to be sent.
	Mobile string `json:"mobile,omitempty"`

	// The password
	Password string `json:"password,omitempty"`

	// The tan
	Tan string `json:"tan,omitempty"`

	// The tan methods available for the user
	TanMethods []string `json:"tanMethods"`

	// The token
	Token string `json:"token,omitempty"`

	// The user
	User string `json:"user,omitempty"`
}

// Validate validates this new password
func (m *NewPassword) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NewPassword) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NewPassword) UnmarshalBinary(b []byte) error {
	var res NewPassword
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

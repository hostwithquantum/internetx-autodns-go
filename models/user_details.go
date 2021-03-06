// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UserDetails user details
//
// swagger:model UserDetails
type UserDetails struct {

	// The first name.
	Fname string `json:"fname,omitempty"`

	// The last name.
	Lname string `json:"lname,omitempty"`

	// The organization.
	Organization string `json:"organization,omitempty"`

	// The email address for the password reset tan.
	PasswordResetEmail string `json:"passwordResetEmail,omitempty"`

	// The mobile phone number.
	PasswordResetMobile string `json:"passwordResetMobile,omitempty"`

	// The email address for the verification of the password reset process.
	PasswordResetVerifyEmail string `json:"passwordResetVerifyEmail,omitempty"`

	// The phone number of the guest user
	Phone string `json:"phone,omitempty"`
}

// Validate validates this user details
func (m *UserDetails) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserDetails) UnmarshalBinary(b []byte) error {
	var res UserDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

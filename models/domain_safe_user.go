// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainSafeUser domain safe user
//
// swagger:model DomainSafeUser
type DomainSafeUser struct {

	// The context.
	Context int32 `json:"context,omitempty"`

	// The mobile of the user.
	Mobile Phone `json:"mobile,omitempty"`

	// The pin.
	Pin string `json:"pin,omitempty"`

	// The user name.
	User string `json:"user,omitempty"`
}

// Validate validates this domain safe user
func (m *DomainSafeUser) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain safe user based on context it is used
func (m *DomainSafeUser) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainSafeUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainSafeUser) UnmarshalBinary(b []byte) error {
	var res DomainSafeUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

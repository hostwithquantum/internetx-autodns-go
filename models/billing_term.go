// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BillingTerm billing term
//
// swagger:model BillingTerm
type BillingTerm struct {

	// terms
	Terms []BillingTldTerm `json:"terms"`
}

// Validate validates this billing term
func (m *BillingTerm) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this billing term based on context it is used
func (m *BillingTerm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *BillingTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingTerm) UnmarshalBinary(b []byte) error {
	var res BillingTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

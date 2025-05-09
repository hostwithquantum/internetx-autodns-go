// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainControllValidationStatus domain controll validation status
//
// swagger:model DomainControllValidationStatus
type DomainControllValidationStatus struct {

	// The domain.
	Domain string `json:"domain,omitempty"`

	// The status of the dcv.
	Status string `json:"status,omitempty"`
}

// Validate validates this domain controll validation status
func (m *DomainControllValidationStatus) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this domain controll validation status based on context it is used
func (m *DomainControllValidationStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DomainControllValidationStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainControllValidationStatus) UnmarshalBinary(b []byte) error {
	var res DomainControllValidationStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

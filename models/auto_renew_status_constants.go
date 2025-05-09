// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AutoRenewStatusConstants auto renew status constants
//
// swagger:model AutoRenewStatusConstants
type AutoRenewStatusConstants string

func NewAutoRenewStatusConstants(value AutoRenewStatusConstants) *AutoRenewStatusConstants {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AutoRenewStatusConstants.
func (m AutoRenewStatusConstants) Pointer() *AutoRenewStatusConstants {
	return &m
}

const (

	// AutoRenewStatusConstantsTRUE captures enum value "TRUE"
	AutoRenewStatusConstantsTRUE AutoRenewStatusConstants = "TRUE"

	// AutoRenewStatusConstantsFALSE captures enum value "FALSE"
	AutoRenewStatusConstantsFALSE AutoRenewStatusConstants = "FALSE"

	// AutoRenewStatusConstantsONCE captures enum value "ONCE"
	AutoRenewStatusConstantsONCE AutoRenewStatusConstants = "ONCE"
)

// for schema
var autoRenewStatusConstantsEnum []interface{}

func init() {
	var res []AutoRenewStatusConstants
	if err := json.Unmarshal([]byte(`["TRUE","FALSE","ONCE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		autoRenewStatusConstantsEnum = append(autoRenewStatusConstantsEnum, v)
	}
}

func (m AutoRenewStatusConstants) validateAutoRenewStatusConstantsEnum(path, location string, value AutoRenewStatusConstants) error {
	if err := validate.EnumCase(path, location, value, autoRenewStatusConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this auto renew status constants
func (m AutoRenewStatusConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAutoRenewStatusConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this auto renew status constants based on context it is used
func (m AutoRenewStatusConstants) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

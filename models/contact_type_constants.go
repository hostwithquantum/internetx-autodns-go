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

// ContactTypeConstants contact type constants
//
// swagger:model ContactTypeConstants
type ContactTypeConstants string

func NewContactTypeConstants(value ContactTypeConstants) *ContactTypeConstants {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ContactTypeConstants.
func (m ContactTypeConstants) Pointer() *ContactTypeConstants {
	return &m
}

const (

	// ContactTypeConstantsPERSON captures enum value "PERSON"
	ContactTypeConstantsPERSON ContactTypeConstants = "PERSON"

	// ContactTypeConstantsORG captures enum value "ORG"
	ContactTypeConstantsORG ContactTypeConstants = "ORG"

	// ContactTypeConstantsROLE captures enum value "ROLE"
	ContactTypeConstantsROLE ContactTypeConstants = "ROLE"
)

// for schema
var contactTypeConstantsEnum []interface{}

func init() {
	var res []ContactTypeConstants
	if err := json.Unmarshal([]byte(`["PERSON","ORG","ROLE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		contactTypeConstantsEnum = append(contactTypeConstantsEnum, v)
	}
}

func (m ContactTypeConstants) validateContactTypeConstantsEnum(path, location string, value ContactTypeConstants) error {
	if err := validate.EnumCase(path, location, value, contactTypeConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this contact type constants
func (m ContactTypeConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContactTypeConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this contact type constants based on context it is used
func (m ContactTypeConstants) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

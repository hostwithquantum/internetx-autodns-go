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

// ProtocolTypeConstants protocol type constants
//
// swagger:model ProtocolTypeConstants
type ProtocolTypeConstants string

func NewProtocolTypeConstants(value ProtocolTypeConstants) *ProtocolTypeConstants {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ProtocolTypeConstants.
func (m ProtocolTypeConstants) Pointer() *ProtocolTypeConstants {
	return &m
}

const (

	// ProtocolTypeConstantsTOTP captures enum value "TOTP"
	ProtocolTypeConstantsTOTP ProtocolTypeConstants = "TOTP"

	// ProtocolTypeConstantsHOTP captures enum value "HOTP"
	ProtocolTypeConstantsHOTP ProtocolTypeConstants = "HOTP"
)

// for schema
var protocolTypeConstantsEnum []interface{}

func init() {
	var res []ProtocolTypeConstants
	if err := json.Unmarshal([]byte(`["TOTP","HOTP"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		protocolTypeConstantsEnum = append(protocolTypeConstantsEnum, v)
	}
}

func (m ProtocolTypeConstants) validateProtocolTypeConstantsEnum(path, location string, value ProtocolTypeConstants) error {
	if err := validate.EnumCase(path, location, value, protocolTypeConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this protocol type constants
func (m ProtocolTypeConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateProtocolTypeConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this protocol type constants based on context it is used
func (m ProtocolTypeConstants) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

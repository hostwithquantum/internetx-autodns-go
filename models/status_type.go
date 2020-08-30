// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// StatusType status type
//
// swagger:model StatusType
type StatusType string

const (

	// StatusTypeSUCCESS captures enum value "SUCCESS"
	StatusTypeSUCCESS StatusType = "SUCCESS"

	// StatusTypeERROR captures enum value "ERROR"
	StatusTypeERROR StatusType = "ERROR"

	// StatusTypeNOTIFY captures enum value "NOTIFY"
	StatusTypeNOTIFY StatusType = "NOTIFY"

	// StatusTypeNOTICE captures enum value "NOTICE"
	StatusTypeNOTICE StatusType = "NOTICE"

	// StatusTypeNICCOMNOTIFY captures enum value "NICCOM_NOTIFY"
	StatusTypeNICCOMNOTIFY StatusType = "NICCOM_NOTIFY"
)

// for schema
var statusTypeEnum []interface{}

func init() {
	var res []StatusType
	if err := json.Unmarshal([]byte(`["SUCCESS","ERROR","NOTIFY","NOTICE","NICCOM_NOTIFY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statusTypeEnum = append(statusTypeEnum, v)
	}
}

func (m StatusType) validateStatusTypeEnum(path, location string, value StatusType) error {
	if err := validate.EnumCase(path, location, value, statusTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this status type
func (m StatusType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateStatusTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

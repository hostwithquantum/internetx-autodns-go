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

// ObjectAssignmentMode object assignment mode
//
// swagger:model ObjectAssignmentMode
type ObjectAssignmentMode string

const (

	// ObjectAssignmentModeASSIGN captures enum value "ASSIGN"
	ObjectAssignmentModeASSIGN ObjectAssignmentMode = "ASSIGN"

	// ObjectAssignmentModeDELETE captures enum value "DELETE"
	ObjectAssignmentModeDELETE ObjectAssignmentMode = "DELETE"
)

// for schema
var objectAssignmentModeEnum []interface{}

func init() {
	var res []ObjectAssignmentMode
	if err := json.Unmarshal([]byte(`["ASSIGN","DELETE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectAssignmentModeEnum = append(objectAssignmentModeEnum, v)
	}
}

func (m ObjectAssignmentMode) validateObjectAssignmentModeEnum(path, location string, value ObjectAssignmentMode) error {
	if err := validate.EnumCase(path, location, value, objectAssignmentModeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object assignment mode
func (m ObjectAssignmentMode) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectAssignmentModeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

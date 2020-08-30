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

// Operator operator
//
// swagger:model Operator
type Operator string

const (

	// OperatorEQUAL captures enum value "EQUAL"
	OperatorEQUAL Operator = "EQUAL"

	// OperatorNOTEQUAL captures enum value "NOT_EQUAL"
	OperatorNOTEQUAL Operator = "NOT_EQUAL"

	// OperatorNOTLIKE captures enum value "NOT_LIKE"
	OperatorNOTLIKE Operator = "NOT_LIKE"

	// OperatorLIKE captures enum value "LIKE"
	OperatorLIKE Operator = "LIKE"

	// OperatorILIKE captures enum value "ILIKE"
	OperatorILIKE Operator = "ILIKE"

	// OperatorGREATER captures enum value "GREATER"
	OperatorGREATER Operator = "GREATER"

	// OperatorGREATEREQUAL captures enum value "GREATER_EQUAL"
	OperatorGREATEREQUAL Operator = "GREATER_EQUAL"

	// OperatorLESS captures enum value "LESS"
	OperatorLESS Operator = "LESS"

	// OperatorLESSEQUAL captures enum value "LESS_EQUAL"
	OperatorLESSEQUAL Operator = "LESS_EQUAL"

	// OperatorISNULL captures enum value "IS_NULL"
	OperatorISNULL Operator = "IS_NULL"

	// OperatorISNOTNULL captures enum value "IS_NOT_NULL"
	OperatorISNOTNULL Operator = "IS_NOT_NULL"

	// OperatorIN captures enum value "IN"
	OperatorIN Operator = "IN"
)

// for schema
var operatorEnum []interface{}

func init() {
	var res []Operator
	if err := json.Unmarshal([]byte(`["EQUAL","NOT_EQUAL","NOT_LIKE","LIKE","ILIKE","GREATER","GREATER_EQUAL","LESS","LESS_EQUAL","IS_NULL","IS_NOT_NULL","IN"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		operatorEnum = append(operatorEnum, v)
	}
}

func (m Operator) validateOperatorEnum(path, location string, value Operator) error {
	if err := validate.EnumCase(path, location, value, operatorEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this operator
func (m Operator) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateOperatorEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

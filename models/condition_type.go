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

// ConditionType condition type
//
// swagger:model ConditionType
type ConditionType string

func NewConditionType(value ConditionType) *ConditionType {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ConditionType.
func (m ConditionType) Pointer() *ConditionType {
	return &m
}

const (

	// ConditionTypeAND captures enum value "AND"
	ConditionTypeAND ConditionType = "AND"

	// ConditionTypeOR captures enum value "OR"
	ConditionTypeOR ConditionType = "OR"
)

// for schema
var conditionTypeEnum []interface{}

func init() {
	var res []ConditionType
	if err := json.Unmarshal([]byte(`["AND","OR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		conditionTypeEnum = append(conditionTypeEnum, v)
	}
}

func (m ConditionType) validateConditionTypeEnum(path, location string, value ConditionType) error {
	if err := validate.EnumCase(path, location, value, conditionTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this condition type
func (m ConditionType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateConditionTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this condition type based on context it is used
func (m ConditionType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

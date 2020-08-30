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

// UserProfileFlag user profile flag
//
// swagger:model UserProfileFlag
type UserProfileFlag string

const (

	// UserProfileFlagOPTIONAL captures enum value "OPTIONAL"
	UserProfileFlagOPTIONAL UserProfileFlag = "OPTIONAL"

	// UserProfileFlagFIX captures enum value "FIX"
	UserProfileFlagFIX UserProfileFlag = "FIX"

	// UserProfileFlagRECURSE captures enum value "RECURSE"
	UserProfileFlagRECURSE UserProfileFlag = "RECURSE"

	// UserProfileFlagHIDDEN captures enum value "HIDDEN"
	UserProfileFlagHIDDEN UserProfileFlag = "HIDDEN"

	// UserProfileFlagPARENTFIX captures enum value "PARENT_FIX"
	UserProfileFlagPARENTFIX UserProfileFlag = "PARENT_FIX"
)

// for schema
var userProfileFlagEnum []interface{}

func init() {
	var res []UserProfileFlag
	if err := json.Unmarshal([]byte(`["OPTIONAL","FIX","RECURSE","HIDDEN","PARENT_FIX"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		userProfileFlagEnum = append(userProfileFlagEnum, v)
	}
}

func (m UserProfileFlag) validateUserProfileFlagEnum(path, location string, value UserProfileFlag) error {
	if err := validate.EnumCase(path, location, value, userProfileFlagEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this user profile flag
func (m UserProfileFlag) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateUserProfileFlagEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
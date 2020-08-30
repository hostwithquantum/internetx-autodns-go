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

// ACLRestriction ACL restriction
//
// swagger:model ACLRestriction
type ACLRestriction string

const (

	// ACLRestrictionUSERLOCKED captures enum value "USER_LOCKED"
	ACLRestrictionUSERLOCKED ACLRestriction = "USER_LOCKED"

	// ACLRestrictionCHILDRENLOCKED captures enum value "CHILDREN_LOCKED"
	ACLRestrictionCHILDRENLOCKED ACLRestriction = "CHILDREN_LOCKED"

	// ACLRestrictionPARENTLOCK captures enum value "PARENT_LOCK"
	ACLRestrictionPARENTLOCK ACLRestriction = "PARENT_LOCK"

	// ACLRestrictionNOTLOCKED captures enum value "NOT_LOCKED"
	ACLRestrictionNOTLOCKED ACLRestriction = "NOT_LOCKED"
)

// for schema
var aclRestrictionEnum []interface{}

func init() {
	var res []ACLRestriction
	if err := json.Unmarshal([]byte(`["USER_LOCKED","CHILDREN_LOCKED","PARENT_LOCK","NOT_LOCKED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		aclRestrictionEnum = append(aclRestrictionEnum, v)
	}
}

func (m ACLRestriction) validateACLRestrictionEnum(path, location string, value ACLRestriction) error {
	if err := validate.EnumCase(path, location, value, aclRestrictionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ACL restriction
func (m ACLRestriction) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateACLRestrictionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

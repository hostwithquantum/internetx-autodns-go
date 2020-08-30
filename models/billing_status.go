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

// BillingStatus billing status
//
// swagger:model BillingStatus
type BillingStatus string

const (

	// BillingStatusADD captures enum value "ADD"
	BillingStatusADD BillingStatus = "ADD"

	// BillingStatusREMOVE captures enum value "REMOVE"
	BillingStatusREMOVE BillingStatus = "REMOVE"

	// BillingStatusUPDATE captures enum value "UPDATE"
	BillingStatusUPDATE BillingStatus = "UPDATE"

	// BillingStatusIGNORE captures enum value "IGNORE"
	BillingStatusIGNORE BillingStatus = "IGNORE"

	// BillingStatusCUSTOMERCHANGED captures enum value "CUSTOMER_CHANGED"
	BillingStatusCUSTOMERCHANGED BillingStatus = "CUSTOMER_CHANGED"

	// BillingStatusUSERCHANGED captures enum value "USER_CHANGED"
	BillingStatusUSERCHANGED BillingStatus = "USER_CHANGED"

	// BillingStatusACTIVE captures enum value "ACTIVE"
	BillingStatusACTIVE BillingStatus = "ACTIVE"

	// BillingStatusCANCELED captures enum value "CANCELED"
	BillingStatusCANCELED BillingStatus = "CANCELED"

	// BillingStatusCANCELEDEXPIRE captures enum value "CANCELED_EXPIRE"
	BillingStatusCANCELEDEXPIRE BillingStatus = "CANCELED_EXPIRE"

	// BillingStatusRESTORE captures enum value "RESTORE"
	BillingStatusRESTORE BillingStatus = "RESTORE"

	// BillingStatusAUTODELETE captures enum value "AUTODELETE"
	BillingStatusAUTODELETE BillingStatus = "AUTODELETE"

	// BillingStatusREMOVED captures enum value "REMOVED"
	BillingStatusREMOVED BillingStatus = "REMOVED"
)

// for schema
var billingStatusEnum []interface{}

func init() {
	var res []BillingStatus
	if err := json.Unmarshal([]byte(`["ADD","REMOVE","UPDATE","IGNORE","CUSTOMER_CHANGED","USER_CHANGED","ACTIVE","CANCELED","CANCELED_EXPIRE","RESTORE","AUTODELETE","REMOVED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billingStatusEnum = append(billingStatusEnum, v)
	}
}

func (m BillingStatus) validateBillingStatusEnum(path, location string, value BillingStatus) error {
	if err := validate.EnumCase(path, location, value, billingStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this billing status
func (m BillingStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBillingStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

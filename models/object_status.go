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

// ObjectStatus object status
//
// swagger:model ObjectStatus
type ObjectStatus string

func NewObjectStatus(value ObjectStatus) *ObjectStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated ObjectStatus.
func (m ObjectStatus) Pointer() *ObjectStatus {
	return &m
}

const (

	// ObjectStatusSPOOL captures enum value "SPOOL"
	ObjectStatusSPOOL ObjectStatus = "SPOOL"

	// ObjectStatusREQUEST captures enum value "REQUEST"
	ObjectStatusREQUEST ObjectStatus = "REQUEST"

	// ObjectStatusPENDINGNOTIFY captures enum value "PENDING_NOTIFY"
	ObjectStatusPENDINGNOTIFY ObjectStatus = "PENDING_NOTIFY"

	// ObjectStatusPENDING captures enum value "PENDING"
	ObjectStatusPENDING ObjectStatus = "PENDING"

	// ObjectStatusFAILED captures enum value "FAILED"
	ObjectStatusFAILED ObjectStatus = "FAILED"

	// ObjectStatusSUCCESS captures enum value "SUCCESS"
	ObjectStatusSUCCESS ObjectStatus = "SUCCESS"

	// ObjectStatusPENDINGDOMAIN captures enum value "PENDING_DOMAIN"
	ObjectStatusPENDINGDOMAIN ObjectStatus = "PENDING_DOMAIN"

	// ObjectStatusCONNECT captures enum value "CONNECT"
	ObjectStatusCONNECT ObjectStatus = "CONNECT"
)

// for schema
var objectStatusEnum []interface{}

func init() {
	var res []ObjectStatus
	if err := json.Unmarshal([]byte(`["SPOOL","REQUEST","PENDING_NOTIFY","PENDING","FAILED","SUCCESS","PENDING_DOMAIN","CONNECT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		objectStatusEnum = append(objectStatusEnum, v)
	}
}

func (m ObjectStatus) validateObjectStatusEnum(path, location string, value ObjectStatus) error {
	if err := validate.EnumCase(path, location, value, objectStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this object status
func (m ObjectStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateObjectStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this object status based on context it is used
func (m ObjectStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

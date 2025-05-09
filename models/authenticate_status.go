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

// AuthenticateStatus authenticate status
//
// swagger:model AuthenticateStatus
type AuthenticateStatus string

func NewAuthenticateStatus(value AuthenticateStatus) *AuthenticateStatus {
	return &value
}

// Pointer returns a pointer to a freshly-allocated AuthenticateStatus.
func (m AuthenticateStatus) Pointer() *AuthenticateStatus {
	return &m
}

const (

	// AuthenticateStatusCOMPLETED captures enum value "COMPLETED"
	AuthenticateStatusCOMPLETED AuthenticateStatus = "COMPLETED"

	// AuthenticateStatusADDITIONALINFORMATIONREQUIRED captures enum value "ADDITIONAL_INFORMATION_REQUIRED"
	AuthenticateStatusADDITIONALINFORMATIONREQUIRED AuthenticateStatus = "ADDITIONAL_INFORMATION_REQUIRED"

	// AuthenticateStatusINPROGRESS captures enum value "IN_PROGRESS"
	AuthenticateStatusINPROGRESS AuthenticateStatus = "IN_PROGRESS"

	// AuthenticateStatusCOMMUNICATIONSENTTOCUSTOMER captures enum value "COMMUNICATION_SENT_TO_CUSTOMER"
	AuthenticateStatusCOMMUNICATIONSENTTOCUSTOMER AuthenticateStatus = "COMMUNICATION_SENT_TO_CUSTOMER"

	// AuthenticateStatusATTEMPTEDTOREACHCUSTOMER captures enum value "ATTEMPTED_TO_REACH_CUSTOMER"
	AuthenticateStatusATTEMPTEDTOREACHCUSTOMER AuthenticateStatus = "ATTEMPTED_TO_REACH_CUSTOMER"

	// AuthenticateStatusPENDINGCUSTOMERAPPROVAL captures enum value "PENDING_CUSTOMER_APPROVAL"
	AuthenticateStatusPENDINGCUSTOMERAPPROVAL AuthenticateStatus = "PENDING_CUSTOMER_APPROVAL"

	// AuthenticateStatusNOTCOMPLETED captures enum value "NOT_COMPLETED"
	AuthenticateStatusNOTCOMPLETED AuthenticateStatus = "NOT_COMPLETED"

	// AuthenticateStatusCOMPLETEDPASSED captures enum value "COMPLETED_PASSED"
	AuthenticateStatusCOMPLETEDPASSED AuthenticateStatus = "COMPLETED_PASSED"

	// AuthenticateStatusCOMPLETEDFAILED captures enum value "COMPLETED_FAILED"
	AuthenticateStatusCOMPLETEDFAILED AuthenticateStatus = "COMPLETED_FAILED"

	// AuthenticateStatusUNABLETOCOMPLETE captures enum value "UNABLE_TO_COMPLETE"
	AuthenticateStatusUNABLETOCOMPLETE AuthenticateStatus = "UNABLE_TO_COMPLETE"

	// AuthenticateStatusNOTSTARTED captures enum value "NOT_STARTED"
	AuthenticateStatusNOTSTARTED AuthenticateStatus = "NOT_STARTED"
)

// for schema
var authenticateStatusEnum []interface{}

func init() {
	var res []AuthenticateStatus
	if err := json.Unmarshal([]byte(`["COMPLETED","ADDITIONAL_INFORMATION_REQUIRED","IN_PROGRESS","COMMUNICATION_SENT_TO_CUSTOMER","ATTEMPTED_TO_REACH_CUSTOMER","PENDING_CUSTOMER_APPROVAL","NOT_COMPLETED","COMPLETED_PASSED","COMPLETED_FAILED","UNABLE_TO_COMPLETE","NOT_STARTED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		authenticateStatusEnum = append(authenticateStatusEnum, v)
	}
}

func (m AuthenticateStatus) validateAuthenticateStatusEnum(path, location string, value AuthenticateStatus) error {
	if err := validate.EnumCase(path, location, value, authenticateStatusEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this authenticate status
func (m AuthenticateStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAuthenticateStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this authenticate status based on context it is used
func (m AuthenticateStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

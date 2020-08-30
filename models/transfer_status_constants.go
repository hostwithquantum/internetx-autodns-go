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

// TransferStatusConstants transfer status constants
//
// swagger:model TransferStatusConstants
type TransferStatusConstants string

const (

	// TransferStatusConstantsNOTSET captures enum value "NOT_SET"
	TransferStatusConstantsNOTSET TransferStatusConstants = "NOT_SET"

	// TransferStatusConstantsSTART captures enum value "START"
	TransferStatusConstantsSTART TransferStatusConstants = "START"

	// TransferStatusConstantsFAILED captures enum value "FAILED"
	TransferStatusConstantsFAILED TransferStatusConstants = "FAILED"

	// TransferStatusConstantsNACK captures enum value "NACK"
	TransferStatusConstantsNACK TransferStatusConstants = "NACK"

	// TransferStatusConstantsACK captures enum value "ACK"
	TransferStatusConstantsACK TransferStatusConstants = "ACK"

	// TransferStatusConstantsFOA1SENT captures enum value "FOA1_SENT"
	TransferStatusConstantsFOA1SENT TransferStatusConstants = "FOA1_SENT"

	// TransferStatusConstantsAUTOUPDATESUCCESS captures enum value "AUTOUPDATE_SUCCESS"
	TransferStatusConstantsAUTOUPDATESUCCESS TransferStatusConstants = "AUTOUPDATE_SUCCESS"

	// TransferStatusConstantsAUTOUPDATEFAILED captures enum value "AUTOUPDATE_FAILED"
	TransferStatusConstantsAUTOUPDATEFAILED TransferStatusConstants = "AUTOUPDATE_FAILED"
)

// for schema
var transferStatusConstantsEnum []interface{}

func init() {
	var res []TransferStatusConstants
	if err := json.Unmarshal([]byte(`["NOT_SET","START","FAILED","NACK","ACK","FOA1_SENT","AUTOUPDATE_SUCCESS","AUTOUPDATE_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		transferStatusConstantsEnum = append(transferStatusConstantsEnum, v)
	}
}

func (m TransferStatusConstants) validateTransferStatusConstantsEnum(path, location string, value TransferStatusConstants) error {
	if err := validate.EnumCase(path, location, value, transferStatusConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this transfer status constants
func (m TransferStatusConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTransferStatusConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
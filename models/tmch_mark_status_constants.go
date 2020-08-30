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

// TmchMarkStatusConstants tmch mark status constants
//
// swagger:model TmchMarkStatusConstants
type TmchMarkStatusConstants string

const (

	// TmchMarkStatusConstantsOPEN captures enum value "OPEN"
	TmchMarkStatusConstantsOPEN TmchMarkStatusConstants = "OPEN"

	// TmchMarkStatusConstantsPAYMENT captures enum value "PAYMENT"
	TmchMarkStatusConstantsPAYMENT TmchMarkStatusConstants = "PAYMENT"

	// TmchMarkStatusConstantsPENDING captures enum value "PENDING"
	TmchMarkStatusConstantsPENDING TmchMarkStatusConstants = "PENDING"

	// TmchMarkStatusConstantsINCORRECT captures enum value "INCORRECT"
	TmchMarkStatusConstantsINCORRECT TmchMarkStatusConstants = "INCORRECT"

	// TmchMarkStatusConstantsVERIFIED captures enum value "VERIFIED"
	TmchMarkStatusConstantsVERIFIED TmchMarkStatusConstants = "VERIFIED"

	// TmchMarkStatusConstantsSENT captures enum value "SENT"
	TmchMarkStatusConstantsSENT TmchMarkStatusConstants = "SENT"

	// TmchMarkStatusConstantsACTIVE captures enum value "ACTIVE"
	TmchMarkStatusConstantsACTIVE TmchMarkStatusConstants = "ACTIVE"

	// TmchMarkStatusConstantsDEACTIVATED captures enum value "DEACTIVATED"
	TmchMarkStatusConstantsDEACTIVATED TmchMarkStatusConstants = "DEACTIVATED"

	// TmchMarkStatusConstantsCANCELED captures enum value "CANCELED"
	TmchMarkStatusConstantsCANCELED TmchMarkStatusConstants = "CANCELED"

	// TmchMarkStatusConstantsRENEW captures enum value "RENEW"
	TmchMarkStatusConstantsRENEW TmchMarkStatusConstants = "RENEW"

	// TmchMarkStatusConstantsFAILED captures enum value "FAILED"
	TmchMarkStatusConstantsFAILED TmchMarkStatusConstants = "FAILED"

	// TmchMarkStatusConstantsEXTERNAL captures enum value "EXTERNAL"
	TmchMarkStatusConstantsEXTERNAL TmchMarkStatusConstants = "EXTERNAL"

	// TmchMarkStatusConstantsPENDINGTRANSFER captures enum value "PENDING_TRANSFER"
	TmchMarkStatusConstantsPENDINGTRANSFER TmchMarkStatusConstants = "PENDING_TRANSFER"

	// TmchMarkStatusConstantsPENDINGRENEW captures enum value "PENDING_RENEW"
	TmchMarkStatusConstantsPENDINGRENEW TmchMarkStatusConstants = "PENDING_RENEW"

	// TmchMarkStatusConstantsPENDINGDELETE captures enum value "PENDING_DELETE"
	TmchMarkStatusConstantsPENDINGDELETE TmchMarkStatusConstants = "PENDING_DELETE"

	// TmchMarkStatusConstantsPENDINGPAYMENT captures enum value "PENDING_PAYMENT"
	TmchMarkStatusConstantsPENDINGPAYMENT TmchMarkStatusConstants = "PENDING_PAYMENT"

	// TmchMarkStatusConstantsPENDINGTRANSFERPAYMENT captures enum value "PENDING_TRANSFER_PAYMENT"
	TmchMarkStatusConstantsPENDINGTRANSFERPAYMENT TmchMarkStatusConstants = "PENDING_TRANSFER_PAYMENT"

	// TmchMarkStatusConstantsPENDINGRENEWPAYMENT captures enum value "PENDING_RENEW_PAYMENT"
	TmchMarkStatusConstantsPENDINGRENEWPAYMENT TmchMarkStatusConstants = "PENDING_RENEW_PAYMENT"
)

// for schema
var tmchMarkStatusConstantsEnum []interface{}

func init() {
	var res []TmchMarkStatusConstants
	if err := json.Unmarshal([]byte(`["OPEN","PAYMENT","PENDING","INCORRECT","VERIFIED","SENT","ACTIVE","DEACTIVATED","CANCELED","RENEW","FAILED","EXTERNAL","PENDING_TRANSFER","PENDING_RENEW","PENDING_DELETE","PENDING_PAYMENT","PENDING_TRANSFER_PAYMENT","PENDING_RENEW_PAYMENT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tmchMarkStatusConstantsEnum = append(tmchMarkStatusConstantsEnum, v)
	}
}

func (m TmchMarkStatusConstants) validateTmchMarkStatusConstantsEnum(path, location string, value TmchMarkStatusConstants) error {
	if err := validate.EnumCase(path, location, value, tmchMarkStatusConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this tmch mark status constants
func (m TmchMarkStatusConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateTmchMarkStatusConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
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

// ModifierConstants modifier constants
//
// swagger:model ModifierConstants
type ModifierConstants string

const (

	// ModifierConstantsTTL captures enum value "TTL"
	ModifierConstantsTTL ModifierConstants = "TTL"

	// ModifierConstantsMX captures enum value "MX"
	ModifierConstantsMX ModifierConstants = "MX"

	// ModifierConstantsA captures enum value "A"
	ModifierConstantsA ModifierConstants = "A"

	// ModifierConstantsAAAA captures enum value "AAAA"
	ModifierConstantsAAAA ModifierConstants = "AAAA"

	// ModifierConstantsSOAEMAIL captures enum value "SOA_EMAIL"
	ModifierConstantsSOAEMAIL ModifierConstants = "SOA_EMAIL"

	// ModifierConstantsNSERVER captures enum value "NSERVER"
	ModifierConstantsNSERVER ModifierConstants = "NSERVER"

	// ModifierConstantsCNAME captures enum value "CNAME"
	ModifierConstantsCNAME ModifierConstants = "CNAME"

	// ModifierConstantsTXT captures enum value "TXT"
	ModifierConstantsTXT ModifierConstants = "TXT"

	// ModifierConstantsALL captures enum value "ALL"
	ModifierConstantsALL ModifierConstants = "ALL"

	// ModifierConstantsMAINIP captures enum value "MAIN_IP"
	ModifierConstantsMAINIP ModifierConstants = "MAIN_IP"

	// ModifierConstantsOWNERC captures enum value "OWNERC"
	ModifierConstantsOWNERC ModifierConstants = "OWNERC"

	// ModifierConstantsADMINC captures enum value "ADMINC"
	ModifierConstantsADMINC ModifierConstants = "ADMINC"

	// ModifierConstantsTECHC captures enum value "TECHC"
	ModifierConstantsTECHC ModifierConstants = "TECHC"

	// ModifierConstantsZONEC captures enum value "ZONEC"
	ModifierConstantsZONEC ModifierConstants = "ZONEC"

	// ModifierConstantsBILLINGC captures enum value "BILLINGC"
	ModifierConstantsBILLINGC ModifierConstants = "BILLINGC"

	// ModifierConstantsALIAS captures enum value "ALIAS"
	ModifierConstantsALIAS ModifierConstants = "ALIAS"
)

// for schema
var modifierConstantsEnum []interface{}

func init() {
	var res []ModifierConstants
	if err := json.Unmarshal([]byte(`["TTL","MX","A","AAAA","SOA_EMAIL","NSERVER","CNAME","TXT","ALL","MAIN_IP","OWNERC","ADMINC","TECHC","ZONEC","BILLINGC","ALIAS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		modifierConstantsEnum = append(modifierConstantsEnum, v)
	}
}

func (m ModifierConstants) validateModifierConstantsEnum(path, location string, value ModifierConstants) error {
	if err := validate.EnumCase(path, location, value, modifierConstantsEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this modifier constants
func (m ModifierConstants) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateModifierConstantsEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

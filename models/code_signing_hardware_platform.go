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

// CodeSigningHardwarePlatform code signing hardware platform
//
// swagger:model CodeSigningHardwarePlatform
type CodeSigningHardwarePlatform string

const (

	// CodeSigningHardwarePlatformAEPKEYPER captures enum value "AEP_KEYPER"
	CodeSigningHardwarePlatformAEPKEYPER CodeSigningHardwarePlatform = "AEP_KEYPER"

	// CodeSigningHardwarePlatformARXPRIVATESERVER captures enum value "ARX_PRIVATE_SERVER"
	CodeSigningHardwarePlatformARXPRIVATESERVER CodeSigningHardwarePlatform = "ARX_PRIVATE_SERVER"

	// CodeSigningHardwarePlatformBULLTRUSTWAYCRYPTOPCI captures enum value "BULL_TRUSTWAY_CRYPTO_PCI"
	CodeSigningHardwarePlatformBULLTRUSTWAYCRYPTOPCI CodeSigningHardwarePlatform = "BULL_TRUSTWAY_CRYPTO_PCI"

	// CodeSigningHardwarePlatformEPASS3003 captures enum value "E_PASS_3003"
	CodeSigningHardwarePlatformEPASS3003 CodeSigningHardwarePlatform = "E_PASS_3003"

	// CodeSigningHardwarePlatformSAFENETETOKEN5100 captures enum value "SAFE_NET_E_TOKEN_5100"
	CodeSigningHardwarePlatformSAFENETETOKEN5100 CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5100"

	// CodeSigningHardwarePlatformSAFENETETOKEN5105 captures enum value "SAFE_NET_E_TOKEN_5105"
	CodeSigningHardwarePlatformSAFENETETOKEN5105 CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5105"

	// CodeSigningHardwarePlatformSAFENETETOKEN5110 captures enum value "SAFE_NET_E_TOKEN_5110"
	CodeSigningHardwarePlatformSAFENETETOKEN5110 CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5110"

	// CodeSigningHardwarePlatformSAFENETETOKEN5110FIPS captures enum value "SAFE_NET_E_TOKEN_5110_FIPS"
	CodeSigningHardwarePlatformSAFENETETOKEN5110FIPS CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5110_FIPS"

	// CodeSigningHardwarePlatformSAFENETETOKEN5200 captures enum value "SAFE_NET_E_TOKEN_5200"
	CodeSigningHardwarePlatformSAFENETETOKEN5200 CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5200"

	// CodeSigningHardwarePlatformSAFENETETOKEN5205 captures enum value "SAFE_NET_E_TOKEN_5205"
	CodeSigningHardwarePlatformSAFENETETOKEN5205 CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_5205"

	// CodeSigningHardwarePlatformSAFENETETOKENPRO72K captures enum value "SAFE_NET_E_TOKEN_PRO_72K"
	CodeSigningHardwarePlatformSAFENETETOKENPRO72K CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_PRO_72K"

	// CodeSigningHardwarePlatformSAFENETETOKENPROANYWHERE captures enum value "SAFE_NET_E_TOKEN_PRO_ANYWHERE"
	CodeSigningHardwarePlatformSAFENETETOKENPROANYWHERE CodeSigningHardwarePlatform = "SAFE_NET_E_TOKEN_PRO_ANYWHERE"

	// CodeSigningHardwarePlatformSAFENETIKEY4000 captures enum value "SAFE_NET_I_KEY_4000"
	CodeSigningHardwarePlatformSAFENETIKEY4000 CodeSigningHardwarePlatform = "SAFE_NET_I_KEY_4000"

	// CodeSigningHardwarePlatformSAFENETLUNA captures enum value "SAFE_NET_LUNA"
	CodeSigningHardwarePlatformSAFENETLUNA CodeSigningHardwarePlatform = "SAFE_NET_LUNA"

	// CodeSigningHardwarePlatformTHALESNSHIELD captures enum value "THALES_N_SHIELD"
	CodeSigningHardwarePlatformTHALESNSHIELD CodeSigningHardwarePlatform = "THALES_N_SHIELD"

	// CodeSigningHardwarePlatformULTIMACOCRYPTOSERVER captures enum value "ULTIMACO_CRYPTO_SERVER"
	CodeSigningHardwarePlatformULTIMACOCRYPTOSERVER CodeSigningHardwarePlatform = "ULTIMACO_CRYPTO_SERVER"

	// CodeSigningHardwarePlatformOTHER captures enum value "OTHER"
	CodeSigningHardwarePlatformOTHER CodeSigningHardwarePlatform = "OTHER"
)

// for schema
var codeSigningHardwarePlatformEnum []interface{}

func init() {
	var res []CodeSigningHardwarePlatform
	if err := json.Unmarshal([]byte(`["AEP_KEYPER","ARX_PRIVATE_SERVER","BULL_TRUSTWAY_CRYPTO_PCI","E_PASS_3003","SAFE_NET_E_TOKEN_5100","SAFE_NET_E_TOKEN_5105","SAFE_NET_E_TOKEN_5110","SAFE_NET_E_TOKEN_5110_FIPS","SAFE_NET_E_TOKEN_5200","SAFE_NET_E_TOKEN_5205","SAFE_NET_E_TOKEN_PRO_72K","SAFE_NET_E_TOKEN_PRO_ANYWHERE","SAFE_NET_I_KEY_4000","SAFE_NET_LUNA","THALES_N_SHIELD","ULTIMACO_CRYPTO_SERVER","OTHER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		codeSigningHardwarePlatformEnum = append(codeSigningHardwarePlatformEnum, v)
	}
}

func (m CodeSigningHardwarePlatform) validateCodeSigningHardwarePlatformEnum(path, location string, value CodeSigningHardwarePlatform) error {
	if err := validate.EnumCase(path, location, value, codeSigningHardwarePlatformEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this code signing hardware platform
func (m CodeSigningHardwarePlatform) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateCodeSigningHardwarePlatformEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

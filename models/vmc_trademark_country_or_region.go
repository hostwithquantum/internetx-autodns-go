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

// VmcTrademarkCountryOrRegion vmc trademark country or region
//
// swagger:model VmcTrademarkCountryOrRegion
type VmcTrademarkCountryOrRegion string

func NewVmcTrademarkCountryOrRegion(value VmcTrademarkCountryOrRegion) *VmcTrademarkCountryOrRegion {
	return &value
}

// Pointer returns a pointer to a freshly-allocated VmcTrademarkCountryOrRegion.
func (m VmcTrademarkCountryOrRegion) Pointer() *VmcTrademarkCountryOrRegion {
	return &m
}

const (

	// VmcTrademarkCountryOrRegionAU captures enum value "AU"
	VmcTrademarkCountryOrRegionAU VmcTrademarkCountryOrRegion = "AU"

	// VmcTrademarkCountryOrRegionBR captures enum value "BR"
	VmcTrademarkCountryOrRegionBR VmcTrademarkCountryOrRegion = "BR"

	// VmcTrademarkCountryOrRegionBX captures enum value "BX"
	VmcTrademarkCountryOrRegionBX VmcTrademarkCountryOrRegion = "BX"

	// VmcTrademarkCountryOrRegionCA captures enum value "CA"
	VmcTrademarkCountryOrRegionCA VmcTrademarkCountryOrRegion = "CA"

	// VmcTrademarkCountryOrRegionCH captures enum value "CH"
	VmcTrademarkCountryOrRegionCH VmcTrademarkCountryOrRegion = "CH"

	// VmcTrademarkCountryOrRegionDE captures enum value "DE"
	VmcTrademarkCountryOrRegionDE VmcTrademarkCountryOrRegion = "DE"

	// VmcTrademarkCountryOrRegionDK captures enum value "DK"
	VmcTrademarkCountryOrRegionDK VmcTrademarkCountryOrRegion = "DK"

	// VmcTrademarkCountryOrRegionEM captures enum value "EM"
	VmcTrademarkCountryOrRegionEM VmcTrademarkCountryOrRegion = "EM"

	// VmcTrademarkCountryOrRegionES captures enum value "ES"
	VmcTrademarkCountryOrRegionES VmcTrademarkCountryOrRegion = "ES"

	// VmcTrademarkCountryOrRegionFR captures enum value "FR"
	VmcTrademarkCountryOrRegionFR VmcTrademarkCountryOrRegion = "FR"

	// VmcTrademarkCountryOrRegionGB captures enum value "GB"
	VmcTrademarkCountryOrRegionGB VmcTrademarkCountryOrRegion = "GB"

	// VmcTrademarkCountryOrRegionIN captures enum value "IN"
	VmcTrademarkCountryOrRegionIN VmcTrademarkCountryOrRegion = "IN"

	// VmcTrademarkCountryOrRegionJP captures enum value "JP"
	VmcTrademarkCountryOrRegionJP VmcTrademarkCountryOrRegion = "JP"

	// VmcTrademarkCountryOrRegionKR captures enum value "KR"
	VmcTrademarkCountryOrRegionKR VmcTrademarkCountryOrRegion = "KR"

	// VmcTrademarkCountryOrRegionNZ captures enum value "NZ"
	VmcTrademarkCountryOrRegionNZ VmcTrademarkCountryOrRegion = "NZ"

	// VmcTrademarkCountryOrRegionSE captures enum value "SE"
	VmcTrademarkCountryOrRegionSE VmcTrademarkCountryOrRegion = "SE"

	// VmcTrademarkCountryOrRegionUS captures enum value "US"
	VmcTrademarkCountryOrRegionUS VmcTrademarkCountryOrRegion = "US"
)

// for schema
var vmcTrademarkCountryOrRegionEnum []interface{}

func init() {
	var res []VmcTrademarkCountryOrRegion
	if err := json.Unmarshal([]byte(`["AU","BR","BX","CA","CH","DE","DK","EM","ES","FR","GB","IN","JP","KR","NZ","SE","US"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		vmcTrademarkCountryOrRegionEnum = append(vmcTrademarkCountryOrRegionEnum, v)
	}
}

func (m VmcTrademarkCountryOrRegion) validateVmcTrademarkCountryOrRegionEnum(path, location string, value VmcTrademarkCountryOrRegion) error {
	if err := validate.EnumCase(path, location, value, vmcTrademarkCountryOrRegionEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this vmc trademark country or region
func (m VmcTrademarkCountryOrRegion) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateVmcTrademarkCountryOrRegionEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this vmc trademark country or region based on context it is used
func (m VmcTrademarkCountryOrRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TmchMarkAddon tmch mark addon
//
// swagger:model TmchMarkAddon
type TmchMarkAddon struct {

	// court name
	CourtName string `json:"courtName,omitempty"`

	// court protection country
	CourtProtectionCountry string `json:"courtProtectionCountry,omitempty"`

	// court reference
	CourtReference string `json:"courtReference,omitempty"`

	// parent reference
	ParentReference string `json:"parentReference,omitempty"`

	// protection
	// Format: date-time
	Protection strfmt.DateTime `json:"protection,omitempty"`

	// trademark classification
	TrademarkClassification string `json:"trademarkClassification,omitempty"`

	// trademark expire
	// Format: date-time
	TrademarkExpire strfmt.DateTime `json:"trademarkExpire,omitempty"`

	// trademark jurisdiction
	TrademarkJurisdiction string `json:"trademarkJurisdiction,omitempty"`

	// trademark number
	TrademarkNumber string `json:"trademarkNumber,omitempty"`

	// trademark registration
	// Format: date-time
	TrademarkRegistration strfmt.DateTime `json:"trademarkRegistration,omitempty"`

	// treaty execution
	// Format: date-time
	TreatyExecution strfmt.DateTime `json:"treatyExecution,omitempty"`

	// treaty title
	TreatyTitle string `json:"treatyTitle,omitempty"`
}

// Validate validates this tmch mark addon
func (m *TmchMarkAddon) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrademarkExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTrademarkRegistration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTreatyExecution(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TmchMarkAddon) validateProtection(formats strfmt.Registry) error {
	if swag.IsZero(m.Protection) { // not required
		return nil
	}

	if err := validate.FormatOf("protection", "body", "date-time", m.Protection.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TmchMarkAddon) validateTrademarkExpire(formats strfmt.Registry) error {
	if swag.IsZero(m.TrademarkExpire) { // not required
		return nil
	}

	if err := validate.FormatOf("trademarkExpire", "body", "date-time", m.TrademarkExpire.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TmchMarkAddon) validateTrademarkRegistration(formats strfmt.Registry) error {
	if swag.IsZero(m.TrademarkRegistration) { // not required
		return nil
	}

	if err := validate.FormatOf("trademarkRegistration", "body", "date-time", m.TrademarkRegistration.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TmchMarkAddon) validateTreatyExecution(formats strfmt.Registry) error {
	if swag.IsZero(m.TreatyExecution) { // not required
		return nil
	}

	if err := validate.FormatOf("treatyExecution", "body", "date-time", m.TreatyExecution.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this tmch mark addon based on context it is used
func (m *TmchMarkAddon) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TmchMarkAddon) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TmchMarkAddon) UnmarshalBinary(b []byte) error {
	var res TmchMarkAddon
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

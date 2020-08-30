// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ContactExtensions contact extensions
//
// swagger:model ContactExtensions
type ContactExtensions struct {

	// The .aero contact extensions.
	Aero *ContactAeroExtensions `json:"aero,omitempty"`

	// The .au contact extensions.
	Au *ContactAuExtensions `json:"au,omitempty"`

	// The .bank contact extensions.
	Bank *ContactBankExtensions `json:"bank,omitempty"`

	// The .barcelona contact extensions.
	Barcelona *ContactBarcelonaExtensions `json:"barcelona,omitempty"`

	// The .ca contact extensions.
	Ca *ContactCaExtensions `json:"ca,omitempty"`

	// The .cat contact extensions.
	Cat *ContactCatExtensions `json:"cat,omitempty"`

	// The general contact extensions.
	General *ContactGeneralExtensions `json:"general,omitempty"`

	// The .hk contact extensions.
	Hk *ContactHkExtensions `json:"hk,omitempty"`

	// The .it contact extensions.
	It *ContactItExtensions `json:"it,omitempty"`

	// The .jobs contact extensions.
	Jobs *ContactJobsExtensions `json:"jobs,omitempty"`

	// The .jp contact extensions.
	Jp *ContactJpExtensions `json:"jp,omitempty"`

	// The .luxe contact extensions.
	Luxe *ContactLuxeExtensions `json:"luxe,omitempty"`

	// The .madrid contact extensions.
	Madrid *ContactMadridExtensions `json:"madrid,omitempty"`

	// The .nz contact extensions
	Nz *ContactNzExtensions `json:"nz,omitempty"`

	// The .ro contact extensions.
	Ro *ContactRoExtensions `json:"ro,omitempty"`

	// The .ru contact extensions.
	Ru *ContactRuExtensions `json:"ru,omitempty"`

	// The .sport contact extensions.
	Sport *ContactSportExtensions `json:"sport,omitempty"`

	// The .swiss contact extensions.
	Swiss *ContactSwissExtensions `json:"swiss,omitempty"`

	// The .uk contact extensions.
	Uk *ContactUkExtensions `json:"uk,omitempty"`

	// The .xxx contact extensions.
	Xxx *ContactXxxExtensions `json:"xxx,omitempty"`
}

// Validate validates this contact extensions
func (m *ContactExtensions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAero(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBank(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBarcelona(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCat(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeneral(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLuxe(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMadrid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNz(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRu(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwiss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUk(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXxx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContactExtensions) validateAero(formats strfmt.Registry) error {

	if swag.IsZero(m.Aero) { // not required
		return nil
	}

	if m.Aero != nil {
		if err := m.Aero.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aero")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateAu(formats strfmt.Registry) error {

	if swag.IsZero(m.Au) { // not required
		return nil
	}

	if m.Au != nil {
		if err := m.Au.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("au")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateBank(formats strfmt.Registry) error {

	if swag.IsZero(m.Bank) { // not required
		return nil
	}

	if m.Bank != nil {
		if err := m.Bank.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bank")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateBarcelona(formats strfmt.Registry) error {

	if swag.IsZero(m.Barcelona) { // not required
		return nil
	}

	if m.Barcelona != nil {
		if err := m.Barcelona.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("barcelona")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateCa(formats strfmt.Registry) error {

	if swag.IsZero(m.Ca) { // not required
		return nil
	}

	if m.Ca != nil {
		if err := m.Ca.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ca")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateCat(formats strfmt.Registry) error {

	if swag.IsZero(m.Cat) { // not required
		return nil
	}

	if m.Cat != nil {
		if err := m.Cat.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cat")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateGeneral(formats strfmt.Registry) error {

	if swag.IsZero(m.General) { // not required
		return nil
	}

	if m.General != nil {
		if err := m.General.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("general")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateHk(formats strfmt.Registry) error {

	if swag.IsZero(m.Hk) { // not required
		return nil
	}

	if m.Hk != nil {
		if err := m.Hk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hk")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateIt(formats strfmt.Registry) error {

	if swag.IsZero(m.It) { // not required
		return nil
	}

	if m.It != nil {
		if err := m.It.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("it")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateJobs(formats strfmt.Registry) error {

	if swag.IsZero(m.Jobs) { // not required
		return nil
	}

	if m.Jobs != nil {
		if err := m.Jobs.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jobs")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateJp(formats strfmt.Registry) error {

	if swag.IsZero(m.Jp) { // not required
		return nil
	}

	if m.Jp != nil {
		if err := m.Jp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("jp")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateLuxe(formats strfmt.Registry) error {

	if swag.IsZero(m.Luxe) { // not required
		return nil
	}

	if m.Luxe != nil {
		if err := m.Luxe.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("luxe")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateMadrid(formats strfmt.Registry) error {

	if swag.IsZero(m.Madrid) { // not required
		return nil
	}

	if m.Madrid != nil {
		if err := m.Madrid.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("madrid")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateNz(formats strfmt.Registry) error {

	if swag.IsZero(m.Nz) { // not required
		return nil
	}

	if m.Nz != nil {
		if err := m.Nz.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nz")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateRo(formats strfmt.Registry) error {

	if swag.IsZero(m.Ro) { // not required
		return nil
	}

	if m.Ro != nil {
		if err := m.Ro.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ro")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateRu(formats strfmt.Registry) error {

	if swag.IsZero(m.Ru) { // not required
		return nil
	}

	if m.Ru != nil {
		if err := m.Ru.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ru")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateSport(formats strfmt.Registry) error {

	if swag.IsZero(m.Sport) { // not required
		return nil
	}

	if m.Sport != nil {
		if err := m.Sport.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sport")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateSwiss(formats strfmt.Registry) error {

	if swag.IsZero(m.Swiss) { // not required
		return nil
	}

	if m.Swiss != nil {
		if err := m.Swiss.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swiss")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateUk(formats strfmt.Registry) error {

	if swag.IsZero(m.Uk) { // not required
		return nil
	}

	if m.Uk != nil {
		if err := m.Uk.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("uk")
			}
			return err
		}
	}

	return nil
}

func (m *ContactExtensions) validateXxx(formats strfmt.Registry) error {

	if swag.IsZero(m.Xxx) { // not required
		return nil
	}

	if m.Xxx != nil {
		if err := m.Xxx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("xxx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContactExtensions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContactExtensions) UnmarshalBinary(b []byte) error {
	var res ContactExtensions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

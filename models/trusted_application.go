// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TrustedApplication trusted application
//
// swagger:model TrustedApplication
type TrustedApplication struct {

	// The kind of application.
	// Required: true
	Application *Application `json:"application"`

	// A comment for the trusted application.
	Comment string `json:"comment,omitempty"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The device, which uses the trusted app.
	Device string `json:"device,omitempty"`

	// The wrapper of the allowed functions for the trusted application.
	Functions *TrustedApplicationFunctions `json:"functions,omitempty"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// The password for the trusted application.
	Password string `json:"password,omitempty"`

	// The wrapper of the ip restrictions for the trusted application.
	Restrictions *IPRestrictions `json:"restrictions,omitempty"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating user of the object.
	Updater *BasicUser `json:"updater,omitempty"`

	// The user the trusted application belongs to.
	User *BasicUser `json:"user,omitempty"`

	// The unique identifier for the trusted application.
	// Required: true
	UUID *string `json:"uuid"`
}

// Validate validates this trusted application
func (m *TrustedApplication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApplication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFunctions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdater(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TrustedApplication) validateApplication(formats strfmt.Registry) error {

	if err := validate.Required("application", "body", m.Application); err != nil {
		return err
	}

	if m.Application != nil {
		if err := m.Application.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("application")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrustedApplication) validateFunctions(formats strfmt.Registry) error {

	if swag.IsZero(m.Functions) { // not required
		return nil
	}

	if m.Functions != nil {
		if err := m.Functions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("functions")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.Restrictions) { // not required
		return nil
	}

	if m.Restrictions != nil {
		if err := m.Restrictions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restrictions")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TrustedApplication) validateUpdater(formats strfmt.Registry) error {

	if swag.IsZero(m.Updater) { // not required
		return nil
	}

	if m.Updater != nil {
		if err := m.Updater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateUser(formats strfmt.Registry) error {

	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

func (m *TrustedApplication) validateUUID(formats strfmt.Registry) error {

	if err := validate.Required("uuid", "body", m.UUID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TrustedApplication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TrustedApplication) UnmarshalBinary(b []byte) error {
	var res TrustedApplication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

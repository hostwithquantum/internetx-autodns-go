// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Subject subject
//
// swagger:model Subject
type Subject struct {

	// List of metadata
	// Unique: true
	Customs []*Custom `json:"customs"`

	// The name of the booking item, e.g. domain.de
	// Required: true
	Name *string `json:"name"`

	// The internal unique name of the object, e.g. xn--
	Object string `json:"object,omitempty"`
}

// Validate validates this subject
func (m *Subject) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustoms(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Subject) validateCustoms(formats strfmt.Registry) error {

	if swag.IsZero(m.Customs) { // not required
		return nil
	}

	if err := validate.UniqueItems("customs", "body", m.Customs); err != nil {
		return err
	}

	for i := 0; i < len(m.Customs); i++ {
		if swag.IsZero(m.Customs[i]) { // not required
			continue
		}

		if m.Customs[i] != nil {
			if err := m.Customs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Subject) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Subject) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Subject) UnmarshalBinary(b []byte) error {
	var res Subject
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
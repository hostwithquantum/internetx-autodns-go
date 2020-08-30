// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ServiceProfiles service profiles
//
// swagger:model ServiceProfiles
type ServiceProfiles struct {

	// The service profiles of the user.
	ServiceProfiles []*ServiceUsersProfile `json:"serviceProfiles"`
}

// Validate validates this service profiles
func (m *ServiceProfiles) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServiceProfiles(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ServiceProfiles) validateServiceProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceProfiles) { // not required
		return nil
	}

	for i := 0; i < len(m.ServiceProfiles); i++ {
		if swag.IsZero(m.ServiceProfiles[i]) { // not required
			continue
		}

		if m.ServiceProfiles[i] != nil {
			if err := m.ServiceProfiles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("serviceProfiles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ServiceProfiles) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ServiceProfiles) UnmarshalBinary(b []byte) error {
	var res ServiceProfiles
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

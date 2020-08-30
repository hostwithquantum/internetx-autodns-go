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

// DomainStudioSourceRecommended domain studio source recommended
//
// swagger:model DomainStudioSourceRecommended
type DomainStudioSourceRecommended struct {

	// The maximum amount of generated similar domains.
	Max int32 `json:"max,omitempty"`

	// Defines whether to return only free domain names when service WHOIS is used for a source.
	OnlyAvailable bool `json:"onlyAvailable,omitempty"`

	// The services to fetch extra data from for this source
	Services []DomainEnvelopeSearchService `json:"services"`
}

// Validate validates this domain studio source recommended
func (m *DomainStudioSourceRecommended) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSourceRecommended) validateServices(formats strfmt.Registry) error {

	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {

		if err := m.Services[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainStudioSourceRecommended) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainStudioSourceRecommended) UnmarshalBinary(b []byte) error {
	var res DomainStudioSourceRecommended
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

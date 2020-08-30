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

// DomainStudioSourceSuggestion domain studio source suggestion
//
// swagger:model DomainStudioSourceSuggestion
type DomainStudioSourceSuggestion struct {

	// Suggestion language
	Language string `json:"language,omitempty"`

	// Maximum fetched suggested domains
	Max int32 `json:"max,omitempty"`

	// Maximum sld length for suggested domains
	MaxSldLength int32 `json:"maxSldLength,omitempty"`

	// Defines whether to return only free domain names when service WHOIS is used for a source.
	OnlyAvailable bool `json:"onlyAvailable,omitempty"`

	// The services to fetch extra data from for this source
	Services []DomainEnvelopeSearchService `json:"services"`

	// Selected tlds
	Tlds []string `json:"tlds"`

	// Suggested domains with dash
	UseDash bool `json:"useDash,omitempty"`

	// Suggested domains with idn
	UseIdn bool `json:"useIdn,omitempty"`

	// Suggested domains with numbers
	UseNumber bool `json:"useNumber,omitempty"`
}

// Validate validates this domain studio source suggestion
func (m *DomainStudioSourceSuggestion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSourceSuggestion) validateServices(formats strfmt.Registry) error {

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
func (m *DomainStudioSourceSuggestion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainStudioSourceSuggestion) UnmarshalBinary(b []byte) error {
	var res DomainStudioSourceSuggestion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

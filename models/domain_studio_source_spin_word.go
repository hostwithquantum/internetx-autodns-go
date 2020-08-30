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

// DomainStudioSourceSpinWord domain studio source spin word
//
// swagger:model DomainStudioSourceSpinWord
type DomainStudioSourceSpinWord struct {

	// Maximum fetched suggested domains
	Max int32 `json:"max,omitempty"`

	// Maximum sld length for suggested domains
	MaxSldLength int32 `json:"maxSldLength,omitempty"`

	// Defines whether to return only free domain names when service WHOIS is used for a source.
	OnlyAvailable bool `json:"onlyAvailable,omitempty"`

	// Position of the word for selective wordsuggestion starting with “0” for the firstword. If no position is supplied, allpositions are spinned.
	Position int32 `json:"position,omitempty"`

	// The services to fetch extra data from for this source
	Services []DomainEnvelopeSearchService `json:"services"`

	// Double with a value between 0.00 and 1.00.The higher this value, the closer thealternate word will be to the original word.
	Similarity float64 `json:"similarity,omitempty"`

	// Selected tlds
	Tlds []string `json:"tlds"`

	// Suggested domains with idn
	UseIdn bool `json:"useIdn,omitempty"`
}

// Validate validates this domain studio source spin word
func (m *DomainStudioSourceSpinWord) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSourceSpinWord) validateServices(formats strfmt.Registry) error {

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
func (m *DomainStudioSourceSpinWord) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainStudioSourceSpinWord) UnmarshalBinary(b []byte) error {
	var res DomainStudioSourceSpinWord
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

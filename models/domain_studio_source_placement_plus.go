// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DomainStudioSourcePlacementPlus domain studio source placement plus
//
// swagger:model DomainStudioSourcePlacementPlus
type DomainStudioSourcePlacementPlus struct {

	// Exclude adult related and possibly offensive domains
	Adult bool `json:"adult,omitempty"`

	// The acronym for AccountIdentifier
	Aid string `json:"aid,omitempty"`

	// Override the default selection of generic TLDs and their base score. TLD and their custom score must be separated by the column sign “:”. Example org:1,club:0.5,xyz:0.0
	DefScore []string `json:"defScore"`

	// The generated domains of this source
	Domains []string `json:"domains"`

	// Defines the type of recommendation. Accepted values: 0 (default),1,2,3 and 4. 0 allows both TLD and Domain Recommendation.
	// 1 allows only TLD Recommendations to be displayed.
	// 2 allows only Domain Recommendations to be displayed
	// 3 allows only TLD Recommendations to be displayed plus simple domain variations
	// allows both TLD and Domain Recommendation (like func=0) BUT only returns Domains
	// Recommendations that are either a top legacy TLD such as com, net or org or any TLD IF the confidence score is higher than a minimum threshold.
	Func string `json:"func,omitempty"`

	// Suggestion language
	Language string `json:"language,omitempty"`

	// Maximum fetched suggested domains
	Max int32 `json:"max,omitempty"`

	// filter our domains whose confidence score is lower than your threshold. This parameters accepts values between 0.0 and 1.0. The default value is set to 0.0.
	MinConfidence float32 `json:"minConfidence,omitempty"`

	// Defines whether to return only free domain names when service WHOIS is used for a source.
	OnlyAvailable bool `json:"onlyAvailable,omitempty"`

	// The acronym for ResellerIdentifier
	Rid string `json:"rid,omitempty"`

	// The services to fetch extra data from for this source
	Services []DomainEnvelopeSearchService `json:"services"`

	// The acronym for Session Identifier
	Sid string `json:"sid,omitempty"`

	// Selected tlds
	// Read Only: true
	Tlds []string `json:"tlds"`

	// Suggest domains with dash
	UseDash bool `json:"useDash,omitempty"`

	// Suggest domains with idn
	UseIdn bool `json:"useIdn,omitempty"`
}

// Validate validates this domain studio source placement plus
func (m *DomainStudioSourcePlacementPlus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateServices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSourcePlacementPlus) validateServices(formats strfmt.Registry) error {
	if swag.IsZero(m.Services) { // not required
		return nil
	}

	for i := 0; i < len(m.Services); i++ {

		if err := m.Services[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("services" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// ContextValidate validate this domain studio source placement plus based on the context it is used
func (m *DomainStudioSourcePlacementPlus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateServices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTlds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSourcePlacementPlus) contextValidateServices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Services); i++ {

		if swag.IsZero(m.Services[i]) { // not required
			return nil
		}

		if err := m.Services[i].ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("services" + "." + strconv.Itoa(i))
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("services" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

func (m *DomainStudioSourcePlacementPlus) contextValidateTlds(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "tlds", "body", []string(m.Tlds)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainStudioSourcePlacementPlus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainStudioSourcePlacementPlus) UnmarshalBinary(b []byte) error {
	var res DomainStudioSourcePlacementPlus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

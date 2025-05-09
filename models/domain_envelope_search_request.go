// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainEnvelopeSearchRequest domain envelope search request
//
// swagger:model DomainEnvelopeSearchRequest
type DomainEnvelopeSearchRequest struct {

	// The user agent of the request. Only readable!
	Agent string `json:"agent,omitempty"`

	// Allow duplicate domain names from different sources.
	AllowDuplicates bool `json:"allowDuplicates,omitempty"`

	// Activates the check for each domain whether the user already owns it.
	CheckPortfolio bool `json:"checkPortfolio,omitempty"`

	// The ip of the client
	ClientIP string `json:"clientIp,omitempty"`

	// The ctid of the request. Only readable!
	Ctid string `json:"ctid,omitempty"`

	// The currency for every price lookup
	Currency string `json:"currency,omitempty"`

	// Activates debugging
	Debug bool `json:"debug,omitempty"`

	// All whois checks will be done via dns check.
	ForceDNSCheck bool `json:"forceDnsCheck,omitempty"`

	// If set to true, the inital, recommended, geo, custom  and suggestion sources are filtered for market domains. This can result in empty lists!
	IgnoreMarket bool `json:"ignoreMarket,omitempty"`

	// If set to true, the inital, recommended, geo, custom and suggestion sources are filtered for premium domains. This can result in empty lists!
	IgnorePremium bool `json:"ignorePremium,omitempty"`

	// Defines whether to return only free domain names when service WHOIS is used for a source.
	OnlyAvailable bool `json:"onlyAvailable,omitempty"`

	// Domain search token
	SearchToken string `json:"searchToken,omitempty"`

	// Wrapper for the configuration for each source
	Sources *DomainStudioSources `json:"sources,omitempty"`

	// The stid of the request. Only readable!
	Stid string `json:"stid,omitempty"`

	// Defines the timeout for the whois duration in seconds
	WhoisTimeout int32 `json:"whoisTimeout,omitempty"`
}

// Validate validates this domain envelope search request
func (m *DomainEnvelopeSearchRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEnvelopeSearchRequest) validateSources(formats strfmt.Registry) error {
	if swag.IsZero(m.Sources) { // not required
		return nil
	}

	if m.Sources != nil {
		if err := m.Sources.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sources")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this domain envelope search request based on the context it is used
func (m *DomainEnvelopeSearchRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainEnvelopeSearchRequest) contextValidateSources(ctx context.Context, formats strfmt.Registry) error {

	if m.Sources != nil {

		if swag.IsZero(m.Sources) { // not required
			return nil
		}

		if err := m.Sources.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sources")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("sources")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainEnvelopeSearchRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainEnvelopeSearchRequest) UnmarshalBinary(b []byte) error {
	var res DomainEnvelopeSearchRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

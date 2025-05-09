// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DomainEnvelopeSearchService domain envelope search service
//
// swagger:model DomainEnvelopeSearchService
type DomainEnvelopeSearchService string

func NewDomainEnvelopeSearchService(value DomainEnvelopeSearchService) *DomainEnvelopeSearchService {
	return &value
}

// Pointer returns a pointer to a freshly-allocated DomainEnvelopeSearchService.
func (m DomainEnvelopeSearchService) Pointer() *DomainEnvelopeSearchService {
	return &m
}

const (

	// DomainEnvelopeSearchServiceWHOIS captures enum value "WHOIS"
	DomainEnvelopeSearchServiceWHOIS DomainEnvelopeSearchService = "WHOIS"

	// DomainEnvelopeSearchServicePRICE captures enum value "PRICE"
	DomainEnvelopeSearchServicePRICE DomainEnvelopeSearchService = "PRICE"

	// DomainEnvelopeSearchServiceESTIMATION captures enum value "ESTIMATION"
	DomainEnvelopeSearchServiceESTIMATION DomainEnvelopeSearchService = "ESTIMATION"
)

// for schema
var domainEnvelopeSearchServiceEnum []interface{}

func init() {
	var res []DomainEnvelopeSearchService
	if err := json.Unmarshal([]byte(`["WHOIS","PRICE","ESTIMATION"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		domainEnvelopeSearchServiceEnum = append(domainEnvelopeSearchServiceEnum, v)
	}
}

func (m DomainEnvelopeSearchService) validateDomainEnvelopeSearchServiceEnum(path, location string, value DomainEnvelopeSearchService) error {
	if err := validate.EnumCase(path, location, value, domainEnvelopeSearchServiceEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this domain envelope search service
func (m DomainEnvelopeSearchService) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateDomainEnvelopeSearchServiceEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this domain envelope search service based on context it is used
func (m DomainEnvelopeSearchService) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

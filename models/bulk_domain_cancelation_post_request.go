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

// BulkDomainCancelationPostRequest bulk domain cancelation post request
//
// swagger:model BulkDomainCancelationPostRequest
type BulkDomainCancelationPostRequest struct {

	// The objects to process
	Objects []*DomainCancelation `json:"objects"`

	// The template for objects to process
	Template *DomainCancelation `json:"template,omitempty"`
}

// Validate validates this bulk domain cancelation post request
func (m *BulkDomainCancelationPostRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkDomainCancelationPostRequest) validateObjects(formats strfmt.Registry) error {

	if swag.IsZero(m.Objects) { // not required
		return nil
	}

	for i := 0; i < len(m.Objects); i++ {
		if swag.IsZero(m.Objects[i]) { // not required
			continue
		}

		if m.Objects[i] != nil {
			if err := m.Objects[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkDomainCancelationPostRequest) validateTemplate(formats strfmt.Registry) error {

	if swag.IsZero(m.Template) { // not required
		return nil
	}

	if m.Template != nil {
		if err := m.Template.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("template")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkDomainCancelationPostRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkDomainCancelationPostRequest) UnmarshalBinary(b []byte) error {
	var res BulkDomainCancelationPostRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

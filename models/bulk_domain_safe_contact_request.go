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
)

// BulkDomainSafeContactRequest bulk domain safe contact request
//
// swagger:model BulkDomainSafeContactRequest
type BulkDomainSafeContactRequest struct {

	// The objects to process
	Objects []*Contact `json:"objects"`

	// The query to fetch the object instead of defined list
	Query *Query `json:"query,omitempty"`
}

// Validate validates this bulk domain safe contact request
func (m *BulkDomainSafeContactRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateObjects(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQuery(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkDomainSafeContactRequest) validateObjects(formats strfmt.Registry) error {
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
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkDomainSafeContactRequest) validateQuery(formats strfmt.Registry) error {
	if swag.IsZero(m.Query) { // not required
		return nil
	}

	if m.Query != nil {
		if err := m.Query.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this bulk domain safe contact request based on the context it is used
func (m *BulkDomainSafeContactRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateObjects(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQuery(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BulkDomainSafeContactRequest) contextValidateObjects(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Objects); i++ {

		if m.Objects[i] != nil {

			if swag.IsZero(m.Objects[i]) { // not required
				return nil
			}

			if err := m.Objects[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("objects" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("objects" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BulkDomainSafeContactRequest) contextValidateQuery(ctx context.Context, formats strfmt.Registry) error {

	if m.Query != nil {

		if swag.IsZero(m.Query) { // not required
			return nil
		}

		if err := m.Query.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("query")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("query")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BulkDomainSafeContactRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BulkDomainSafeContactRequest) UnmarshalBinary(b []byte) error {
	var res BulkDomainSafeContactRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

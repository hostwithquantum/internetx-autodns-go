// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PriceChangeExclude price change exclude
//
// swagger:model PriceChangeExclude
type PriceChangeExclude struct {

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Customer
	Customer *GenericCustomer `json:"customer,omitempty"`

	// The unique identifier of the price
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// The priceChange
	PriceChange *PriceChange `json:"priceChange,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this price change exclude
func (m *PriceChangeExclude) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceChangeExclude) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *PriceChangeExclude) validateCustomer(formats strfmt.Registry) error {
	if swag.IsZero(m.Customer) { // not required
		return nil
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *PriceChangeExclude) validatePriceChange(formats strfmt.Registry) error {
	if swag.IsZero(m.PriceChange) { // not required
		return nil
	}

	if m.PriceChange != nil {
		if err := m.PriceChange.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priceChange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priceChange")
			}
			return err
		}
	}

	return nil
}

func (m *PriceChangeExclude) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this price change exclude based on the context it is used
func (m *PriceChangeExclude) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriceChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PriceChangeExclude) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

	if m.Customer != nil {

		if swag.IsZero(m.Customer) { // not required
			return nil
		}

		if err := m.Customer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *PriceChangeExclude) contextValidatePriceChange(ctx context.Context, formats strfmt.Registry) error {

	if m.PriceChange != nil {

		if swag.IsZero(m.PriceChange) { // not required
			return nil
		}

		if err := m.PriceChange.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priceChange")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priceChange")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PriceChangeExclude) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PriceChangeExclude) UnmarshalBinary(b []byte) error {
	var res PriceChangeExclude
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

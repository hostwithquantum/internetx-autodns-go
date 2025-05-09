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

// CustomerPriceList customer price list
//
// swagger:model CustomerPriceList
type CustomerPriceList struct {

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// customer
	Customer *GenericCustomer `json:"customer,omitempty"`

	// The id.
	ID int32 `json:"id,omitempty"`

	// The priceList label.
	PriceList *PriceList `json:"priceList,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this customer price list
func (m *CustomerPriceList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceList(formats); err != nil {
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

func (m *CustomerPriceList) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CustomerPriceList) validateCustomer(formats strfmt.Registry) error {
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

func (m *CustomerPriceList) validatePriceList(formats strfmt.Registry) error {
	if swag.IsZero(m.PriceList) { // not required
		return nil
	}

	if m.PriceList != nil {
		if err := m.PriceList.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priceList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priceList")
			}
			return err
		}
	}

	return nil
}

func (m *CustomerPriceList) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this customer price list based on the context it is used
func (m *CustomerPriceList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriceList(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomerPriceList) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *CustomerPriceList) contextValidatePriceList(ctx context.Context, formats strfmt.Registry) error {

	if m.PriceList != nil {

		if swag.IsZero(m.PriceList) { // not required
			return nil
		}

		if err := m.PriceList.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("priceList")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("priceList")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomerPriceList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomerPriceList) UnmarshalBinary(b []byte) error {
	var res CustomerPriceList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

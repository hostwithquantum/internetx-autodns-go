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

// ProductPriceTemplate product price template
//
// swagger:model ProductPriceTemplate
type ProductPriceTemplate struct {

	// The actual amount, before the price change processing
	ActualAmount float64 `json:"actualAmount,omitempty"`

	// Actual currency
	// Example: EUR
	ActualCurrency string `json:"actualCurrency,omitempty"`

	// Amount
	Amount float64 `json:"amount,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Used currency
	// Example: EUR
	Currency string `json:"currency,omitempty"`

	// Customer
	Customer *GenericCustomer `json:"customer,omitempty"`

	// Indicates if price discountable is possible
	Discountable bool `json:"discountable,omitempty"`

	// The unique identifier of the price
	// Example: 1
	ID int32 `json:"id,omitempty"`

	// The margin between amount and purchase amount
	Margin float32 `json:"margin,omitempty"`

	// The activity period of a product
	Period *TimePeriod `json:"period,omitempty"`

	// The priceChange
	PriceChange *PriceChange `json:"priceChange,omitempty"`

	// ProductPriceTemplate Price condition
	PriceConditions []*PriceServiceEntity `json:"priceConditions"`

	// Priority
	Priority PriorityConstants `json:"priority,omitempty"`

	// The product
	Product *Product `json:"product,omitempty"`

	// The actual purchase amount
	PurchaseAmount float64 `json:"purchaseAmount,omitempty"`

	// The rounding strategy
	Relative RelativeConstants `json:"relative,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this product price template
func (m *ProductPriceTemplate) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriceConditions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProduct(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRelative(formats); err != nil {
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

func (m *ProductPriceTemplate) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProductPriceTemplate) validateCustomer(formats strfmt.Registry) error {
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

func (m *ProductPriceTemplate) validatePeriod(formats strfmt.Registry) error {
	if swag.IsZero(m.Period) { // not required
		return nil
	}

	if m.Period != nil {
		if err := m.Period.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("period")
			}
			return err
		}
	}

	return nil
}

func (m *ProductPriceTemplate) validatePriceChange(formats strfmt.Registry) error {
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

func (m *ProductPriceTemplate) validatePriceConditions(formats strfmt.Registry) error {
	if swag.IsZero(m.PriceConditions) { // not required
		return nil
	}

	for i := 0; i < len(m.PriceConditions); i++ {
		if swag.IsZero(m.PriceConditions[i]) { // not required
			continue
		}

		if m.PriceConditions[i] != nil {
			if err := m.PriceConditions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priceConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductPriceTemplate) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

func (m *ProductPriceTemplate) validateProduct(formats strfmt.Registry) error {
	if swag.IsZero(m.Product) { // not required
		return nil
	}

	if m.Product != nil {
		if err := m.Product.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *ProductPriceTemplate) validateRelative(formats strfmt.Registry) error {
	if swag.IsZero(m.Relative) { // not required
		return nil
	}

	if err := m.Relative.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("relative")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("relative")
		}
		return err
	}

	return nil
}

func (m *ProductPriceTemplate) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this product price template based on the context it is used
func (m *ProductPriceTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePeriod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriceChange(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriceConditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePriority(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProduct(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRelative(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProductPriceTemplate) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductPriceTemplate) contextValidatePeriod(ctx context.Context, formats strfmt.Registry) error {

	if m.Period != nil {

		if swag.IsZero(m.Period) { // not required
			return nil
		}

		if err := m.Period.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("period")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("period")
			}
			return err
		}
	}

	return nil
}

func (m *ProductPriceTemplate) contextValidatePriceChange(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ProductPriceTemplate) contextValidatePriceConditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PriceConditions); i++ {

		if m.PriceConditions[i] != nil {

			if swag.IsZero(m.PriceConditions[i]) { // not required
				return nil
			}

			if err := m.PriceConditions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("priceConditions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("priceConditions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ProductPriceTemplate) contextValidatePriority(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	if err := m.Priority.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("priority")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("priority")
		}
		return err
	}

	return nil
}

func (m *ProductPriceTemplate) contextValidateProduct(ctx context.Context, formats strfmt.Registry) error {

	if m.Product != nil {

		if swag.IsZero(m.Product) { // not required
			return nil
		}

		if err := m.Product.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("product")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("product")
			}
			return err
		}
	}

	return nil
}

func (m *ProductPriceTemplate) contextValidateRelative(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Relative) { // not required
		return nil
	}

	if err := m.Relative.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("relative")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("relative")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProductPriceTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProductPriceTemplate) UnmarshalBinary(b []byte) error {
	var res ProductPriceTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

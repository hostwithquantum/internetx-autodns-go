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

// Invoice invoice
//
// swagger:model Invoice
type Invoice struct {

	// Invoice amount
	Amount float64 `json:"amount,omitempty"`

	// Comment to the invoice
	Comment string `json:"comment,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Used currency of the amount
	// Example: EUR
	Currency string `json:"currency,omitempty"`

	// Customer
	Customer *GenericCustomer `json:"customer,omitempty"`

	// The linked pdf
	Document *Document `json:"document,omitempty"`

	// Additional information
	Extension Configuration `json:"extension,omitempty"`

	// Indicator that shows that something is wrong with the process
	Failed bool `json:"failed,omitempty"`

	// internal id
	ID int64 `json:"id,omitempty"`

	// messages
	Messages []string `json:"messages"`

	// invoice number
	Number string `json:"number,omitempty"`

	// The object owner.
	Owner *BasicUser `json:"owner,omitempty"`

	// Date of payment
	// Format: date-time
	Paid strfmt.DateTime `json:"paid,omitempty"`

	// The method of payment: PRE, POST, etc.
	Payment PaymentConstants `json:"payment,omitempty"`

	// Additional payment data, such as provider: ipayment
	PaymentMode string `json:"paymentMode,omitempty"`

	// The payment transaction number
	PaymentTransaction string `json:"paymentTransaction,omitempty"`

	// Date of SEPA collection
	// Format: date-time
	SepaMandateCollection strfmt.DateTime `json:"sepaMandateCollection,omitempty"`

	// SEPA Mandate Reference
	SepaMandateReference string `json:"sepaMandateReference,omitempty"`

	// Invoice status
	Status InvoiceStatusConstants `json:"status,omitempty"`

	// The sub type of the invoice, e.g. domain invoice or server invoice
	SubType string `json:"subType,omitempty"`

	// Type of invoice
	Type AccountingDocumentTypeConstants `json:"type,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// User who performed the last update.
	Updater *BasicUser `json:"updater,omitempty"`

	// Vat amount
	VatAmount float64 `json:"vatAmount,omitempty"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDocument(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaid(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePayment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSepaMandateCollection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdater(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateCustomer(formats strfmt.Registry) error {
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

func (m *Invoice) validateDocument(formats strfmt.Registry) error {
	if swag.IsZero(m.Document) { // not required
		return nil
	}

	if m.Document != nil {
		if err := m.Document.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) validateOwner(formats strfmt.Registry) error {
	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) validatePaid(formats strfmt.Registry) error {
	if swag.IsZero(m.Paid) { // not required
		return nil
	}

	if err := validate.FormatOf("paid", "body", "date-time", m.Paid.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePayment(formats strfmt.Registry) error {
	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if err := m.Payment.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payment")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("payment")
		}
		return err
	}

	return nil
}

func (m *Invoice) validateSepaMandateCollection(formats strfmt.Registry) error {
	if swag.IsZero(m.SepaMandateCollection) { // not required
		return nil
	}

	if err := validate.FormatOf("sepaMandateCollection", "body", "date-time", m.SepaMandateCollection.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Invoice) validateType(formats strfmt.Registry) error {
	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *Invoice) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateUpdater(formats strfmt.Registry) error {
	if swag.IsZero(m.Updater) { // not required
		return nil
	}

	if m.Updater != nil {
		if err := m.Updater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this invoice based on the context it is used
func (m *Invoice) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCustomer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDocument(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePayment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdater(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) contextValidateCustomer(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Invoice) contextValidateDocument(ctx context.Context, formats strfmt.Registry) error {

	if m.Document != nil {

		if swag.IsZero(m.Document) { // not required
			return nil
		}

		if err := m.Document.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("document")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("document")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

	if m.Owner != nil {

		if swag.IsZero(m.Owner) { // not required
			return nil
		}

		if err := m.Owner.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *Invoice) contextValidatePayment(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Payment) { // not required
		return nil
	}

	if err := m.Payment.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("payment")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("payment")
		}
		return err
	}

	return nil
}

func (m *Invoice) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Invoice) contextValidateType(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	if err := m.Type.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *Invoice) contextValidateUpdater(ctx context.Context, formats strfmt.Registry) error {

	if m.Updater != nil {

		if swag.IsZero(m.Updater) { // not required
			return nil
		}

		if err := m.Updater.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

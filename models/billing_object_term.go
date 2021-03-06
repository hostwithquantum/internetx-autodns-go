// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BillingObjectTerm billing object term
//
// swagger:model BillingObjectTerm
type BillingObjectTerm struct {

	// The label of the articel, e.g. domain
	ArticleLabel string `json:"articleLabel,omitempty"`

	// The label of the article group
	ArticleTypeLabel string `json:"articleTypeLabel,omitempty"`

	// Flag indication if the article can be autodeleted
	Autodeleteable bool `json:"autodeleteable,omitempty"`

	// cancelation
	// Required: true
	Cancelation *TimePeriod `json:"cancelation"`

	// cancelation expire only
	CancelationExpireOnly bool `json:"cancelationExpireOnly,omitempty"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// the related customer or group
	// Required: true
	Customer *GenericCustomer `json:"customer"`

	// initial
	// Required: true
	Initial *TimePeriod `json:"initial"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// renew
	// Required: true
	Renew *TimePeriod `json:"renew"`

	// renew term
	RenewTerm *TimePeriod `json:"renewTerm,omitempty"`

	// Flag indication if the article can be restored
	Restoreable bool `json:"restoreable,omitempty"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating user of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this billing object term
func (m *BillingObjectTerm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCancelation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenew(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRenewTerm(formats); err != nil {
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

func (m *BillingObjectTerm) validateCancelation(formats strfmt.Registry) error {

	if err := validate.Required("cancelation", "body", m.Cancelation); err != nil {
		return err
	}

	if m.Cancelation != nil {
		if err := m.Cancelation.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cancelation")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingObjectTerm) validateCustomer(formats strfmt.Registry) error {

	if err := validate.Required("customer", "body", m.Customer); err != nil {
		return err
	}

	if m.Customer != nil {
		if err := m.Customer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customer")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateInitial(formats strfmt.Registry) error {

	if err := validate.Required("initial", "body", m.Initial); err != nil {
		return err
	}

	if m.Initial != nil {
		if err := m.Initial.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initial")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateOwner(formats strfmt.Registry) error {

	if swag.IsZero(m.Owner) { // not required
		return nil
	}

	if m.Owner != nil {
		if err := m.Owner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("owner")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateRenew(formats strfmt.Registry) error {

	if err := validate.Required("renew", "body", m.Renew); err != nil {
		return err
	}

	if m.Renew != nil {
		if err := m.Renew.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renew")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateRenewTerm(formats strfmt.Registry) error {

	if swag.IsZero(m.RenewTerm) { // not required
		return nil
	}

	if m.RenewTerm != nil {
		if err := m.RenewTerm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("renewTerm")
			}
			return err
		}
	}

	return nil
}

func (m *BillingObjectTerm) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillingObjectTerm) validateUpdater(formats strfmt.Registry) error {

	if swag.IsZero(m.Updater) { // not required
		return nil
	}

	if m.Updater != nil {
		if err := m.Updater.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updater")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BillingObjectTerm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillingObjectTerm) UnmarshalBinary(b []byte) error {
	var res BillingObjectTerm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

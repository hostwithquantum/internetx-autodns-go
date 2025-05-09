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

// Transfer transfer
//
// swagger:model Transfer
type Transfer struct {

	// Date of the automatic ACK on which the transfer is confirmed.
	// Format: date-time
	AutoAck strfmt.DateTime `json:"autoAck,omitempty"`

	// Automatic response to the transfer request.
	// false = not active
	// true = active
	// Default value = false
	// For XML, 0 (false) and 1 (true) can also be used.
	AutoAnswer bool `json:"autoAnswer,omitempty"`

	// Date of the automatic NACK on which the transfer is rejected.
	// Format: date-time
	AutoNack strfmt.DateTime `json:"autoNack,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The domain name.
	Domain string `json:"domain,omitempty"`

	// Date on which the transfer process ends.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The gaining registrar.
	GainingRegistrar string `json:"gainingRegistrar,omitempty"`

	// The loosing registrar.
	LoosingRegistrar string `json:"loosingRegistrar,omitempty"`

	// The object owner.
	Owner *BasicUser `json:"owner,omitempty"`

	// Receiver of the reminder email.
	Recipient string `json:"recipient,omitempty"`

	// Date on which the transfer reminder mail is sent.
	// Format: date-time
	Reminder strfmt.DateTime `json:"reminder,omitempty"`

	// Date on which the transfer started.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// Status of a running transfer.
	Status TransferStatusConstants `json:"status,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// User who performed the last update.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this transfer
func (m *Transfer) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAutoAck(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAutoNack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReminder(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *Transfer) validateAutoAck(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoAck) { // not required
		return nil
	}

	if err := validate.FormatOf("autoAck", "body", "date-time", m.AutoAck.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateAutoNack(formats strfmt.Registry) error {
	if swag.IsZero(m.AutoNack) { // not required
		return nil
	}

	if err := validate.FormatOf("autoNack", "body", "date-time", m.AutoNack.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateEnd(formats strfmt.Registry) error {
	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateOwner(formats strfmt.Registry) error {
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

func (m *Transfer) validateReminder(formats strfmt.Registry) error {
	if swag.IsZero(m.Reminder) { // not required
		return nil
	}

	if err := validate.FormatOf("reminder", "body", "date-time", m.Reminder.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateStart(formats strfmt.Registry) error {
	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateStatus(formats strfmt.Registry) error {
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

func (m *Transfer) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Transfer) validateUpdater(formats strfmt.Registry) error {
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

// ContextValidate validate this transfer based on the context it is used
func (m *Transfer) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
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

func (m *Transfer) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Transfer) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *Transfer) contextValidateUpdater(ctx context.Context, formats strfmt.Registry) error {

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
func (m *Transfer) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Transfer) UnmarshalBinary(b []byte) error {
	var res Transfer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

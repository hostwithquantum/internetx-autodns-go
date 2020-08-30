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

// TransferOut transfer out
//
// swagger:model TransferOut
type TransferOut struct {

	// The auto ack date.
	// Format: date-time
	AutoAck strfmt.DateTime `json:"autoAck,omitempty"`

	// Autoanswer active.
	AutoAnswer bool `json:"autoAnswer,omitempty"`

	// The auto nack date.
	// Format: date-time
	AutoNack strfmt.DateTime `json:"autoNack,omitempty"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The delivered date.
	// Format: date-time
	Delivered strfmt.DateTime `json:"delivered,omitempty"`

	// The delivered mailserver.
	DeliveredMailserver string `json:"deliveredMailserver,omitempty"`

	// The domain name.
	// Required: true
	Domain *string `json:"domain"`

	// The end date.
	// Format: date-time
	End strfmt.DateTime `json:"end,omitempty"`

	// The gaining registrar.
	GainingRegistrar string `json:"gainingRegistrar,omitempty"`

	// The loosing registrar.
	LoosingRegistrar string `json:"loosingRegistrar,omitempty"`

	// The mailserver.
	Mailserver string `json:"mailserver,omitempty"`

	// The reason. Possible values are : 1 = Evidence of fraud / 2 = Current UDRP action / 3 = Court order / 4 = Identity dispute / 5 = No payment for previous registration period / 6 = Express written objection to the transfer from the transfer contact.
	NackReason int32 `json:"nackReason,omitempty"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// The recipient.
	Recipient string `json:"recipient,omitempty"`

	// The reminder date.
	// Format: date-time
	Reminder strfmt.DateTime `json:"reminder,omitempty"`

	// The start date.
	// Format: date-time
	Start strfmt.DateTime `json:"start,omitempty"`

	// The ctid.
	Transaction string `json:"transaction,omitempty"`

	// The type of the transfer.
	// Required: true
	Type TransferAnswer `json:"type"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating user of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this transfer out
func (m *TransferOut) Validate(formats strfmt.Registry) error {
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

	if err := m.validateDelivered(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
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

func (m *TransferOut) validateAutoAck(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoAck) { // not required
		return nil
	}

	if err := validate.FormatOf("autoAck", "body", "date-time", m.AutoAck.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateAutoNack(formats strfmt.Registry) error {

	if swag.IsZero(m.AutoNack) { // not required
		return nil
	}

	if err := validate.FormatOf("autoNack", "body", "date-time", m.AutoNack.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateDelivered(formats strfmt.Registry) error {

	if swag.IsZero(m.Delivered) { // not required
		return nil
	}

	if err := validate.FormatOf("delivered", "body", "date-time", m.Delivered.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateEnd(formats strfmt.Registry) error {

	if swag.IsZero(m.End) { // not required
		return nil
	}

	if err := validate.FormatOf("end", "body", "date-time", m.End.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateOwner(formats strfmt.Registry) error {

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

func (m *TransferOut) validateReminder(formats strfmt.Registry) error {

	if swag.IsZero(m.Reminder) { // not required
		return nil
	}

	if err := validate.FormatOf("reminder", "body", "date-time", m.Reminder.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateStart(formats strfmt.Registry) error {

	if swag.IsZero(m.Start) { // not required
		return nil
	}

	if err := validate.FormatOf("start", "body", "date-time", m.Start.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *TransferOut) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *TransferOut) validateUpdater(formats strfmt.Registry) error {

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
func (m *TransferOut) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TransferOut) UnmarshalBinary(b []byte) error {
	var res TransferOut
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

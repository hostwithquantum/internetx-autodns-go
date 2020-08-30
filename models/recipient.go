// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Recipient recipient
//
// swagger:model Recipient
type Recipient struct {

	// The date of the next attempt
	// Format: date-time
	Attempt strfmt.DateTime `json:"attempt,omitempty"`

	// The number of attempts
	AttemptCount int32 `json:"attemptCount,omitempty"`

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// the email address of the recipient
	Email string `json:"email,omitempty"`

	// The date after the mail delivery will be expire.
	// Format: date-time
	Expire strfmt.DateTime `json:"expire,omitempty"`

	// The trace logs of each attempt
	Logs []*RecipientLog `json:"logs"`

	// The recipients role
	Role RecipientRole `json:"role,omitempty"`

	// The actual delivery status of the email
	Status DeliveryStatus `json:"status,omitempty"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`
}

// Validate validates this recipient
func (m *Recipient) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAttempt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpire(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogs(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
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

func (m *Recipient) validateAttempt(formats strfmt.Registry) error {

	if swag.IsZero(m.Attempt) { // not required
		return nil
	}

	if err := validate.FormatOf("attempt", "body", "date-time", m.Attempt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateExpire(formats strfmt.Registry) error {

	if swag.IsZero(m.Expire) { // not required
		return nil
	}

	if err := validate.FormatOf("expire", "body", "date-time", m.Expire.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Recipient) validateLogs(formats strfmt.Registry) error {

	if swag.IsZero(m.Logs) { // not required
		return nil
	}

	for i := 0; i < len(m.Logs); i++ {
		if swag.IsZero(m.Logs[i]) { // not required
			continue
		}

		if m.Logs[i] != nil {
			if err := m.Logs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("logs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Recipient) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if err := m.Role.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("role")
		}
		return err
	}

	return nil
}

func (m *Recipient) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.Status.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("status")
		}
		return err
	}

	return nil
}

func (m *Recipient) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Recipient) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Recipient) UnmarshalBinary(b []byte) error {
	var res Recipient
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
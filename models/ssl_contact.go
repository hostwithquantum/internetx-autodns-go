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

// SslContact ssl contact
//
// swagger:model SslContact
type SslContact struct {

	// The address of the contact.
	Address []string `json:"address"`

	// The city of the contact
	// Example: Anytown
	City string `json:"city,omitempty"`

	// The country of the contact
	// Example: DE
	Country string `json:"country,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The email address of the contact.
	// Example: john.doe@domain.com
	Email string `json:"email,omitempty"`

	// The contact extensions
	Extensions *SslContactExtensions `json:"extensions,omitempty"`

	// The fax number of the contact
	// Example: +49-123-12345
	Fax string `json:"fax,omitempty"`

	// The first name of the contact
	// Example: John
	Fname string `json:"fname,omitempty"`

	// Unique identifier of the object
	ID int32 `json:"id,omitempty"`

	// The last name of the contact
	// Example: Doe
	Lname string `json:"lname,omitempty"`

	// The name of organisation of the contact.
	// Example: Company
	Organization string `json:"organization,omitempty"`

	// The owner of the object
	Owner *BasicUser `json:"owner,omitempty"`

	// The postal code of the contact.
	// Example: 12345
	Pcode string `json:"pcode,omitempty"`

	// The phone number of the contact
	// Example: +49-123-12345
	Phone string `json:"phone,omitempty"`

	// The contact references
	References []*SslContactReference `json:"references"`

	// The local country state of the contact
	// Example: BY
	State string `json:"state,omitempty"`

	// The title of the contact
	// Example: Dr.
	Title string `json:"title,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updating using of the object
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this ssl contact
func (m *SslContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReferences(formats); err != nil {
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

func (m *SslContact) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateExtensions(formats strfmt.Registry) error {
	if swag.IsZero(m.Extensions) { // not required
		return nil
	}

	if m.Extensions != nil {
		if err := m.Extensions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

func (m *SslContact) validateOwner(formats strfmt.Registry) error {
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

func (m *SslContact) validateReferences(formats strfmt.Registry) error {
	if swag.IsZero(m.References) { // not required
		return nil
	}

	for i := 0; i < len(m.References); i++ {
		if swag.IsZero(m.References[i]) { // not required
			continue
		}

		if m.References[i] != nil {
			if err := m.References[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SslContact) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *SslContact) validateUpdater(formats strfmt.Registry) error {
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

// ContextValidate validate this ssl contact based on the context it is used
func (m *SslContact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateExtensions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateReferences(ctx, formats); err != nil {
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

func (m *SslContact) contextValidateExtensions(ctx context.Context, formats strfmt.Registry) error {

	if m.Extensions != nil {

		if swag.IsZero(m.Extensions) { // not required
			return nil
		}

		if err := m.Extensions.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("extensions")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("extensions")
			}
			return err
		}
	}

	return nil
}

func (m *SslContact) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *SslContact) contextValidateReferences(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.References); i++ {

		if m.References[i] != nil {

			if swag.IsZero(m.References[i]) { // not required
				return nil
			}

			if err := m.References[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("references" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("references" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *SslContact) contextValidateUpdater(ctx context.Context, formats strfmt.Registry) error {

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
func (m *SslContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SslContact) UnmarshalBinary(b []byte) error {
	var res SslContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

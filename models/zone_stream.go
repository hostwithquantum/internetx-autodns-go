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

// ZoneStream zone stream
//
// swagger:model ZoneStream
type ZoneStream struct {

	// the records to add.
	Adds []*ResourceRecord `json:"adds"`

	// the records to remove.
	Rems []*ResourceRecord `json:"rems"`
}

// Validate validates this zone stream
func (m *ZoneStream) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneStream) validateAdds(formats strfmt.Registry) error {
	if swag.IsZero(m.Adds) { // not required
		return nil
	}

	for i := 0; i < len(m.Adds); i++ {
		if swag.IsZero(m.Adds[i]) { // not required
			continue
		}

		if m.Adds[i] != nil {
			if err := m.Adds[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZoneStream) validateRems(formats strfmt.Registry) error {
	if swag.IsZero(m.Rems) { // not required
		return nil
	}

	for i := 0; i < len(m.Rems); i++ {
		if swag.IsZero(m.Rems[i]) { // not required
			continue
		}

		if m.Rems[i] != nil {
			if err := m.Rems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this zone stream based on the context it is used
func (m *ZoneStream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAdds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ZoneStream) contextValidateAdds(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Adds); i++ {

		if m.Adds[i] != nil {

			if swag.IsZero(m.Adds[i]) { // not required
				return nil
			}

			if err := m.Adds[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("adds" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("adds" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ZoneStream) contextValidateRems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Rems); i++ {

		if m.Rems[i] != nil {

			if swag.IsZero(m.Rems[i]) { // not required
				return nil
			}

			if err := m.Rems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("rems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ZoneStream) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ZoneStream) UnmarshalBinary(b []byte) error {
	var res ZoneStream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// JSONResponseDataListJSONResponseDataContact Json response data list Json response data contact
//
// swagger:model JsonResponseDataListJsonResponseDataContact
type JSONResponseDataListJSONResponseDataContact struct {

	// The CTID (Client Transaction ID) of the response.
	Ctid string `json:"ctid,omitempty"`

	// The data for the response. The type of the objects are depending on the request and are also specified in the ResponseObject value of the response.
	Data [][]*JSONResponseDataContact `json:"data"`

	// The messages belonging to the response.
	Messages []*Message `json:"messages"`

	// The response object.
	Object *ResponseObject `json:"object,omitempty"`

	// The status of the response.
	Status *ResponseStatus `json:"status,omitempty"`

	// The server transaction ID for the response.
	Stid string `json:"stid,omitempty"`
}

// Validate validates this Json response data list Json response data contact
func (m *JSONResponseDataListJSONResponseDataContact) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessages(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) validateData(formats strfmt.Registry) error {
	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {

		for ii := 0; ii < len(m.Data[i]); ii++ {
			if swag.IsZero(m.Data[i][ii]) { // not required
				continue
			}

			if m.Data[i][ii] != nil {
				if err := m.Data[i][ii].Validate(formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("data" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("data" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) validateMessages(formats strfmt.Registry) error {
	if swag.IsZero(m.Messages) { // not required
		return nil
	}

	for i := 0; i < len(m.Messages); i++ {
		if swag.IsZero(m.Messages[i]) { // not required
			continue
		}

		if m.Messages[i] != nil {
			if err := m.Messages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) validateObject(formats strfmt.Registry) error {
	if swag.IsZero(m.Object) { // not required
		return nil
	}

	if m.Object != nil {
		if err := m.Object.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this Json response data list Json response data contact based on the context it is used
func (m *JSONResponseDataListJSONResponseDataContact) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMessages(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObject(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Data); i++ {

		for ii := 0; ii < len(m.Data[i]); ii++ {

			if m.Data[i][ii] != nil {

				if swag.IsZero(m.Data[i][ii]) { // not required
					return nil
				}

				if err := m.Data[i][ii].ContextValidate(ctx, formats); err != nil {
					if ve, ok := err.(*errors.Validation); ok {
						return ve.ValidateName("data" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					} else if ce, ok := err.(*errors.CompositeError); ok {
						return ce.ValidateName("data" + "." + strconv.Itoa(i) + "." + strconv.Itoa(ii))
					}
					return err
				}
			}

		}

	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) contextValidateMessages(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Messages); i++ {

		if m.Messages[i] != nil {

			if swag.IsZero(m.Messages[i]) { // not required
				return nil
			}

			if err := m.Messages[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("messages" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("messages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) contextValidateObject(ctx context.Context, formats strfmt.Registry) error {

	if m.Object != nil {

		if swag.IsZero(m.Object) { // not required
			return nil
		}

		if err := m.Object.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("object")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("object")
			}
			return err
		}
	}

	return nil
}

func (m *JSONResponseDataListJSONResponseDataContact) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {

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
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONResponseDataListJSONResponseDataContact) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONResponseDataListJSONResponseDataContact) UnmarshalBinary(b []byte) error {
	var res JSONResponseDataListJSONResponseDataContact
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

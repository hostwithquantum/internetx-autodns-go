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

// ContextHost context host
//
// swagger:model ContextHost
type ContextHost struct {

	// The intermediates of the certificate
	CaCert string `json:"caCert,omitempty"`

	// The zones of the host
	Certificate *BasicCertificate `json:"certificate,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The csr of the certificate
	Csr string `json:"csr,omitempty"`

	// Show if host is default value
	Master bool `json:"master,omitempty"`

	// The host name
	Name string `json:"name,omitempty"`

	// The server certificate.
	ServerCert string `json:"serverCert,omitempty"`

	// The job status
	Status JobStatusConstants `json:"status,omitempty"`

	// The job substatus
	SubStatus string `json:"subStatus,omitempty"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The zones of the host
	Zone *ContextHostZone `json:"zone,omitempty"`
}

// Validate validates this context host
func (m *ContextHost) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCertificate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateZone(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContextHost) validateCertificate(formats strfmt.Registry) error {
	if swag.IsZero(m.Certificate) { // not required
		return nil
	}

	if m.Certificate != nil {
		if err := m.Certificate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *ContextHost) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContextHost) validateStatus(formats strfmt.Registry) error {
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

func (m *ContextHost) validateUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ContextHost) validateZone(formats strfmt.Registry) error {
	if swag.IsZero(m.Zone) { // not required
		return nil
	}

	if m.Zone != nil {
		if err := m.Zone.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this context host based on the context it is used
func (m *ContextHost) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCertificate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateZone(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ContextHost) contextValidateCertificate(ctx context.Context, formats strfmt.Registry) error {

	if m.Certificate != nil {

		if swag.IsZero(m.Certificate) { // not required
			return nil
		}

		if err := m.Certificate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("certificate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("certificate")
			}
			return err
		}
	}

	return nil
}

func (m *ContextHost) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

func (m *ContextHost) contextValidateZone(ctx context.Context, formats strfmt.Registry) error {

	if m.Zone != nil {

		if swag.IsZero(m.Zone) { // not required
			return nil
		}

		if err := m.Zone.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("zone")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("zone")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ContextHost) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ContextHost) UnmarshalBinary(b []byte) error {
	var res ContextHost
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

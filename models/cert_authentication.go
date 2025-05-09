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

// CertAuthentication cert authentication
//
// swagger:model CertAuthentication
type CertAuthentication struct {

	// The approver addresses for email based authentication.
	ApproverEmails []string `json:"approverEmails"`

	// The dns entry for dns based authentication. A zone record as string.
	DNS string `json:"dns,omitempty"`

	// The authentication data for every domain.
	Domains []*CertAuthentication `json:"domains"`

	// The content for file based authentication.
	FileContent string `json:"fileContent,omitempty"`

	// The filename for file based authentication.
	FileName string `json:"fileName,omitempty"`

	// The authentication method.
	Method AuthMethodConstants `json:"method,omitempty"`

	// The domain name the authentication data belongs to.
	Name string `json:"name,omitempty"`

	// Automatic zone provisioning for DNS validation.
	// Example: false
	Provisioning bool `json:"provisioning,omitempty"`

	// The authentication scope.
	Scope DcvValidationScope `json:"scope,omitempty"`

	// The status of the DCV.
	Status AuthenticateStatus `json:"status,omitempty"`
}

// Validate validates this cert authentication
func (m *CertAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDomains(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
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

func (m *CertAuthentication) validateDomains(formats strfmt.Registry) error {
	if swag.IsZero(m.Domains) { // not required
		return nil
	}

	for i := 0; i < len(m.Domains); i++ {
		if swag.IsZero(m.Domains[i]) { // not required
			continue
		}

		if m.Domains[i] != nil {
			if err := m.Domains[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CertAuthentication) validateMethod(formats strfmt.Registry) error {
	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if err := m.Method.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("method")
		}
		return err
	}

	return nil
}

func (m *CertAuthentication) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *CertAuthentication) validateStatus(formats strfmt.Registry) error {
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

// ContextValidate validate this cert authentication based on the context it is used
func (m *CertAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDomains(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMethod(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateScope(ctx, formats); err != nil {
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

func (m *CertAuthentication) contextValidateDomains(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Domains); i++ {

		if m.Domains[i] != nil {

			if swag.IsZero(m.Domains[i]) { // not required
				return nil
			}

			if err := m.Domains[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("domains" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("domains" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CertAuthentication) contextValidateMethod(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Method) { // not required
		return nil
	}

	if err := m.Method.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("method")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("method")
		}
		return err
	}

	return nil
}

func (m *CertAuthentication) contextValidateScope(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if err := m.Scope.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("scope")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("scope")
		}
		return err
	}

	return nil
}

func (m *CertAuthentication) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CertAuthentication) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CertAuthentication) UnmarshalBinary(b []byte) error {
	var res CertAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

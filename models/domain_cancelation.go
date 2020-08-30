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

// DomainCancelation domain cancelation
//
// swagger:model DomainCancelation
type DomainCancelation struct {

	// The created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// disconnect
	Disconnect bool `json:"disconnect,omitempty"`

	// domain
	// Required: true
	Domain *string `json:"domain"`

	// The execution type.
	// Required: true
	Execution ExecutionTypeConstants `json:"execution"`

	// gaining registrar
	GainingRegistrar string `json:"gainingRegistrar,omitempty"`

	// log Id
	LogID int64 `json:"logId,omitempty"`

	// Some remarks
	Notice string `json:"notice,omitempty"`

	// The owner of the object.
	Owner *BasicUser `json:"owner,omitempty"`

	// registry status
	RegistryStatus RegistryStatusConstants `json:"registryStatus,omitempty"`

	// The date of the execution. Only necessary when ExecutionType equals DATE.
	// Required: true
	// Format: date-time
	RegistryWhen *strfmt.DateTime `json:"registryWhen"`

	// The cancelation type. TRANSIT is only possible for certain TLDs.
	// Required: true
	Type CancelationTypeConstants `json:"type"`

	// The updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The updater of the object.
	Updater *BasicUser `json:"updater,omitempty"`
}

// Validate validates this domain cancelation
func (m *DomainCancelation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExecution(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryWhen(formats); err != nil {
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

func (m *DomainCancelation) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainCancelation) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *DomainCancelation) validateExecution(formats strfmt.Registry) error {

	if err := m.Execution.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("execution")
		}
		return err
	}

	return nil
}

func (m *DomainCancelation) validateOwner(formats strfmt.Registry) error {

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

func (m *DomainCancelation) validateRegistryStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.RegistryStatus) { // not required
		return nil
	}

	if err := m.RegistryStatus.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("registryStatus")
		}
		return err
	}

	return nil
}

func (m *DomainCancelation) validateRegistryWhen(formats strfmt.Registry) error {

	if err := validate.Required("registryWhen", "body", m.RegistryWhen); err != nil {
		return err
	}

	if err := validate.FormatOf("registryWhen", "body", "date-time", m.RegistryWhen.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainCancelation) validateType(formats strfmt.Registry) error {

	if err := m.Type.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("type")
		}
		return err
	}

	return nil
}

func (m *DomainCancelation) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DomainCancelation) validateUpdater(formats strfmt.Registry) error {

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
func (m *DomainCancelation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainCancelation) UnmarshalBinary(b []byte) error {
	var res DomainCancelation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

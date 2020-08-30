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

// User user
//
// swagger:model User
type User struct {

	// Wrapper for the user acls.
	Acls *UserAcls `json:"acls,omitempty"`

	// The ancestors of the user.
	Ancestors []*BasicUser `json:"ancestors"`

	// Holds the used applications.
	Applications []*TrustedApplication `json:"applications"`

	// The users authorization method.
	AuthType AuthType `json:"authType,omitempty"`

	// The context.
	// Required: true
	Context *int32 `json:"context"`

	// The user created date.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The customer belonging to the user.
	Customer *BasicCustomer `json:"customer,omitempty"`

	// The default email.
	DefaultEmail string `json:"defaultEmail,omitempty"`

	// The user details.
	Details *UserDetails `json:"details,omitempty"`

	// direct customer
	DirectCustomer bool `json:"directCustomer,omitempty"`

	// The language.
	Language string `json:"language,omitempty"`

	// The lock status of the user.
	Lock UserLock `json:"lock,omitempty"`

	// The available name server groups
	NameServerGroups []*VirtualNameServerGroup `json:"nameServerGroups"`

	// The parent.
	Parent *User `json:"parent,omitempty"`

	// The password.
	Password string `json:"password,omitempty"`

	// Wrapper for the user profiles.
	Profiles *UserProfileViews `json:"profiles,omitempty"`

	// The wrapper of the ip restrictions for the user.
	Restrictions *IPRestrictions `json:"restrictions,omitempty"`

	// Wrapper for the service user profiles.
	ServiceProfiles *ServiceProfiles `json:"serviceProfiles,omitempty"`

	// The status.
	Status int32 `json:"status,omitempty"`

	// Wrapper for the subscriptions.
	Subscriptions []*Subscription `json:"subscriptions"`

	// The substatus.
	Substatus int32 `json:"substatus,omitempty"`

	// The user updated date.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The user name.
	// Required: true
	// Pattern: ^[^_].*
	User *string `json:"user"`
}

// Validate validates this user
func (m *User) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAcls(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAncestors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplications(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuthType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCustomer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLock(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNameServerGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServiceProfiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubscriptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *User) validateAcls(formats strfmt.Registry) error {

	if swag.IsZero(m.Acls) { // not required
		return nil
	}

	if m.Acls != nil {
		if err := m.Acls.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("acls")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateAncestors(formats strfmt.Registry) error {

	if swag.IsZero(m.Ancestors) { // not required
		return nil
	}

	for i := 0; i < len(m.Ancestors); i++ {
		if swag.IsZero(m.Ancestors[i]) { // not required
			continue
		}

		if m.Ancestors[i] != nil {
			if err := m.Ancestors[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ancestors" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateApplications(formats strfmt.Registry) error {

	if swag.IsZero(m.Applications) { // not required
		return nil
	}

	for i := 0; i < len(m.Applications); i++ {
		if swag.IsZero(m.Applications[i]) { // not required
			continue
		}

		if m.Applications[i] != nil {
			if err := m.Applications[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("applications" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateAuthType(formats strfmt.Registry) error {

	if swag.IsZero(m.AuthType) { // not required
		return nil
	}

	if err := m.AuthType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("authType")
		}
		return err
	}

	return nil
}

func (m *User) validateContext(formats strfmt.Registry) error {

	if err := validate.Required("context", "body", m.Context); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateCustomer(formats strfmt.Registry) error {

	if swag.IsZero(m.Customer) { // not required
		return nil
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

func (m *User) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details) { // not required
		return nil
	}

	if m.Details != nil {
		if err := m.Details.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateLock(formats strfmt.Registry) error {

	if swag.IsZero(m.Lock) { // not required
		return nil
	}

	if err := m.Lock.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("lock")
		}
		return err
	}

	return nil
}

func (m *User) validateNameServerGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.NameServerGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.NameServerGroups); i++ {
		if swag.IsZero(m.NameServerGroups[i]) { // not required
			continue
		}

		if m.NameServerGroups[i] != nil {
			if err := m.NameServerGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("nameServerGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.Profiles) { // not required
		return nil
	}

	if m.Profiles != nil {
		if err := m.Profiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("profiles")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateRestrictions(formats strfmt.Registry) error {

	if swag.IsZero(m.Restrictions) { // not required
		return nil
	}

	if m.Restrictions != nil {
		if err := m.Restrictions.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("restrictions")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateServiceProfiles(formats strfmt.Registry) error {

	if swag.IsZero(m.ServiceProfiles) { // not required
		return nil
	}

	if m.ServiceProfiles != nil {
		if err := m.ServiceProfiles.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("serviceProfiles")
			}
			return err
		}
	}

	return nil
}

func (m *User) validateSubscriptions(formats strfmt.Registry) error {

	if swag.IsZero(m.Subscriptions) { // not required
		return nil
	}

	for i := 0; i < len(m.Subscriptions); i++ {
		if swag.IsZero(m.Subscriptions[i]) { // not required
			continue
		}

		if m.Subscriptions[i] != nil {
			if err := m.Subscriptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("subscriptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *User) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *User) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	if err := validate.Pattern("user", "body", string(*m.User), `^[^_].*`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *User) UnmarshalBinary(b []byte) error {
	var res User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MailProxy mail proxy
//
// swagger:model MailProxy
type MailProxy struct {

	// email address of the administrator
	Admin string `json:"admin,omitempty"`

	// The banned files options to use
	// Enum: [DISABLED QUARANTINE DISCARD ACCEPT]
	BannedFiles string `json:"bannedFiles,omitempty"`

	// The black listed email addresses
	Blacklist *MailList `json:"blacklist,omitempty"`

	// The date of the creation
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// The domain of the mail exchange to backup
	// Required: true
	Domain *string `json:"domain"`

	// The exculded listed email addresses
	Excludelist *MailList `json:"excludelist,omitempty"`

	// The grey listing policy
	Greylisting bool `json:"greylisting,omitempty"`

	// The mail header options to use
	// Enum: [DISABLED QUARANTINE DISCARD ACCEPT]
	Header string `json:"header,omitempty"`

	// The idn version of the domain.
	Idn string `json:"idn,omitempty"`

	// The owner of the entry
	Owner *BasicUser `json:"owner,omitempty"`

	// The protection level
	Protection ProtectionConstants `json:"protection,omitempty"`

	// The spam policy options
	Spam *SpamPolicy `json:"spam,omitempty"`

	// The hostname of the target mailserver
	// Required: true
	Target *string `json:"target"`

	// The date of the last updated
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// The last updater of the entry
	Updater *BasicUser `json:"updater,omitempty"`

	// The virus options to use
	// Enum: [DISABLED QUARANTINE DISCARD ACCEPT]
	Virus string `json:"virus,omitempty"`

	// The white listed email addresses
	Whitelist *MailList `json:"whitelist,omitempty"`
}

// Validate validates this mail proxy
func (m *MailProxy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBannedFiles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBlacklist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExcludelist(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHeader(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtection(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpam(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTarget(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdater(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVirus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhitelist(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var mailProxyTypeBannedFilesPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISABLED","QUARANTINE","DISCARD","ACCEPT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mailProxyTypeBannedFilesPropEnum = append(mailProxyTypeBannedFilesPropEnum, v)
	}
}

const (

	// MailProxyBannedFilesDISABLED captures enum value "DISABLED"
	MailProxyBannedFilesDISABLED string = "DISABLED"

	// MailProxyBannedFilesQUARANTINE captures enum value "QUARANTINE"
	MailProxyBannedFilesQUARANTINE string = "QUARANTINE"

	// MailProxyBannedFilesDISCARD captures enum value "DISCARD"
	MailProxyBannedFilesDISCARD string = "DISCARD"

	// MailProxyBannedFilesACCEPT captures enum value "ACCEPT"
	MailProxyBannedFilesACCEPT string = "ACCEPT"
)

// prop value enum
func (m *MailProxy) validateBannedFilesEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mailProxyTypeBannedFilesPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MailProxy) validateBannedFiles(formats strfmt.Registry) error {

	if swag.IsZero(m.BannedFiles) { // not required
		return nil
	}

	// value enum
	if err := m.validateBannedFilesEnum("bannedFiles", "body", m.BannedFiles); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateBlacklist(formats strfmt.Registry) error {

	if swag.IsZero(m.Blacklist) { // not required
		return nil
	}

	if m.Blacklist != nil {
		if err := m.Blacklist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blacklist")
			}
			return err
		}
	}

	return nil
}

func (m *MailProxy) validateCreated(formats strfmt.Registry) error {

	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateDomain(formats strfmt.Registry) error {

	if err := validate.Required("domain", "body", m.Domain); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateExcludelist(formats strfmt.Registry) error {

	if swag.IsZero(m.Excludelist) { // not required
		return nil
	}

	if m.Excludelist != nil {
		if err := m.Excludelist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludelist")
			}
			return err
		}
	}

	return nil
}

var mailProxyTypeHeaderPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISABLED","QUARANTINE","DISCARD","ACCEPT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mailProxyTypeHeaderPropEnum = append(mailProxyTypeHeaderPropEnum, v)
	}
}

const (

	// MailProxyHeaderDISABLED captures enum value "DISABLED"
	MailProxyHeaderDISABLED string = "DISABLED"

	// MailProxyHeaderQUARANTINE captures enum value "QUARANTINE"
	MailProxyHeaderQUARANTINE string = "QUARANTINE"

	// MailProxyHeaderDISCARD captures enum value "DISCARD"
	MailProxyHeaderDISCARD string = "DISCARD"

	// MailProxyHeaderACCEPT captures enum value "ACCEPT"
	MailProxyHeaderACCEPT string = "ACCEPT"
)

// prop value enum
func (m *MailProxy) validateHeaderEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mailProxyTypeHeaderPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MailProxy) validateHeader(formats strfmt.Registry) error {

	if swag.IsZero(m.Header) { // not required
		return nil
	}

	// value enum
	if err := m.validateHeaderEnum("header", "body", m.Header); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateOwner(formats strfmt.Registry) error {

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

func (m *MailProxy) validateProtection(formats strfmt.Registry) error {

	if swag.IsZero(m.Protection) { // not required
		return nil
	}

	if err := m.Protection.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protection")
		}
		return err
	}

	return nil
}

func (m *MailProxy) validateSpam(formats strfmt.Registry) error {

	if swag.IsZero(m.Spam) { // not required
		return nil
	}

	if m.Spam != nil {
		if err := m.Spam.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spam")
			}
			return err
		}
	}

	return nil
}

func (m *MailProxy) validateTarget(formats strfmt.Registry) error {

	if err := validate.Required("target", "body", m.Target); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateUpdated(formats strfmt.Registry) error {

	if swag.IsZero(m.Updated) { // not required
		return nil
	}

	if err := validate.FormatOf("updated", "body", "date-time", m.Updated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateUpdater(formats strfmt.Registry) error {

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

var mailProxyTypeVirusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DISABLED","QUARANTINE","DISCARD","ACCEPT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mailProxyTypeVirusPropEnum = append(mailProxyTypeVirusPropEnum, v)
	}
}

const (

	// MailProxyVirusDISABLED captures enum value "DISABLED"
	MailProxyVirusDISABLED string = "DISABLED"

	// MailProxyVirusQUARANTINE captures enum value "QUARANTINE"
	MailProxyVirusQUARANTINE string = "QUARANTINE"

	// MailProxyVirusDISCARD captures enum value "DISCARD"
	MailProxyVirusDISCARD string = "DISCARD"

	// MailProxyVirusACCEPT captures enum value "ACCEPT"
	MailProxyVirusACCEPT string = "ACCEPT"
)

// prop value enum
func (m *MailProxy) validateVirusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mailProxyTypeVirusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MailProxy) validateVirus(formats strfmt.Registry) error {

	if swag.IsZero(m.Virus) { // not required
		return nil
	}

	// value enum
	if err := m.validateVirusEnum("virus", "body", m.Virus); err != nil {
		return err
	}

	return nil
}

func (m *MailProxy) validateWhitelist(formats strfmt.Registry) error {

	if swag.IsZero(m.Whitelist) { // not required
		return nil
	}

	if m.Whitelist != nil {
		if err := m.Whitelist.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whitelist")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MailProxy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MailProxy) UnmarshalBinary(b []byte) error {
	var res MailProxy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

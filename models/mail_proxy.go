// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
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

	// Email address of the administrator to whom notifications are sent to.
	Admin string `json:"admin,omitempty"`

	// Defines whether files should be checked and how banned files should be avoided.
	// Enum: ["DISABLED","QUARANTINE","DISCARD","ACCEPT"]
	BannedFiles string `json:"bannedFiles,omitempty"`

	// Specification of email addresses whose mails are always to be marked as spam.
	Blacklist *MailList `json:"blacklist,omitempty"`

	// Date of creation.
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Domain name for which the BackupMX Record is to be created.
	// Required: true
	Domain *string `json:"domain"`

	// Definition of administrative addresses that should never be ignored by spam filters. An example of this is the "Hostmaster" addresses, e.g. hostmaster@example.com.
	Excludelist *MailList `json:"excludelist,omitempty"`

	// If greylisting is activated, the first email from an unknown sender is rejected at first. Mails from this sender will only be accepted after a further delayed delivery attempt.
	// Enum: ["ENABLED","DISABLED"]
	Greylisting string `json:"greylisting,omitempty"`

	// Defines whether headers are to be checked and how banned headers are to be handled.
	// Enum: ["DISABLED","QUARANTINE","DISCARD","ACCEPT"]
	Header string `json:"header,omitempty"`

	// IDN version of the domain name written in Punycode.
	Idn string `json:"idn,omitempty"`

	// The object owner.
	Owner *BasicUser `json:"owner,omitempty"`

	// Security settings for handling infected mails.
	Protection ProtectionConstants `json:"protection,omitempty"`

	// The spam policy options.
	Spam *SpamPolicy `json:"spam,omitempty"`

	// Mail server to which the MailProxy should forward the emails.
	// Note that the MX record of your mail server must be removed from the zone.
	// Required: true
	Target *string `json:"target"`

	// Date of the last update.
	// Format: date-time
	Updated strfmt.DateTime `json:"updated,omitempty"`

	// User who performed the last update.
	Updater *BasicUser `json:"updater,omitempty"`

	// Defines whether to check for viruses and how to deal with detected viruses.
	// Enum: ["DISABLED","QUARANTINE","DISCARD","ACCEPT"]
	Virus string `json:"virus,omitempty"`

	// Define email addresses whose mails should be trusted and never marked as spam.
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

	if err := m.validateGreylisting(formats); err != nil {
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blacklist")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludelist")
			}
			return err
		}
	}

	return nil
}

var mailProxyTypeGreylistingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		mailProxyTypeGreylistingPropEnum = append(mailProxyTypeGreylistingPropEnum, v)
	}
}

const (

	// MailProxyGreylistingENABLED captures enum value "ENABLED"
	MailProxyGreylistingENABLED string = "ENABLED"

	// MailProxyGreylistingDISABLED captures enum value "DISABLED"
	MailProxyGreylistingDISABLED string = "DISABLED"
)

// prop value enum
func (m *MailProxy) validateGreylistingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, mailProxyTypeGreylistingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *MailProxy) validateGreylisting(formats strfmt.Registry) error {
	if swag.IsZero(m.Greylisting) { // not required
		return nil
	}

	// value enum
	if err := m.validateGreylistingEnum("greylisting", "body", m.Greylisting); err != nil {
		return err
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("owner")
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
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protection")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spam")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updater")
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
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("whitelist")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mail proxy based on the context it is used
func (m *MailProxy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateBlacklist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateExcludelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateOwner(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSpam(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUpdater(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWhitelist(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MailProxy) contextValidateBlacklist(ctx context.Context, formats strfmt.Registry) error {

	if m.Blacklist != nil {

		if swag.IsZero(m.Blacklist) { // not required
			return nil
		}

		if err := m.Blacklist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blacklist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("blacklist")
			}
			return err
		}
	}

	return nil
}

func (m *MailProxy) contextValidateExcludelist(ctx context.Context, formats strfmt.Registry) error {

	if m.Excludelist != nil {

		if swag.IsZero(m.Excludelist) { // not required
			return nil
		}

		if err := m.Excludelist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("excludelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("excludelist")
			}
			return err
		}
	}

	return nil
}

func (m *MailProxy) contextValidateOwner(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MailProxy) contextValidateProtection(ctx context.Context, formats strfmt.Registry) error {

	if swag.IsZero(m.Protection) { // not required
		return nil
	}

	if err := m.Protection.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("protection")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("protection")
		}
		return err
	}

	return nil
}

func (m *MailProxy) contextValidateSpam(ctx context.Context, formats strfmt.Registry) error {

	if m.Spam != nil {

		if swag.IsZero(m.Spam) { // not required
			return nil
		}

		if err := m.Spam.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spam")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("spam")
			}
			return err
		}
	}

	return nil
}

func (m *MailProxy) contextValidateUpdater(ctx context.Context, formats strfmt.Registry) error {

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

func (m *MailProxy) contextValidateWhitelist(ctx context.Context, formats strfmt.Registry) error {

	if m.Whitelist != nil {

		if swag.IsZero(m.Whitelist) { // not required
			return nil
		}

		if err := m.Whitelist.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("whitelist")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("whitelist")
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

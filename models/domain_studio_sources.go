// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DomainStudioSources domain studio sources
//
// swagger:model DomainStudioSources
type DomainStudioSources struct {

	// The configuration for the custom source
	Custom *DomainStudioSourceCustom `json:"custom,omitempty"`

	// The configuration for the geo source
	Geo *DomainStudioSourceGeo `json:"geo,omitempty"`

	// The configuration for the initial source
	Initial *DomainStudioSourceInitial `json:"initial,omitempty"`

	// The configuration for the online presence source
	OnlinePresence *DomainStudioSourceOnlinePresence `json:"onlinePresence,omitempty"`

	// The configuration for the personal names source
	PersonalNames *DomainStudioSourcePersonalNames `json:"personalNames,omitempty"`

	// The configuration for the prefix suffix source
	PrefixSuffix *DomainStudioSourcePrefixSuffix `json:"prefixSuffix,omitempty"`

	// The configuration for the premium source
	Premium *DomainStudioSourcePremium `json:"premium,omitempty"`

	// The configuration for the recommended source
	Recommended *DomainStudioSourceRecommended `json:"recommended,omitempty"`

	// The configuration for the similar source
	Similar *DomainStudioSourceSimilar `json:"similar,omitempty"`

	// The configuration for the spin word source
	SpinWord *DomainStudioSourceSpinWord `json:"spinWord,omitempty"`

	// The configuration for the suggestion source
	Suggestion *DomainStudioSourceSuggestion `json:"suggestion,omitempty"`

	// The configuration for the upcoming source
	Upcoming *DomainStudioSourceUpcoming `json:"upcoming,omitempty"`
}

// Validate validates this domain studio sources
func (m *DomainStudioSources) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGeo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitial(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOnlinePresence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePersonalNames(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefixSuffix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePremium(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRecommended(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSimilar(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSpinWord(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuggestion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpcoming(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DomainStudioSources) validateCustom(formats strfmt.Registry) error {

	if swag.IsZero(m.Custom) { // not required
		return nil
	}

	if m.Custom != nil {
		if err := m.Custom.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("custom")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateGeo(formats strfmt.Registry) error {

	if swag.IsZero(m.Geo) { // not required
		return nil
	}

	if m.Geo != nil {
		if err := m.Geo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("geo")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateInitial(formats strfmt.Registry) error {

	if swag.IsZero(m.Initial) { // not required
		return nil
	}

	if m.Initial != nil {
		if err := m.Initial.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initial")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateOnlinePresence(formats strfmt.Registry) error {

	if swag.IsZero(m.OnlinePresence) { // not required
		return nil
	}

	if m.OnlinePresence != nil {
		if err := m.OnlinePresence.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("onlinePresence")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validatePersonalNames(formats strfmt.Registry) error {

	if swag.IsZero(m.PersonalNames) { // not required
		return nil
	}

	if m.PersonalNames != nil {
		if err := m.PersonalNames.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("personalNames")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validatePrefixSuffix(formats strfmt.Registry) error {

	if swag.IsZero(m.PrefixSuffix) { // not required
		return nil
	}

	if m.PrefixSuffix != nil {
		if err := m.PrefixSuffix.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("prefixSuffix")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validatePremium(formats strfmt.Registry) error {

	if swag.IsZero(m.Premium) { // not required
		return nil
	}

	if m.Premium != nil {
		if err := m.Premium.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("premium")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateRecommended(formats strfmt.Registry) error {

	if swag.IsZero(m.Recommended) { // not required
		return nil
	}

	if m.Recommended != nil {
		if err := m.Recommended.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("recommended")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateSimilar(formats strfmt.Registry) error {

	if swag.IsZero(m.Similar) { // not required
		return nil
	}

	if m.Similar != nil {
		if err := m.Similar.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("similar")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateSpinWord(formats strfmt.Registry) error {

	if swag.IsZero(m.SpinWord) { // not required
		return nil
	}

	if m.SpinWord != nil {
		if err := m.SpinWord.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("spinWord")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateSuggestion(formats strfmt.Registry) error {

	if swag.IsZero(m.Suggestion) { // not required
		return nil
	}

	if m.Suggestion != nil {
		if err := m.Suggestion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("suggestion")
			}
			return err
		}
	}

	return nil
}

func (m *DomainStudioSources) validateUpcoming(formats strfmt.Registry) error {

	if swag.IsZero(m.Upcoming) { // not required
		return nil
	}

	if m.Upcoming != nil {
		if err := m.Upcoming.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("upcoming")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DomainStudioSources) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DomainStudioSources) UnmarshalBinary(b []byte) error {
	var res DomainStudioSources
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
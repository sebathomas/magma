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

// ReleaseChannel release channel
//
// swagger:model release_channel
type ReleaseChannel struct {

	// id
	// Required: true
	ID ChannelID `json:"id"`

	// name
	Name string `json:"name,omitempty"`

	// supported versions
	// Required: true
	SupportedVersions []string `json:"supported_versions"`
}

// Validate validates this release channel
func (m *ReleaseChannel) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSupportedVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseChannel) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", ChannelID(m.ID)); err != nil {
		return err
	}

	if err := m.ID.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

func (m *ReleaseChannel) validateSupportedVersions(formats strfmt.Registry) error {

	if err := validate.Required("supported_versions", "body", m.SupportedVersions); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this release channel based on the context it is used
func (m *ReleaseChannel) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReleaseChannel) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ID.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("id")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("id")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseChannel) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseChannel) UnmarshalBinary(b []byte) error {
	var res ReleaseChannel
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

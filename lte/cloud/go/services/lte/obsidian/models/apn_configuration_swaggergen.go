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

// ApnConfiguration apn configuration
//
// swagger:model apn_configuration
type ApnConfiguration struct {

	// ambr
	// Required: true
	Ambr *AggregatedMaximumBitrate `json:"ambr"`

	// Value identifier for PDN type (0=IPv4 1=IPv6 2=IPv4v6 3=IPv4orv6)
	// Enum: [0 1 2 3]
	PdnType uint32 `json:"pdn_type,omitempty"`

	// qos profile
	// Required: true
	QosProfile *QosProfile `json:"qos_profile"`
}

// Validate validates this apn configuration
func (m *ApnConfiguration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePdnType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateQosProfile(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApnConfiguration) validateAmbr(formats strfmt.Registry) error {

	if err := validate.Required("ambr", "body", m.Ambr); err != nil {
		return err
	}

	if m.Ambr != nil {
		if err := m.Ambr.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambr")
			}
			return err
		}
	}

	return nil
}

var apnConfigurationTypePdnTypePropEnum []interface{}

func init() {
	var res []uint32
	if err := json.Unmarshal([]byte(`[0,1,2,3]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		apnConfigurationTypePdnTypePropEnum = append(apnConfigurationTypePdnTypePropEnum, v)
	}
}

// prop value enum
func (m *ApnConfiguration) validatePdnTypeEnum(path, location string, value uint32) error {
	if err := validate.EnumCase(path, location, value, apnConfigurationTypePdnTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *ApnConfiguration) validatePdnType(formats strfmt.Registry) error {
	if swag.IsZero(m.PdnType) { // not required
		return nil
	}

	// value enum
	if err := m.validatePdnTypeEnum("pdn_type", "body", m.PdnType); err != nil {
		return err
	}

	return nil
}

func (m *ApnConfiguration) validateQosProfile(formats strfmt.Registry) error {

	if err := validate.Required("qos_profile", "body", m.QosProfile); err != nil {
		return err
	}

	if m.QosProfile != nil {
		if err := m.QosProfile.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos_profile")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this apn configuration based on the context it is used
func (m *ApnConfiguration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAmbr(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateQosProfile(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApnConfiguration) contextValidateAmbr(ctx context.Context, formats strfmt.Registry) error {

	if m.Ambr != nil {
		if err := m.Ambr.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambr")
			}
			return err
		}
	}

	return nil
}

func (m *ApnConfiguration) contextValidateQosProfile(ctx context.Context, formats strfmt.Registry) error {

	if m.QosProfile != nil {
		if err := m.QosProfile.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("qos_profile")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApnConfiguration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApnConfiguration) UnmarshalBinary(b []byte) error {
	var res ApnConfiguration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

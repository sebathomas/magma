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

// ApnResource apn resource
//
// swagger:model apn_resource
type ApnResource struct {

	// apn name
	// Required: true
	ApnName ApnName `json:"apn_name" magma_alt_name:"service_selection"`

	// gateway ip
	// Format: ipv4
	GatewayIP strfmt.IPv4 `json:"gateway_ip,omitempty"`

	// gateway mac
	// Format: mac
	GatewayMac strfmt.MAC `json:"gateway_mac,omitempty"`

	// APN resource ID must be unique across all gateways in a network
	// Required: true
	// Min Length: 1
	ID string `json:"id"`

	// vlan id
	VlanID uint32 `json:"vlan_id,omitempty"`
}

// Validate validates this apn resource
func (m *ApnResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateApnName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayIP(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGatewayMac(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApnResource) validateApnName(formats strfmt.Registry) error {

	if err := validate.Required("apn_name", "body", ApnName(m.ApnName)); err != nil {
		return err
	}

	if err := m.ApnName.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apn_name")
		}
		return err
	}

	return nil
}

func (m *ApnResource) validateGatewayIP(formats strfmt.Registry) error {
	if swag.IsZero(m.GatewayIP) { // not required
		return nil
	}

	if err := validate.FormatOf("gateway_ip", "body", "ipv4", m.GatewayIP.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApnResource) validateGatewayMac(formats strfmt.Registry) error {
	if swag.IsZero(m.GatewayMac) { // not required
		return nil
	}

	if err := validate.FormatOf("gateway_mac", "body", "mac", m.GatewayMac.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ApnResource) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", m.ID); err != nil {
		return err
	}

	if err := validate.MinLength("id", "body", m.ID, 1); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this apn resource based on the context it is used
func (m *ApnResource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateApnName(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApnResource) contextValidateApnName(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ApnName.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("apn_name")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("apn_name")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApnResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApnResource) UnmarshalBinary(b []byte) error {
	var res ApnResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

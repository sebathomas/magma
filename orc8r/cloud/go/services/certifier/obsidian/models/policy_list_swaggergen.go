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

// PolicyList An object that defines a user's permissions to access resources
// Example: {"policies":[{"action":"WRITE","effect":"ALLOW","path":"**","resourceType":"URI"},{"action":"WRITE","effect":"DENY","resourceIDs":["test_network1","test_network2"],"resourceType":"NETWORK_ID"},{"action":"WRITE","effect":"ALLOW","resourceIDs":[0,1,2],"resourceType":"TENANT_ID"}],"token":"op_6YHy0uT7DeuWyT3N9nkAOyoeyOI25fletJE69yHGGl4ifjfoq"}
//
// swagger:model policyList
type PolicyList struct {

	// policies
	// Required: true
	Policies Policies `json:"policies"`

	// token
	// Required: true
	Token *string `json:"token"`
}

// Validate validates this policy list
func (m *PolicyList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateToken(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyList) validatePolicies(formats strfmt.Registry) error {

	if err := validate.Required("policies", "body", m.Policies); err != nil {
		return err
	}

	if err := m.Policies.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policies")
		}
		return err
	}

	return nil
}

func (m *PolicyList) validateToken(formats strfmt.Registry) error {

	if err := validate.Required("token", "body", m.Token); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this policy list based on the context it is used
func (m *PolicyList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicies(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PolicyList) contextValidatePolicies(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Policies.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("policies")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PolicyList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PolicyList) UnmarshalBinary(b []byte) error {
	var res PolicyList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

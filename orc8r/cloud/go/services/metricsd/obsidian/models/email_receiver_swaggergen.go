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

// EmailReceiver email receiver
//
// swagger:model email_receiver
type EmailReceiver struct {

	// auth identity
	AuthIdentity string `json:"auth_identity,omitempty"`

	// auth password
	AuthPassword string `json:"auth_password,omitempty"`

	// auth secret
	AuthSecret string `json:"auth_secret,omitempty"`

	// auth username
	AuthUsername string `json:"auth_username,omitempty"`

	// from
	// Required: true
	From *string `json:"from"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// hello
	Hello string `json:"hello,omitempty"`

	// html
	HTML string `json:"html,omitempty"`

	// send resolved
	SendResolved bool `json:"send_resolved,omitempty"`

	// smarthost
	// Required: true
	Smarthost *string `json:"smarthost"`

	// text
	Text string `json:"text,omitempty"`

	// to
	// Required: true
	To *string `json:"to"`
}

// Validate validates this email receiver
func (m *EmailReceiver) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSmarthost(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EmailReceiver) validateFrom(formats strfmt.Registry) error {

	if err := validate.Required("from", "body", m.From); err != nil {
		return err
	}

	return nil
}

func (m *EmailReceiver) validateSmarthost(formats strfmt.Registry) error {

	if err := validate.Required("smarthost", "body", m.Smarthost); err != nil {
		return err
	}

	return nil
}

func (m *EmailReceiver) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this email receiver based on context it is used
func (m *EmailReceiver) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EmailReceiver) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailReceiver) UnmarshalBinary(b []byte) error {
	var res EmailReceiver
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

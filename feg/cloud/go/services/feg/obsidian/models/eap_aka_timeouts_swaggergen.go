// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EapAkaTimeouts eap aka timeouts
//
// swagger:model eap_aka_timeouts
type EapAkaTimeouts struct {

	// challenge ms
	// Example: 20000
	ChallengeMs uint32 `json:"challenge_ms,omitempty"`

	// error notification ms
	// Example: 10000
	ErrorNotificationMs uint32 `json:"error_notification_ms,omitempty"`

	// session authenticated ms
	// Example: 5000
	SessionAuthenticatedMs uint32 `json:"session_authenticated_ms,omitempty"`

	// session ms
	// Example: 43200000
	SessionMs uint32 `json:"session_ms,omitempty"`
}

// Validate validates this eap aka timeouts
func (m *EapAkaTimeouts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this eap aka timeouts based on context it is used
func (m *EapAkaTimeouts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EapAkaTimeouts) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EapAkaTimeouts) UnmarshalBinary(b []byte) error {
	var res EapAkaTimeouts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

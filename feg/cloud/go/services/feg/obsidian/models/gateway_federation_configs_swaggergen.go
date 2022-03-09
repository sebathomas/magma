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

// GatewayFederationConfigs Federation configuration for a gateway
//
// swagger:model gateway_federation_configs
type GatewayFederationConfigs struct {

	// aaa server
	// Required: true
	AaaServer *AaaServer `json:"aaa_server"`

	// csfb
	Csfb *Csfb `json:"csfb,omitempty"`

	// eap aka
	// Required: true
	EapAka *EapAka `json:"eap_aka"`

	// eap sim
	EapSim *EapSim `json:"eap_sim,omitempty"`

	// gx
	// Required: true
	Gx *Gx `json:"gx"`

	// gy
	// Required: true
	Gy *Gy `json:"gy"`

	// health
	// Required: true
	Health *Health `json:"health" magma_alt_name:"HEALTH"`

	// hss
	// Required: true
	Hss *Hss `json:"hss" magma_alt_name:"HSS"`

	// nh routes
	NhRoutes NhRoutes `json:"nh_routes,omitempty"`

	// s6a
	// Required: true
	S6a *S6a `json:"s6a" magma_alt_name:"S6A"`

	// s8
	S8 *S8 `json:"s8,omitempty"`

	// served network ids
	// Required: true
	ServedNetworkIds ServedNetworkIds `json:"served_network_ids"`

	// served nh ids
	ServedNhIds ServedNhIds `json:"served_nh_ids,omitempty"`

	// swx
	// Required: true
	Swx *Swx `json:"swx" magma_alt_name:"SWX"`
}

// Validate validates this gateway federation configs
func (m *GatewayFederationConfigs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAaaServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCsfb(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEapAka(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEapSim(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGx(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHealth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHss(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNhRoutes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS6a(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateS8(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServedNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServedNhIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwx(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayFederationConfigs) validateAaaServer(formats strfmt.Registry) error {

	if err := validate.Required("aaa_server", "body", m.AaaServer); err != nil {
		return err
	}

	if m.AaaServer != nil {
		if err := m.AaaServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aaa_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aaa_server")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateCsfb(formats strfmt.Registry) error {
	if swag.IsZero(m.Csfb) { // not required
		return nil
	}

	if m.Csfb != nil {
		if err := m.Csfb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csfb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csfb")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateEapAka(formats strfmt.Registry) error {

	if err := validate.Required("eap_aka", "body", m.EapAka); err != nil {
		return err
	}

	if m.EapAka != nil {
		if err := m.EapAka.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eap_aka")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateEapSim(formats strfmt.Registry) error {
	if swag.IsZero(m.EapSim) { // not required
		return nil
	}

	if m.EapSim != nil {
		if err := m.EapSim.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_sim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eap_sim")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateGx(formats strfmt.Registry) error {

	if err := validate.Required("gx", "body", m.Gx); err != nil {
		return err
	}

	if m.Gx != nil {
		if err := m.Gx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gx")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateGy(formats strfmt.Registry) error {

	if err := validate.Required("gy", "body", m.Gy); err != nil {
		return err
	}

	if m.Gy != nil {
		if err := m.Gy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gy")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateHealth(formats strfmt.Registry) error {

	if err := validate.Required("health", "body", m.Health); err != nil {
		return err
	}

	if m.Health != nil {
		if err := m.Health.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateHss(formats strfmt.Registry) error {

	if err := validate.Required("hss", "body", m.Hss); err != nil {
		return err
	}

	if m.Hss != nil {
		if err := m.Hss.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hss")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateNhRoutes(formats strfmt.Registry) error {
	if swag.IsZero(m.NhRoutes) { // not required
		return nil
	}

	if m.NhRoutes != nil {
		if err := m.NhRoutes.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nh_routes")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nh_routes")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateS6a(formats strfmt.Registry) error {

	if err := validate.Required("s6a", "body", m.S6a); err != nil {
		return err
	}

	if m.S6a != nil {
		if err := m.S6a.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s6a")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s6a")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateS8(formats strfmt.Registry) error {
	if swag.IsZero(m.S8) { // not required
		return nil
	}

	if m.S8 != nil {
		if err := m.S8.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s8")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s8")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) validateServedNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("served_network_ids", "body", m.ServedNetworkIds); err != nil {
		return err
	}

	if err := m.ServedNetworkIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("served_network_ids")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("served_network_ids")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) validateServedNhIds(formats strfmt.Registry) error {
	if swag.IsZero(m.ServedNhIds) { // not required
		return nil
	}

	if err := m.ServedNhIds.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("served_nh_ids")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("served_nh_ids")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) validateSwx(formats strfmt.Registry) error {

	if err := validate.Required("swx", "body", m.Swx); err != nil {
		return err
	}

	if m.Swx != nil {
		if err := m.Swx.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swx")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this gateway federation configs based on the context it is used
func (m *GatewayFederationConfigs) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAaaServer(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCsfb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEapAka(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEapSim(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHealth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHss(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNhRoutes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS6a(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateS8(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServedNetworkIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateServedNhIds(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwx(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GatewayFederationConfigs) contextValidateAaaServer(ctx context.Context, formats strfmt.Registry) error {

	if m.AaaServer != nil {
		if err := m.AaaServer.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aaa_server")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("aaa_server")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateCsfb(ctx context.Context, formats strfmt.Registry) error {

	if m.Csfb != nil {
		if err := m.Csfb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("csfb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("csfb")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateEapAka(ctx context.Context, formats strfmt.Registry) error {

	if m.EapAka != nil {
		if err := m.EapAka.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_aka")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eap_aka")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateEapSim(ctx context.Context, formats strfmt.Registry) error {

	if m.EapSim != nil {
		if err := m.EapSim.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eap_sim")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eap_sim")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateGx(ctx context.Context, formats strfmt.Registry) error {

	if m.Gx != nil {
		if err := m.Gx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gx")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateGy(ctx context.Context, formats strfmt.Registry) error {

	if m.Gy != nil {
		if err := m.Gy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("gy")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateHealth(ctx context.Context, formats strfmt.Registry) error {

	if m.Health != nil {
		if err := m.Health.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("health")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("health")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateHss(ctx context.Context, formats strfmt.Registry) error {

	if m.Hss != nil {
		if err := m.Hss.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hss")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("hss")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateNhRoutes(ctx context.Context, formats strfmt.Registry) error {

	if err := m.NhRoutes.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("nh_routes")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("nh_routes")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateS6a(ctx context.Context, formats strfmt.Registry) error {

	if m.S6a != nil {
		if err := m.S6a.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s6a")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s6a")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateS8(ctx context.Context, formats strfmt.Registry) error {

	if m.S8 != nil {
		if err := m.S8.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("s8")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("s8")
			}
			return err
		}
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateServedNetworkIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServedNetworkIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("served_network_ids")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("served_network_ids")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateServedNhIds(ctx context.Context, formats strfmt.Registry) error {

	if err := m.ServedNhIds.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("served_nh_ids")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("served_nh_ids")
		}
		return err
	}

	return nil
}

func (m *GatewayFederationConfigs) contextValidateSwx(ctx context.Context, formats strfmt.Registry) error {

	if m.Swx != nil {
		if err := m.Swx.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("swx")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("swx")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GatewayFederationConfigs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GatewayFederationConfigs) UnmarshalBinary(b []byte) error {
	var res GatewayFederationConfigs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

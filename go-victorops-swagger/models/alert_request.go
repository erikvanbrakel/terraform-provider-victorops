// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AlertRequest alert request
// swagger:model AlertRequest
type AlertRequest struct {

	// alert Id
	// Required: true
	AlertID *string `json:"alertId"`
}

// Validate validates this alert request
func (m *AlertRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AlertRequest) validateAlertID(formats strfmt.Registry) error {

	if err := validate.Required("alertId", "body", m.AlertID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AlertRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertRequest) UnmarshalBinary(b []byte) error {
	var res AlertRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

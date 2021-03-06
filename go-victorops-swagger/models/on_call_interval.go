// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OnCallInterval on call interval
// swagger:model OnCallInterval
type OnCallInterval struct {

	// duration
	Duration *OnCallIntervalDuration `json:"duration,omitempty"`

	// escalation policy
	EscalationPolicy *EscalationPolicy `json:"escalationPolicy,omitempty"`

	// off
	Off string `json:"off,omitempty"`

	// on
	On string `json:"on,omitempty"`
}

// Validate validates this on call interval
func (m *OnCallInterval) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDuration(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEscalationPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *OnCallInterval) validateDuration(formats strfmt.Registry) error {

	if swag.IsZero(m.Duration) { // not required
		return nil
	}

	if m.Duration != nil {
		if err := m.Duration.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("duration")
			}
			return err
		}
	}

	return nil
}

func (m *OnCallInterval) validateEscalationPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.EscalationPolicy) { // not required
		return nil
	}

	if m.EscalationPolicy != nil {
		if err := m.EscalationPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("escalationPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *OnCallInterval) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnCallInterval) UnmarshalBinary(b []byte) error {
	var res OnCallInterval
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// TeamScheduleRollResource team schedule roll resource
// swagger:model TeamScheduleRollResource
type TeamScheduleRollResource struct {

	// The time the period is scheduled to start
	// Required: true
	Change *float64 `json:"change"`

	// is roll
	// Required: true
	IsRoll *bool `json:"isRoll"`

	// The user scheduled on call (if any)
	Oncall string `json:"oncall,omitempty"`

	// The time the period is scheduled to end
	// Required: true
	Until *float64 `json:"until"`
}

// Validate validates this team schedule roll resource
func (m *TeamScheduleRollResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChange(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsRoll(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUntil(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *TeamScheduleRollResource) validateChange(formats strfmt.Registry) error {

	if err := validate.Required("change", "body", m.Change); err != nil {
		return err
	}

	return nil
}

func (m *TeamScheduleRollResource) validateIsRoll(formats strfmt.Registry) error {

	if err := validate.Required("isRoll", "body", m.IsRoll); err != nil {
		return err
	}

	return nil
}

func (m *TeamScheduleRollResource) validateUntil(formats strfmt.Registry) error {

	if err := validate.Required("until", "body", m.Until); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TeamScheduleRollResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TeamScheduleRollResource) UnmarshalBinary(b []byte) error {
	var res TeamScheduleRollResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

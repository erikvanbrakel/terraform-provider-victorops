// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CreatedIncident created incident
// swagger:model CreatedIncident
type CreatedIncident struct {

	// An error message
	Error string `json:"error,omitempty"`

	// The VictorOps incident number
	IncidentNumber string `json:"incidentNumber,omitempty"`
}

// Validate validates this created incident
func (m *CreatedIncident) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CreatedIncident) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreatedIncident) UnmarshalBinary(b []byte) error {
	var res CreatedIncident
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ActiveIncidentList The list of incidents in various states
// swagger:model ActiveIncidentList
type ActiveIncidentList struct {

	// The incident data
	Incidents []*ActiveIncidentInfo `json:"incidents"`
}

// Validate validates this active incident list
func (m *ActiveIncidentList) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIncidents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ActiveIncidentList) validateIncidents(formats strfmt.Registry) error {

	if swag.IsZero(m.Incidents) { // not required
		return nil
	}

	for i := 0; i < len(m.Incidents); i++ {
		if swag.IsZero(m.Incidents[i]) { // not required
			continue
		}

		if m.Incidents[i] != nil {
			if err := m.Incidents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("incidents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *ActiveIncidentList) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ActiveIncidentList) UnmarshalBinary(b []byte) error {
	var res ActiveIncidentList
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

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

// CreateBatchAlertRequestPayload A collection of IDs (uuid) for the alerts to be returned
// swagger:model CreateBatchAlertRequestPayload
type CreateBatchAlertRequestPayload struct {

	// alert ids
	AlertIds []AlertID `json:"alertIds"`
}

// Validate validates this create batch alert request payload
func (m *CreateBatchAlertRequestPayload) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlertIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBatchAlertRequestPayload) validateAlertIds(formats strfmt.Registry) error {

	if swag.IsZero(m.AlertIds) { // not required
		return nil
	}

	for i := 0; i < len(m.AlertIds); i++ {

		if err := m.AlertIds[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertIds" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateBatchAlertRequestPayload) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBatchAlertRequestPayload) UnmarshalBinary(b []byte) error {
	var res CreateBatchAlertRequestPayload
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

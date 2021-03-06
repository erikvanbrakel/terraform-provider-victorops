// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// PostAPIPublicV1ProfileUsernamePoliciesStepOKBody post Api public v1 profile username policies step o k body
// swagger:model postApiPublicV1ProfileUsernamePoliciesStepOKBody
type PostAPIPublicV1ProfileUsernamePoliciesStepOKBody struct {

	// self Url
	SelfURL string `json:"_selfUrl,omitempty"`

	// step rule
	StepRule *PagingPolicyStepRule `json:"stepRule,omitempty"`
}

// Validate validates this post Api public v1 profile username policies step o k body
func (m *PostAPIPublicV1ProfileUsernamePoliciesStepOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStepRule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostAPIPublicV1ProfileUsernamePoliciesStepOKBody) validateStepRule(formats strfmt.Registry) error {

	if swag.IsZero(m.StepRule) { // not required
		return nil
	}

	if m.StepRule != nil {
		if err := m.StepRule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stepRule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostAPIPublicV1ProfileUsernamePoliciesStepOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostAPIPublicV1ProfileUsernamePoliciesStepOKBody) UnmarshalBinary(b []byte) error {
	var res PostAPIPublicV1ProfileUsernamePoliciesStepOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

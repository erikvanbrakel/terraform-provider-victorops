// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// OnCallUsersResourceOnCallUser on call users resource on call user
// swagger:model onCallUsersResourceOnCallUser
type OnCallUsersResourceOnCallUser struct {

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this on call users resource on call user
func (m *OnCallUsersResourceOnCallUser) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *OnCallUsersResourceOnCallUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OnCallUsersResourceOnCallUser) UnmarshalBinary(b []byte) error {
	var res OnCallUsersResourceOnCallUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new user paging policies API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for user paging policies API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetAPIPublicV1UserUserPolicies gets a list of paging policies for a user

Get paging policies for a user

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) GetAPIPublicV1UserUserPolicies(params *GetAPIPublicV1UserUserPoliciesParams) (*GetAPIPublicV1UserUserPoliciesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIPublicV1UserUserPoliciesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIPublicV1UserUserPolicies",
		Method:             "GET",
		PathPattern:        "/api-public/v1/user/{user}/policies",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIPublicV1UserUserPoliciesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIPublicV1UserUserPoliciesOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

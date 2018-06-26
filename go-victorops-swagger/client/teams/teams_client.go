// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new teams API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for teams API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteAPIPublicV1TeamTeam removes a team

Remove a team from your organization.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) DeleteAPIPublicV1TeamTeam(params *DeleteAPIPublicV1TeamTeamParams) (*DeleteAPIPublicV1TeamTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIPublicV1TeamTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPIPublicV1TeamTeam",
		Method:             "DELETE",
		PathPattern:        "/api-public/v1/team/{team}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIPublicV1TeamTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAPIPublicV1TeamTeamOK), nil

}

/*
DeleteAPIPublicV1TeamTeamMembersUser removes a team member

Remove a team from your organization.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) DeleteAPIPublicV1TeamTeamMembersUser(params *DeleteAPIPublicV1TeamTeamMembersUserParams) (*DeleteAPIPublicV1TeamTeamMembersUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteAPIPublicV1TeamTeamMembersUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteAPIPublicV1TeamTeamMembersUser",
		Method:             "DELETE",
		PathPattern:        "/api-public/v1/team/{team}/members/{user}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteAPIPublicV1TeamTeamMembersUserReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteAPIPublicV1TeamTeamMembersUserOK), nil

}

/*
GetAPIPublicV1Team lists teams

Get a list of teams for your organization.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) GetAPIPublicV1Team(params *GetAPIPublicV1TeamParams) (*GetAPIPublicV1TeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIPublicV1TeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIPublicV1Team",
		Method:             "GET",
		PathPattern:        "/api-public/v1/team",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIPublicV1TeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIPublicV1TeamOK), nil

}

/*
GetAPIPublicV1TeamTeam retrieves information for a team

Get the information for the specified team.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) GetAPIPublicV1TeamTeam(params *GetAPIPublicV1TeamTeamParams) (*GetAPIPublicV1TeamTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIPublicV1TeamTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIPublicV1TeamTeam",
		Method:             "GET",
		PathPattern:        "/api-public/v1/team/{team}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIPublicV1TeamTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIPublicV1TeamTeamOK), nil

}

/*
GetAPIPublicV1TeamTeamMembers retrieves a list of members for a team

Get the members for the specified team.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) GetAPIPublicV1TeamTeamMembers(params *GetAPIPublicV1TeamTeamMembersParams) (*GetAPIPublicV1TeamTeamMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAPIPublicV1TeamTeamMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetAPIPublicV1TeamTeamMembers",
		Method:             "GET",
		PathPattern:        "/api-public/v1/team/{team}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetAPIPublicV1TeamTeamMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAPIPublicV1TeamTeamMembersOK), nil

}

/*
PostAPIPublicV1Team adds a team

Add a team to your organization.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) PostAPIPublicV1Team(params *PostAPIPublicV1TeamParams) (*PostAPIPublicV1TeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIPublicV1TeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPIPublicV1Team",
		Method:             "POST",
		PathPattern:        "/api-public/v1/team",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIPublicV1TeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAPIPublicV1TeamOK), nil

}

/*
PostAPIPublicV1TeamTeamMembers adds a team member

Add a team member to your team.

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) PostAPIPublicV1TeamTeamMembers(params *PostAPIPublicV1TeamTeamMembersParams) (*PostAPIPublicV1TeamTeamMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostAPIPublicV1TeamTeamMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PostAPIPublicV1TeamTeamMembers",
		Method:             "POST",
		PathPattern:        "/api-public/v1/team/{team}/members",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostAPIPublicV1TeamTeamMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostAPIPublicV1TeamTeamMembersOK), nil

}

/*
PutAPIPublicV1TeamTeam updates a team

Update the designated team

This API may be called a maximum of 15 times per minute.

*/
func (a *Client) PutAPIPublicV1TeamTeam(params *PutAPIPublicV1TeamTeamParams) (*PutAPIPublicV1TeamTeamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutAPIPublicV1TeamTeamParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "PutAPIPublicV1TeamTeam",
		Method:             "PUT",
		PathPattern:        "/api-public/v1/team/{team}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PutAPIPublicV1TeamTeamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutAPIPublicV1TeamTeamOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

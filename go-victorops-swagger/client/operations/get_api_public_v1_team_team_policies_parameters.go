// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAPIPublicV1TeamTeamPoliciesParams creates a new GetAPIPublicV1TeamTeamPoliciesParams object
// with the default values initialized.
func NewGetAPIPublicV1TeamTeamPoliciesParams() *GetAPIPublicV1TeamTeamPoliciesParams {
	var ()
	return &GetAPIPublicV1TeamTeamPoliciesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1TeamTeamPoliciesParamsWithTimeout creates a new GetAPIPublicV1TeamTeamPoliciesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1TeamTeamPoliciesParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1TeamTeamPoliciesParams {
	var ()
	return &GetAPIPublicV1TeamTeamPoliciesParams{

		timeout: timeout,
	}
}

// NewGetAPIPublicV1TeamTeamPoliciesParamsWithContext creates a new GetAPIPublicV1TeamTeamPoliciesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1TeamTeamPoliciesParamsWithContext(ctx context.Context) *GetAPIPublicV1TeamTeamPoliciesParams {
	var ()
	return &GetAPIPublicV1TeamTeamPoliciesParams{

		Context: ctx,
	}
}

// NewGetAPIPublicV1TeamTeamPoliciesParamsWithHTTPClient creates a new GetAPIPublicV1TeamTeamPoliciesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1TeamTeamPoliciesParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1TeamTeamPoliciesParams {
	var ()
	return &GetAPIPublicV1TeamTeamPoliciesParams{
		HTTPClient: client,
	}
}

/*GetAPIPublicV1TeamTeamPoliciesParams contains all the parameters to send to the API endpoint
for the get API public v1 team team policies operation typically these are written to a http.Request
*/
type GetAPIPublicV1TeamTeamPoliciesParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Team
	  The VictorOps team to fetch

	*/
	Team string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithContext(ctx context.Context) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithTeam adds the team to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WithTeam(team string) *GetAPIPublicV1TeamTeamPoliciesParams {
	o.SetTeam(team)
	return o
}

// SetTeam adds the team to the get API public v1 team team policies params
func (o *GetAPIPublicV1TeamTeamPoliciesParams) SetTeam(team string) {
	o.Team = team
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1TeamTeamPoliciesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-VO-Api-Id
	if err := r.SetHeaderParam("X-VO-Api-Id", o.XVOAPIID); err != nil {
		return err
	}

	// header param X-VO-Api-Key
	if err := r.SetHeaderParam("X-VO-Api-Key", o.XVOAPIKey); err != nil {
		return err
	}

	// path param team
	if err := r.SetPathParam("team", o.Team); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams creates a new DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized.
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams() *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithTimeout creates a new DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithTimeout(timeout time.Duration) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		timeout: timeout,
	}
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithContext creates a new DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithContext(ctx context.Context) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams{

		Context: ctx,
	}
}

// NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithHTTPClient creates a new DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParamsWithHTTPClient(client *http.Client) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	var ()
	return &DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams{
		HTTPClient: client,
	}
}

/*DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams contains all the parameters to send to the API endpoint
for the delete API public v1 profile username policies step rule operation typically these are written to a http.Request
*/
type DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Rule
	  Rule for a paging policy step

	*/
	Rule float64
	/*Step
	  Paging policy step

	*/
	Step float64
	/*Username
	  Your username

	*/
	Username string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithTimeout(timeout time.Duration) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithContext(ctx context.Context) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithHTTPClient(client *http.Client) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithXVOAPIID(xVOAPIID string) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithXVOAPIKey(xVOAPIKey string) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithRule adds the rule to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithRule(rule float64) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetRule(rule)
	return o
}

// SetRule adds the rule to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetRule(rule float64) {
	o.Rule = rule
}

// WithStep adds the step to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithStep(step float64) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetStep(step)
	return o
}

// SetStep adds the step to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetStep(step float64) {
	o.Step = step
}

// WithUsername adds the username to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WithUsername(username string) *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams {
	o.SetUsername(username)
	return o
}

// SetUsername adds the username to the delete API public v1 profile username policies step rule params
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) SetUsername(username string) {
	o.Username = username
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAPIPublicV1ProfileUsernamePoliciesStepRuleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param rule
	if err := r.SetPathParam("rule", swag.FormatFloat64(o.Rule)); err != nil {
		return err
	}

	// path param step
	if err := r.SetPathParam("step", swag.FormatFloat64(o.Step)); err != nil {
		return err
	}

	// path param username
	if err := r.SetPathParam("username", o.Username); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package user_contact_methods

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

// NewGetAPIPublicV1UserUserContactMethodsDevicesParams creates a new GetAPIPublicV1UserUserContactMethodsDevicesParams object
// with the default values initialized.
func NewGetAPIPublicV1UserUserContactMethodsDevicesParams() *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsDevicesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithTimeout creates a new GetAPIPublicV1UserUserContactMethodsDevicesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithTimeout(timeout time.Duration) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsDevicesParams{

		timeout: timeout,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithContext creates a new GetAPIPublicV1UserUserContactMethodsDevicesParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithContext(ctx context.Context) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsDevicesParams{

		Context: ctx,
	}
}

// NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithHTTPClient creates a new GetAPIPublicV1UserUserContactMethodsDevicesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAPIPublicV1UserUserContactMethodsDevicesParamsWithHTTPClient(client *http.Client) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	var ()
	return &GetAPIPublicV1UserUserContactMethodsDevicesParams{
		HTTPClient: client,
	}
}

/*GetAPIPublicV1UserUserContactMethodsDevicesParams contains all the parameters to send to the API endpoint
for the get API public v1 user user contact methods devices operation typically these are written to a http.Request
*/
type GetAPIPublicV1UserUserContactMethodsDevicesParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*User
	  The VictorOps user ID

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithTimeout(timeout time.Duration) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithContext(ctx context.Context) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithHTTPClient(client *http.Client) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithXVOAPIID(xVOAPIID string) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithXVOAPIKey(xVOAPIKey string) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithUser adds the user to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WithUser(user string) *GetAPIPublicV1UserUserContactMethodsDevicesParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get API public v1 user user contact methods devices params
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetAPIPublicV1UserUserContactMethodsDevicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

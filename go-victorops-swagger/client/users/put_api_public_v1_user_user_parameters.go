// Code generated by go-swagger; DO NOT EDIT.

package users

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

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// NewPutAPIPublicV1UserUserParams creates a new PutAPIPublicV1UserUserParams object
// with the default values initialized.
func NewPutAPIPublicV1UserUserParams() *PutAPIPublicV1UserUserParams {
	var ()
	return &PutAPIPublicV1UserUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutAPIPublicV1UserUserParamsWithTimeout creates a new PutAPIPublicV1UserUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutAPIPublicV1UserUserParamsWithTimeout(timeout time.Duration) *PutAPIPublicV1UserUserParams {
	var ()
	return &PutAPIPublicV1UserUserParams{

		timeout: timeout,
	}
}

// NewPutAPIPublicV1UserUserParamsWithContext creates a new PutAPIPublicV1UserUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutAPIPublicV1UserUserParamsWithContext(ctx context.Context) *PutAPIPublicV1UserUserParams {
	var ()
	return &PutAPIPublicV1UserUserParams{

		Context: ctx,
	}
}

// NewPutAPIPublicV1UserUserParamsWithHTTPClient creates a new PutAPIPublicV1UserUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutAPIPublicV1UserUserParamsWithHTTPClient(client *http.Client) *PutAPIPublicV1UserUserParams {
	var ()
	return &PutAPIPublicV1UserUserParams{
		HTTPClient: client,
	}
}

/*PutAPIPublicV1UserUserParams contains all the parameters to send to the API endpoint
for the put API public v1 user user operation typically these are written to a http.Request
*/
type PutAPIPublicV1UserUserParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body*/
	Body *models.AddUserPayload
	/*User
	  The VictorOps user to be updated

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithTimeout(timeout time.Duration) *PutAPIPublicV1UserUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithContext(ctx context.Context) *PutAPIPublicV1UserUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithHTTPClient(client *http.Client) *PutAPIPublicV1UserUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithXVOAPIID(xVOAPIID string) *PutAPIPublicV1UserUserParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithXVOAPIKey(xVOAPIKey string) *PutAPIPublicV1UserUserParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithBody(body *models.AddUserPayload) *PutAPIPublicV1UserUserParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetBody(body *models.AddUserPayload) {
	o.Body = body
}

// WithUser adds the user to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) WithUser(user string) *PutAPIPublicV1UserUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the put API public v1 user user params
func (o *PutAPIPublicV1UserUserParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *PutAPIPublicV1UserUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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

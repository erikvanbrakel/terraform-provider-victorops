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

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// NewPostAPIPublicV1UserUserContactMethodsEmailsParams creates a new PostAPIPublicV1UserUserContactMethodsEmailsParams object
// with the default values initialized.
func NewPostAPIPublicV1UserUserContactMethodsEmailsParams() *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	var ()
	return &PostAPIPublicV1UserUserContactMethodsEmailsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithTimeout creates a new PostAPIPublicV1UserUserContactMethodsEmailsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithTimeout(timeout time.Duration) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	var ()
	return &PostAPIPublicV1UserUserContactMethodsEmailsParams{

		timeout: timeout,
	}
}

// NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithContext creates a new PostAPIPublicV1UserUserContactMethodsEmailsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithContext(ctx context.Context) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	var ()
	return &PostAPIPublicV1UserUserContactMethodsEmailsParams{

		Context: ctx,
	}
}

// NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithHTTPClient creates a new PostAPIPublicV1UserUserContactMethodsEmailsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostAPIPublicV1UserUserContactMethodsEmailsParamsWithHTTPClient(client *http.Client) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	var ()
	return &PostAPIPublicV1UserUserContactMethodsEmailsParams{
		HTTPClient: client,
	}
}

/*PostAPIPublicV1UserUserContactMethodsEmailsParams contains all the parameters to send to the API endpoint
for the post API public v1 user user contact methods emails operation typically these are written to a http.Request
*/
type PostAPIPublicV1UserUserContactMethodsEmailsParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body
	  The contact email

	*/
	Body *models.ContactEmailAdd
	/*User
	  The VictorOps user ID

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithTimeout(timeout time.Duration) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithContext(ctx context.Context) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithHTTPClient(client *http.Client) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithXVOAPIID(xVOAPIID string) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithXVOAPIKey(xVOAPIKey string) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithBody(body *models.ContactEmailAdd) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetBody(body *models.ContactEmailAdd) {
	o.Body = body
}

// WithUser adds the user to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WithUser(user string) *PostAPIPublicV1UserUserContactMethodsEmailsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the post API public v1 user user contact methods emails params
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *PostAPIPublicV1UserUserContactMethodsEmailsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

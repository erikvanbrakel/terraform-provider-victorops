// Code generated by go-swagger; DO NOT EDIT.

package incidents

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

// NewPatchAPIPublicV1IncidentsAckParams creates a new PatchAPIPublicV1IncidentsAckParams object
// with the default values initialized.
func NewPatchAPIPublicV1IncidentsAckParams() *PatchAPIPublicV1IncidentsAckParams {
	var ()
	return &PatchAPIPublicV1IncidentsAckParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchAPIPublicV1IncidentsAckParamsWithTimeout creates a new PatchAPIPublicV1IncidentsAckParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchAPIPublicV1IncidentsAckParamsWithTimeout(timeout time.Duration) *PatchAPIPublicV1IncidentsAckParams {
	var ()
	return &PatchAPIPublicV1IncidentsAckParams{

		timeout: timeout,
	}
}

// NewPatchAPIPublicV1IncidentsAckParamsWithContext creates a new PatchAPIPublicV1IncidentsAckParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchAPIPublicV1IncidentsAckParamsWithContext(ctx context.Context) *PatchAPIPublicV1IncidentsAckParams {
	var ()
	return &PatchAPIPublicV1IncidentsAckParams{

		Context: ctx,
	}
}

// NewPatchAPIPublicV1IncidentsAckParamsWithHTTPClient creates a new PatchAPIPublicV1IncidentsAckParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchAPIPublicV1IncidentsAckParamsWithHTTPClient(client *http.Client) *PatchAPIPublicV1IncidentsAckParams {
	var ()
	return &PatchAPIPublicV1IncidentsAckParams{
		HTTPClient: client,
	}
}

/*PatchAPIPublicV1IncidentsAckParams contains all the parameters to send to the API endpoint
for the patch API public v1 incidents ack operation typically these are written to a http.Request
*/
type PatchAPIPublicV1IncidentsAckParams struct {

	/*XVOAPIID
	  Your API ID

	*/
	XVOAPIID string
	/*XVOAPIKey
	  Your API Key

	*/
	XVOAPIKey string
	/*Body
	  Ack/Resolve payload

	*/
	Body *models.AckOrResolveRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithTimeout(timeout time.Duration) *PatchAPIPublicV1IncidentsAckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithContext(ctx context.Context) *PatchAPIPublicV1IncidentsAckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithHTTPClient(client *http.Client) *PatchAPIPublicV1IncidentsAckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXVOAPIID adds the xVOAPIID to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithXVOAPIID(xVOAPIID string) *PatchAPIPublicV1IncidentsAckParams {
	o.SetXVOAPIID(xVOAPIID)
	return o
}

// SetXVOAPIID adds the xVOApiId to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetXVOAPIID(xVOAPIID string) {
	o.XVOAPIID = xVOAPIID
}

// WithXVOAPIKey adds the xVOAPIKey to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithXVOAPIKey(xVOAPIKey string) *PatchAPIPublicV1IncidentsAckParams {
	o.SetXVOAPIKey(xVOAPIKey)
	return o
}

// SetXVOAPIKey adds the xVOApiKey to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetXVOAPIKey(xVOAPIKey string) {
	o.XVOAPIKey = xVOAPIKey
}

// WithBody adds the body to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) WithBody(body *models.AckOrResolveRequest) *PatchAPIPublicV1IncidentsAckParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch API public v1 incidents ack params
func (o *PatchAPIPublicV1IncidentsAckParams) SetBody(body *models.AckOrResolveRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PatchAPIPublicV1IncidentsAckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

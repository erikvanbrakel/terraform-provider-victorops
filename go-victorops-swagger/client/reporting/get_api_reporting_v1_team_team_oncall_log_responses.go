// Code generated by go-swagger; DO NOT EDIT.

package reporting

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// GetAPIReportingV1TeamTeamOncallLogReader is a Reader for the GetAPIReportingV1TeamTeamOncallLog structure.
type GetAPIReportingV1TeamTeamOncallLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIReportingV1TeamTeamOncallLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIReportingV1TeamTeamOncallLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIReportingV1TeamTeamOncallLogBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIReportingV1TeamTeamOncallLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIReportingV1TeamTeamOncallLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIReportingV1TeamTeamOncallLogNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIReportingV1TeamTeamOncallLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIReportingV1TeamTeamOncallLogOK creates a GetAPIReportingV1TeamTeamOncallLogOK with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogOK() *GetAPIReportingV1TeamTeamOncallLogOK {
	return &GetAPIReportingV1TeamTeamOncallLogOK{}
}

/*GetAPIReportingV1TeamTeamOncallLogOK handles this case with default header values.

The result of the take request
*/
type GetAPIReportingV1TeamTeamOncallLogOK struct {
	Payload *models.OnCallLog
}

func (o *GetAPIReportingV1TeamTeamOncallLogOK) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogOK  %+v", 200, o.Payload)
}

func (o *GetAPIReportingV1TeamTeamOncallLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OnCallLog)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIReportingV1TeamTeamOncallLogBadRequest creates a GetAPIReportingV1TeamTeamOncallLogBadRequest with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogBadRequest() *GetAPIReportingV1TeamTeamOncallLogBadRequest {
	return &GetAPIReportingV1TeamTeamOncallLogBadRequest{}
}

/*GetAPIReportingV1TeamTeamOncallLogBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIReportingV1TeamTeamOncallLogBadRequest struct {
}

func (o *GetAPIReportingV1TeamTeamOncallLogBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogBadRequest ", 400)
}

func (o *GetAPIReportingV1TeamTeamOncallLogBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIReportingV1TeamTeamOncallLogUnauthorized creates a GetAPIReportingV1TeamTeamOncallLogUnauthorized with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogUnauthorized() *GetAPIReportingV1TeamTeamOncallLogUnauthorized {
	return &GetAPIReportingV1TeamTeamOncallLogUnauthorized{}
}

/*GetAPIReportingV1TeamTeamOncallLogUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIReportingV1TeamTeamOncallLogUnauthorized struct {
}

func (o *GetAPIReportingV1TeamTeamOncallLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogUnauthorized ", 401)
}

func (o *GetAPIReportingV1TeamTeamOncallLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIReportingV1TeamTeamOncallLogForbidden creates a GetAPIReportingV1TeamTeamOncallLogForbidden with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogForbidden() *GetAPIReportingV1TeamTeamOncallLogForbidden {
	return &GetAPIReportingV1TeamTeamOncallLogForbidden{}
}

/*GetAPIReportingV1TeamTeamOncallLogForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIReportingV1TeamTeamOncallLogForbidden struct {
}

func (o *GetAPIReportingV1TeamTeamOncallLogForbidden) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogForbidden ", 403)
}

func (o *GetAPIReportingV1TeamTeamOncallLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIReportingV1TeamTeamOncallLogNotFound creates a GetAPIReportingV1TeamTeamOncallLogNotFound with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogNotFound() *GetAPIReportingV1TeamTeamOncallLogNotFound {
	return &GetAPIReportingV1TeamTeamOncallLogNotFound{}
}

/*GetAPIReportingV1TeamTeamOncallLogNotFound handles this case with default header values.

Team or user(s) not found
*/
type GetAPIReportingV1TeamTeamOncallLogNotFound struct {
}

func (o *GetAPIReportingV1TeamTeamOncallLogNotFound) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogNotFound ", 404)
}

func (o *GetAPIReportingV1TeamTeamOncallLogNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIReportingV1TeamTeamOncallLogInternalServerError creates a GetAPIReportingV1TeamTeamOncallLogInternalServerError with default headers values
func NewGetAPIReportingV1TeamTeamOncallLogInternalServerError() *GetAPIReportingV1TeamTeamOncallLogInternalServerError {
	return &GetAPIReportingV1TeamTeamOncallLogInternalServerError{}
}

/*GetAPIReportingV1TeamTeamOncallLogInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIReportingV1TeamTeamOncallLogInternalServerError struct {
}

func (o *GetAPIReportingV1TeamTeamOncallLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-reporting/v1/team/{team}/oncall/log][%d] getApiReportingV1TeamTeamOncallLogInternalServerError ", 500)
}

func (o *GetAPIReportingV1TeamTeamOncallLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
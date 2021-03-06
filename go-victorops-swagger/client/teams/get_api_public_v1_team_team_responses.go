// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// GetAPIPublicV1TeamTeamReader is a Reader for the GetAPIPublicV1TeamTeam structure.
type GetAPIPublicV1TeamTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1TeamTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1TeamTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIPublicV1TeamTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIPublicV1TeamTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1TeamTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetAPIPublicV1TeamTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewGetAPIPublicV1TeamTeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1TeamTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1TeamTeamOK creates a GetAPIPublicV1TeamTeamOK with default headers values
func NewGetAPIPublicV1TeamTeamOK() *GetAPIPublicV1TeamTeamOK {
	return &GetAPIPublicV1TeamTeamOK{}
}

/*GetAPIPublicV1TeamTeamOK handles this case with default header values.

Some details about the team that was added
*/
type GetAPIPublicV1TeamTeamOK struct {
	Payload *models.AddTeamResponse
}

func (o *GetAPIPublicV1TeamTeamOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1TeamTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddTeamResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1TeamTeamBadRequest creates a GetAPIPublicV1TeamTeamBadRequest with default headers values
func NewGetAPIPublicV1TeamTeamBadRequest() *GetAPIPublicV1TeamTeamBadRequest {
	return &GetAPIPublicV1TeamTeamBadRequest{}
}

/*GetAPIPublicV1TeamTeamBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIPublicV1TeamTeamBadRequest struct {
}

func (o *GetAPIPublicV1TeamTeamBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamBadRequest ", 400)
}

func (o *GetAPIPublicV1TeamTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1TeamTeamUnauthorized creates a GetAPIPublicV1TeamTeamUnauthorized with default headers values
func NewGetAPIPublicV1TeamTeamUnauthorized() *GetAPIPublicV1TeamTeamUnauthorized {
	return &GetAPIPublicV1TeamTeamUnauthorized{}
}

/*GetAPIPublicV1TeamTeamUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1TeamTeamUnauthorized struct {
}

func (o *GetAPIPublicV1TeamTeamUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamUnauthorized ", 401)
}

func (o *GetAPIPublicV1TeamTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1TeamTeamForbidden creates a GetAPIPublicV1TeamTeamForbidden with default headers values
func NewGetAPIPublicV1TeamTeamForbidden() *GetAPIPublicV1TeamTeamForbidden {
	return &GetAPIPublicV1TeamTeamForbidden{}
}

/*GetAPIPublicV1TeamTeamForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1TeamTeamForbidden struct {
}

func (o *GetAPIPublicV1TeamTeamForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamForbidden ", 403)
}

func (o *GetAPIPublicV1TeamTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1TeamTeamNotFound creates a GetAPIPublicV1TeamTeamNotFound with default headers values
func NewGetAPIPublicV1TeamTeamNotFound() *GetAPIPublicV1TeamTeamNotFound {
	return &GetAPIPublicV1TeamTeamNotFound{}
}

/*GetAPIPublicV1TeamTeamNotFound handles this case with default header values.

Team not found
*/
type GetAPIPublicV1TeamTeamNotFound struct {
}

func (o *GetAPIPublicV1TeamTeamNotFound) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamNotFound ", 404)
}

func (o *GetAPIPublicV1TeamTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1TeamTeamUnprocessableEntity creates a GetAPIPublicV1TeamTeamUnprocessableEntity with default headers values
func NewGetAPIPublicV1TeamTeamUnprocessableEntity() *GetAPIPublicV1TeamTeamUnprocessableEntity {
	return &GetAPIPublicV1TeamTeamUnprocessableEntity{}
}

/*GetAPIPublicV1TeamTeamUnprocessableEntity handles this case with default header values.

Team name or email is unavailable, or you have reached your team limit.

*/
type GetAPIPublicV1TeamTeamUnprocessableEntity struct {
}

func (o *GetAPIPublicV1TeamTeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamUnprocessableEntity ", 422)
}

func (o *GetAPIPublicV1TeamTeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1TeamTeamInternalServerError creates a GetAPIPublicV1TeamTeamInternalServerError with default headers values
func NewGetAPIPublicV1TeamTeamInternalServerError() *GetAPIPublicV1TeamTeamInternalServerError {
	return &GetAPIPublicV1TeamTeamInternalServerError{}
}

/*GetAPIPublicV1TeamTeamInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1TeamTeamInternalServerError struct {
}

func (o *GetAPIPublicV1TeamTeamInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/team/{team}][%d] getApiPublicV1TeamTeamInternalServerError ", 500)
}

func (o *GetAPIPublicV1TeamTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

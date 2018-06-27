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

// PostAPIPublicV1TeamReader is a Reader for the PostAPIPublicV1Team structure.
type PostAPIPublicV1TeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1TeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1TeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1TeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1TeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1TeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIPublicV1TeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostAPIPublicV1TeamUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1TeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1TeamOK creates a PostAPIPublicV1TeamOK with default headers values
func NewPostAPIPublicV1TeamOK() *PostAPIPublicV1TeamOK {
	return &PostAPIPublicV1TeamOK{}
}

/*PostAPIPublicV1TeamOK handles this case with default header values.

Some details about the team that was added
*/
type PostAPIPublicV1TeamOK struct {
	Payload *models.AddTeamResponse
}

func (o *PostAPIPublicV1TeamOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1TeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddTeamResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1TeamBadRequest creates a PostAPIPublicV1TeamBadRequest with default headers values
func NewPostAPIPublicV1TeamBadRequest() *PostAPIPublicV1TeamBadRequest {
	return &PostAPIPublicV1TeamBadRequest{}
}

/*PostAPIPublicV1TeamBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1TeamBadRequest struct {
}

func (o *PostAPIPublicV1TeamBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamBadRequest ", 400)
}

func (o *PostAPIPublicV1TeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamUnauthorized creates a PostAPIPublicV1TeamUnauthorized with default headers values
func NewPostAPIPublicV1TeamUnauthorized() *PostAPIPublicV1TeamUnauthorized {
	return &PostAPIPublicV1TeamUnauthorized{}
}

/*PostAPIPublicV1TeamUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1TeamUnauthorized struct {
}

func (o *PostAPIPublicV1TeamUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamUnauthorized ", 401)
}

func (o *PostAPIPublicV1TeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamForbidden creates a PostAPIPublicV1TeamForbidden with default headers values
func NewPostAPIPublicV1TeamForbidden() *PostAPIPublicV1TeamForbidden {
	return &PostAPIPublicV1TeamForbidden{}
}

/*PostAPIPublicV1TeamForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1TeamForbidden struct {
}

func (o *PostAPIPublicV1TeamForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamForbidden ", 403)
}

func (o *PostAPIPublicV1TeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamNotFound creates a PostAPIPublicV1TeamNotFound with default headers values
func NewPostAPIPublicV1TeamNotFound() *PostAPIPublicV1TeamNotFound {
	return &PostAPIPublicV1TeamNotFound{}
}

/*PostAPIPublicV1TeamNotFound handles this case with default header values.

Team not found
*/
type PostAPIPublicV1TeamNotFound struct {
}

func (o *PostAPIPublicV1TeamNotFound) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamNotFound ", 404)
}

func (o *PostAPIPublicV1TeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamUnprocessableEntity creates a PostAPIPublicV1TeamUnprocessableEntity with default headers values
func NewPostAPIPublicV1TeamUnprocessableEntity() *PostAPIPublicV1TeamUnprocessableEntity {
	return &PostAPIPublicV1TeamUnprocessableEntity{}
}

/*PostAPIPublicV1TeamUnprocessableEntity handles this case with default header values.

Team name or email is unavailable, or you have reached your team limit.

*/
type PostAPIPublicV1TeamUnprocessableEntity struct {
}

func (o *PostAPIPublicV1TeamUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamUnprocessableEntity ", 422)
}

func (o *PostAPIPublicV1TeamUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1TeamInternalServerError creates a PostAPIPublicV1TeamInternalServerError with default headers values
func NewPostAPIPublicV1TeamInternalServerError() *PostAPIPublicV1TeamInternalServerError {
	return &PostAPIPublicV1TeamInternalServerError{}
}

/*PostAPIPublicV1TeamInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1TeamInternalServerError struct {
}

func (o *PostAPIPublicV1TeamInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/team][%d] postApiPublicV1TeamInternalServerError ", 500)
}

func (o *PostAPIPublicV1TeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

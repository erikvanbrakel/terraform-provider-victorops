// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// PostAPIPublicV1UserReader is a Reader for the PostAPIPublicV1User structure.
type PostAPIPublicV1UserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1UserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1UserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1UserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1UserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1UserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIPublicV1UserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 422:
		result := NewPostAPIPublicV1UserUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1UserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1UserOK creates a PostAPIPublicV1UserOK with default headers values
func NewPostAPIPublicV1UserOK() *PostAPIPublicV1UserOK {
	return &PostAPIPublicV1UserOK{}
}

/*PostAPIPublicV1UserOK handles this case with default header values.

Some details about the user that was added
*/
type PostAPIPublicV1UserOK struct {
	Payload *models.AddUserResponse
}

func (o *PostAPIPublicV1UserOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1UserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddUserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1UserBadRequest creates a PostAPIPublicV1UserBadRequest with default headers values
func NewPostAPIPublicV1UserBadRequest() *PostAPIPublicV1UserBadRequest {
	return &PostAPIPublicV1UserBadRequest{}
}

/*PostAPIPublicV1UserBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1UserBadRequest struct {
}

func (o *PostAPIPublicV1UserBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserBadRequest ", 400)
}

func (o *PostAPIPublicV1UserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUnauthorized creates a PostAPIPublicV1UserUnauthorized with default headers values
func NewPostAPIPublicV1UserUnauthorized() *PostAPIPublicV1UserUnauthorized {
	return &PostAPIPublicV1UserUnauthorized{}
}

/*PostAPIPublicV1UserUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1UserUnauthorized struct {
}

func (o *PostAPIPublicV1UserUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserUnauthorized ", 401)
}

func (o *PostAPIPublicV1UserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserForbidden creates a PostAPIPublicV1UserForbidden with default headers values
func NewPostAPIPublicV1UserForbidden() *PostAPIPublicV1UserForbidden {
	return &PostAPIPublicV1UserForbidden{}
}

/*PostAPIPublicV1UserForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1UserForbidden struct {
}

func (o *PostAPIPublicV1UserForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserForbidden ", 403)
}

func (o *PostAPIPublicV1UserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserNotFound creates a PostAPIPublicV1UserNotFound with default headers values
func NewPostAPIPublicV1UserNotFound() *PostAPIPublicV1UserNotFound {
	return &PostAPIPublicV1UserNotFound{}
}

/*PostAPIPublicV1UserNotFound handles this case with default header values.

User not found
*/
type PostAPIPublicV1UserNotFound struct {
}

func (o *PostAPIPublicV1UserNotFound) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserNotFound ", 404)
}

func (o *PostAPIPublicV1UserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUnprocessableEntity creates a PostAPIPublicV1UserUnprocessableEntity with default headers values
func NewPostAPIPublicV1UserUnprocessableEntity() *PostAPIPublicV1UserUnprocessableEntity {
	return &PostAPIPublicV1UserUnprocessableEntity{}
}

/*PostAPIPublicV1UserUnprocessableEntity handles this case with default header values.

Username or email is unavailable, or you have reached your user limit.

*/
type PostAPIPublicV1UserUnprocessableEntity struct {
}

func (o *PostAPIPublicV1UserUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserUnprocessableEntity ", 422)
}

func (o *PostAPIPublicV1UserUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserInternalServerError creates a PostAPIPublicV1UserInternalServerError with default headers values
func NewPostAPIPublicV1UserInternalServerError() *PostAPIPublicV1UserInternalServerError {
	return &PostAPIPublicV1UserInternalServerError{}
}

/*PostAPIPublicV1UserInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1UserInternalServerError struct {
}

func (o *PostAPIPublicV1UserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user][%d] postApiPublicV1UserInternalServerError ", 500)
}

func (o *PostAPIPublicV1UserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

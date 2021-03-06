// Code generated by go-swagger; DO NOT EDIT.

package user_contact_methods

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// PostAPIPublicV1UserUserContactMethodsPhonesReader is a Reader for the PostAPIPublicV1UserUserContactMethodsPhones structure.
type PostAPIPublicV1UserUserContactMethodsPhonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1UserUserContactMethodsPhonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1UserUserContactMethodsPhonesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesOK creates a PostAPIPublicV1UserUserContactMethodsPhonesOK with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesOK() *PostAPIPublicV1UserUserContactMethodsPhonesOK {
	return &PostAPIPublicV1UserUserContactMethodsPhonesOK{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesOK handles this case with default header values.

The list of contact phones for the user
*/
type PostAPIPublicV1UserUserContactMethodsPhonesOK struct {
	Payload *models.UserContact
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesBadRequest creates a PostAPIPublicV1UserUserContactMethodsPhonesBadRequest with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesBadRequest() *PostAPIPublicV1UserUserContactMethodsPhonesBadRequest {
	return &PostAPIPublicV1UserUserContactMethodsPhonesBadRequest{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1UserUserContactMethodsPhonesBadRequest struct {
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesBadRequest ", 400)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesUnauthorized creates a PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesUnauthorized() *PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized {
	return &PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized struct {
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesUnauthorized ", 401)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesForbidden creates a PostAPIPublicV1UserUserContactMethodsPhonesForbidden with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesForbidden() *PostAPIPublicV1UserUserContactMethodsPhonesForbidden {
	return &PostAPIPublicV1UserUserContactMethodsPhonesForbidden{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1UserUserContactMethodsPhonesForbidden struct {
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesForbidden ", 403)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesNotFound creates a PostAPIPublicV1UserUserContactMethodsPhonesNotFound with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesNotFound() *PostAPIPublicV1UserUserContactMethodsPhonesNotFound {
	return &PostAPIPublicV1UserUserContactMethodsPhonesNotFound{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesNotFound handles this case with default header values.

User not found
*/
type PostAPIPublicV1UserUserContactMethodsPhonesNotFound struct {
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesNotFound) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesNotFound ", 404)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1UserUserContactMethodsPhonesInternalServerError creates a PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError with default headers values
func NewPostAPIPublicV1UserUserContactMethodsPhonesInternalServerError() *PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError {
	return &PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError{}
}

/*PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError struct {
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/user/{user}/contact-methods/phones][%d] postApiPublicV1UserUserContactMethodsPhonesInternalServerError ", 500)
}

func (o *PostAPIPublicV1UserUserContactMethodsPhonesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

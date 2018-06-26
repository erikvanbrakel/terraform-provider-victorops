// Code generated by go-swagger; DO NOT EDIT.

package personal_paging_policies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/erikvanbrakel/terraform-provider-victorops/go-victorops-swagger/models"
)

// PostAPIPublicV1ProfileUsernamePoliciesStepReader is a Reader for the PostAPIPublicV1ProfileUsernamePoliciesStep structure.
type PostAPIPublicV1ProfileUsernamePoliciesStepReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAPIPublicV1ProfileUsernamePoliciesStepReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostAPIPublicV1ProfileUsernamePoliciesStepOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostAPIPublicV1ProfileUsernamePoliciesStepBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostAPIPublicV1ProfileUsernamePoliciesStepForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepOK creates a PostAPIPublicV1ProfileUsernamePoliciesStepOK with default headers values
func NewPostAPIPublicV1ProfileUsernamePoliciesStepOK() *PostAPIPublicV1ProfileUsernamePoliciesStepOK {
	return &PostAPIPublicV1ProfileUsernamePoliciesStepOK{}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepOK handles this case with default header values.

The created rule for the paging policy step
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepOK struct {
	Payload *models.PostAPIPublicV1ProfileUsernamePoliciesStepOKBody
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepOK) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/profile/{username}/policies/{step}][%d] postApiPublicV1ProfileUsernamePoliciesStepOK  %+v", 200, o.Payload)
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PostAPIPublicV1ProfileUsernamePoliciesStepOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepBadRequest creates a PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest with default headers values
func NewPostAPIPublicV1ProfileUsernamePoliciesStepBadRequest() *PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest {
	return &PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest{}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest struct {
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/profile/{username}/policies/{step}][%d] postApiPublicV1ProfileUsernamePoliciesStepBadRequest ", 400)
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized creates a PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized with default headers values
func NewPostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized() *PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized {
	return &PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized{}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized struct {
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/profile/{username}/policies/{step}][%d] postApiPublicV1ProfileUsernamePoliciesStepUnauthorized ", 401)
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepForbidden creates a PostAPIPublicV1ProfileUsernamePoliciesStepForbidden with default headers values
func NewPostAPIPublicV1ProfileUsernamePoliciesStepForbidden() *PostAPIPublicV1ProfileUsernamePoliciesStepForbidden {
	return &PostAPIPublicV1ProfileUsernamePoliciesStepForbidden{}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepForbidden struct {
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepForbidden) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/profile/{username}/policies/{step}][%d] postApiPublicV1ProfileUsernamePoliciesStepForbidden ", 403)
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError creates a PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError with default headers values
func NewPostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError() *PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError {
	return &PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError{}
}

/*PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError struct {
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api-public/v1/profile/{username}/policies/{step}][%d] postApiPublicV1ProfileUsernamePoliciesStepInternalServerError ", 500)
}

func (o *PostAPIPublicV1ProfileUsernamePoliciesStepInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

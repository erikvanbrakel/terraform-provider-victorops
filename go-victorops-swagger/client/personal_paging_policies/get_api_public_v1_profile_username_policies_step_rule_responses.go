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

// GetAPIPublicV1ProfileUsernamePoliciesStepRuleReader is a Reader for the GetAPIPublicV1ProfileUsernamePoliciesStepRule structure.
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleOK creates a GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK with default headers values
func NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleOK() *GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK {
	return &GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK{}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK handles this case with default header values.

The rule from the paging policy step
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK struct {
	Payload *models.GetAPIPublicV1ProfileUsernamePoliciesStepRuleOKBody
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] getApiPublicV1ProfileUsernamePoliciesStepRuleOK  %+v", 200, o.Payload)
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetAPIPublicV1ProfileUsernamePoliciesStepRuleOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest creates a GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest with default headers values
func NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest() *GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest {
	return &GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest{}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest handles this case with default header values.

Problem with the request arguments.  The response payload may include an error message.
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest struct {
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] getApiPublicV1ProfileUsernamePoliciesStepRuleBadRequest ", 400)
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized creates a GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized with default headers values
func NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized() *GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized {
	return &GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized{}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized handles this case with default header values.

Authentication parameters missing
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized struct {
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] getApiPublicV1ProfileUsernamePoliciesStepRuleUnauthorized ", 401)
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden creates a GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden with default headers values
func NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden() *GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden {
	return &GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden{}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden handles this case with default header values.

Authentication failed or rate-limit reached
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden struct {
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] getApiPublicV1ProfileUsernamePoliciesStepRuleForbidden ", 403)
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError creates a GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError with default headers values
func NewGetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError() *GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError {
	return &GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError{}
}

/*GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError struct {
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api-public/v1/profile/{username}/policies/{step}/{rule}][%d] getApiPublicV1ProfileUsernamePoliciesStepRuleInternalServerError ", 500)
}

func (o *GetAPIPublicV1ProfileUsernamePoliciesStepRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

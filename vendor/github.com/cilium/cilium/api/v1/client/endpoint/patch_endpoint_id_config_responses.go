// Code generated by go-swagger; DO NOT EDIT.

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cilium/cilium/api/v1/models"
)

// PatchEndpointIDConfigReader is a Reader for the PatchEndpointIDConfig structure.
type PatchEndpointIDConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchEndpointIDConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchEndpointIDConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPatchEndpointIDConfigInvalid()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPatchEndpointIDConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPatchEndpointIDConfigFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPatchEndpointIDConfigOK creates a PatchEndpointIDConfigOK with default headers values
func NewPatchEndpointIDConfigOK() *PatchEndpointIDConfigOK {
	return &PatchEndpointIDConfigOK{}
}

/*PatchEndpointIDConfigOK handles this case with default header values.

Success
*/
type PatchEndpointIDConfigOK struct {
}

func (o *PatchEndpointIDConfigOK) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/config][%d] patchEndpointIdConfigOK ", 200)
}

func (o *PatchEndpointIDConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEndpointIDConfigInvalid creates a PatchEndpointIDConfigInvalid with default headers values
func NewPatchEndpointIDConfigInvalid() *PatchEndpointIDConfigInvalid {
	return &PatchEndpointIDConfigInvalid{}
}

/*PatchEndpointIDConfigInvalid handles this case with default header values.

Invalid configuration request
*/
type PatchEndpointIDConfigInvalid struct {
}

func (o *PatchEndpointIDConfigInvalid) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/config][%d] patchEndpointIdConfigInvalid ", 400)
}

func (o *PatchEndpointIDConfigInvalid) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEndpointIDConfigNotFound creates a PatchEndpointIDConfigNotFound with default headers values
func NewPatchEndpointIDConfigNotFound() *PatchEndpointIDConfigNotFound {
	return &PatchEndpointIDConfigNotFound{}
}

/*PatchEndpointIDConfigNotFound handles this case with default header values.

Endpoint not found
*/
type PatchEndpointIDConfigNotFound struct {
}

func (o *PatchEndpointIDConfigNotFound) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/config][%d] patchEndpointIdConfigNotFound ", 404)
}

func (o *PatchEndpointIDConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchEndpointIDConfigFailed creates a PatchEndpointIDConfigFailed with default headers values
func NewPatchEndpointIDConfigFailed() *PatchEndpointIDConfigFailed {
	return &PatchEndpointIDConfigFailed{}
}

/*PatchEndpointIDConfigFailed handles this case with default header values.

Update failed. Details in message.
*/
type PatchEndpointIDConfigFailed struct {
	Payload models.Error
}

func (o *PatchEndpointIDConfigFailed) Error() string {
	return fmt.Sprintf("[PATCH /endpoint/{id}/config][%d] patchEndpointIdConfigFailed  %+v", 500, o.Payload)
}

func (o *PatchEndpointIDConfigFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
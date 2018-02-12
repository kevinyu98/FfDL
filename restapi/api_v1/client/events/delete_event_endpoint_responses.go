// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/IBM/FfDL/restapi/api_v1/restmodels"
)

// DeleteEventEndpointReader is a Reader for the DeleteEventEndpoint structure.
type DeleteEventEndpointReader struct {
	formats strfmt.Registry
	writer  io.Writer
}

// ReadResponse reads a server response into the received o.
func (o *DeleteEventEndpointReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEventEndpointOK(o.writer)
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteEventEndpointUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteEventEndpointNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 410:
		result := NewDeleteEventEndpointGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEventEndpointOK creates a DeleteEventEndpointOK with default headers values
func NewDeleteEventEndpointOK(writer io.Writer) *DeleteEventEndpointOK {
	return &DeleteEventEndpointOK{
		Payload: writer,
	}
}

/*DeleteEventEndpointOK handles this case with default header values.

Event updated successfully
*/
type DeleteEventEndpointOK struct {
	Payload io.Writer
}

func (o *DeleteEventEndpointOK) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] deleteEventEndpointOK  %+v", 200, o.Payload)
}

func (o *DeleteEventEndpointOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventEndpointUnauthorized creates a DeleteEventEndpointUnauthorized with default headers values
func NewDeleteEventEndpointUnauthorized() *DeleteEventEndpointUnauthorized {
	return &DeleteEventEndpointUnauthorized{}
}

/*DeleteEventEndpointUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteEventEndpointUnauthorized struct {
	Payload *restmodels.Error
}

func (o *DeleteEventEndpointUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] deleteEventEndpointUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteEventEndpointUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventEndpointNotFound creates a DeleteEventEndpointNotFound with default headers values
func NewDeleteEventEndpointNotFound() *DeleteEventEndpointNotFound {
	return &DeleteEventEndpointNotFound{}
}

/*DeleteEventEndpointNotFound handles this case with default header values.

The model or event type cannot be found.
*/
type DeleteEventEndpointNotFound struct {
	Payload *restmodels.Error
}

func (o *DeleteEventEndpointNotFound) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] deleteEventEndpointNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEventEndpointNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteEventEndpointGone creates a DeleteEventEndpointGone with default headers values
func NewDeleteEventEndpointGone() *DeleteEventEndpointGone {
	return &DeleteEventEndpointGone{}
}

/*DeleteEventEndpointGone handles this case with default header values.

If the trained model storage time has expired and it has been deleted. It only gets deleted if it not stored on an external data store.
*/
type DeleteEventEndpointGone struct {
	Payload *restmodels.Error
}

func (o *DeleteEventEndpointGone) Error() string {
	return fmt.Sprintf("[DELETE /v1/models/{model_id}/events/{event_type}/{endpoint_id}][%d] deleteEventEndpointGone  %+v", 410, o.Payload)
}

func (o *DeleteEventEndpointGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(restmodels.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
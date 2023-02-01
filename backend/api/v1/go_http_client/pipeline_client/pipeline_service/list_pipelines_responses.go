// Code generated by go-swagger; DO NOT EDIT.

package pipeline_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	pipeline_model "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/pipeline_model"
)

// ListPipelinesReader is a Reader for the ListPipelines structure.
type ListPipelinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPipelinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewListPipelinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewListPipelinesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListPipelinesOK creates a ListPipelinesOK with default headers values
func NewListPipelinesOK() *ListPipelinesOK {
	return &ListPipelinesOK{}
}

/*ListPipelinesOK handles this case with default header values.

A successful response.
*/
type ListPipelinesOK struct {
	Payload *pipeline_model.V1ListPipelinesResponse
}

func (o *ListPipelinesOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] listPipelinesOK  %+v", 200, o.Payload)
}

func (o *ListPipelinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V1ListPipelinesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPipelinesDefault creates a ListPipelinesDefault with default headers values
func NewListPipelinesDefault(code int) *ListPipelinesDefault {
	return &ListPipelinesDefault{
		_statusCode: code,
	}
}

/*ListPipelinesDefault handles this case with default header values.

ListPipelinesDefault list pipelines default
*/
type ListPipelinesDefault struct {
	_statusCode int

	Payload *pipeline_model.V1Status
}

// Code gets the status code for the list pipelines default response
func (o *ListPipelinesDefault) Code() int {
	return o._statusCode
}

func (o *ListPipelinesDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1/pipelines][%d] ListPipelines default  %+v", o._statusCode, o.Payload)
}

func (o *ListPipelinesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(pipeline_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

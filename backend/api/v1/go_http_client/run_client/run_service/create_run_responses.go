// Code generated by go-swagger; DO NOT EDIT.

package run_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	run_model "github.com/kubeflow/pipelines/backend/api/v1/go_http_client/run_model"
)

// CreateRunReader is a Reader for the CreateRun structure.
type CreateRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewCreateRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCreateRunOK creates a CreateRunOK with default headers values
func NewCreateRunOK() *CreateRunOK {
	return &CreateRunOK{}
}

/*CreateRunOK handles this case with default header values.

A successful response.
*/
type CreateRunOK struct {
	Payload *run_model.V1RunDetail
}

func (o *CreateRunOK) Error() string {
	return fmt.Sprintf("[POST /apis/v1/runs][%d] createRunOK  %+v", 200, o.Payload)
}

func (o *CreateRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1RunDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRunDefault creates a CreateRunDefault with default headers values
func NewCreateRunDefault(code int) *CreateRunDefault {
	return &CreateRunDefault{
		_statusCode: code,
	}
}

/*CreateRunDefault handles this case with default header values.

CreateRunDefault create run default
*/
type CreateRunDefault struct {
	_statusCode int

	Payload *run_model.V1Status
}

// Code gets the status code for the create run default response
func (o *CreateRunDefault) Code() int {
	return o._statusCode
}

func (o *CreateRunDefault) Error() string {
	return fmt.Sprintf("[POST /apis/v1/runs][%d] CreateRun default  %+v", o._statusCode, o.Payload)
}

func (o *CreateRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

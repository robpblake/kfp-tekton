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

// GetRunReader is a Reader for the GetRun structure.
type GetRunReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRunReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRunOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRunDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRunOK creates a GetRunOK with default headers values
func NewGetRunOK() *GetRunOK {
	return &GetRunOK{}
}

/*GetRunOK handles this case with default header values.

A successful response.
*/
type GetRunOK struct {
	Payload *run_model.V1RunDetail
}

func (o *GetRunOK) Error() string {
	return fmt.Sprintf("[GET /apis/v1/runs/{run_id}][%d] getRunOK  %+v", 200, o.Payload)
}

func (o *GetRunOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1RunDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRunDefault creates a GetRunDefault with default headers values
func NewGetRunDefault(code int) *GetRunDefault {
	return &GetRunDefault{
		_statusCode: code,
	}
}

/*GetRunDefault handles this case with default header values.

GetRunDefault get run default
*/
type GetRunDefault struct {
	_statusCode int

	Payload *run_model.V1Status
}

// Code gets the status code for the get run default response
func (o *GetRunDefault) Code() int {
	return o._statusCode
}

func (o *GetRunDefault) Error() string {
	return fmt.Sprintf("[GET /apis/v1/runs/{run_id}][%d] GetRun default  %+v", o._statusCode, o.Payload)
}

func (o *GetRunDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(run_model.V1Status)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

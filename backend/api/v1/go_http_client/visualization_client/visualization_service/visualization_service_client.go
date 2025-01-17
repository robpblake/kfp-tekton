// Code generated by go-swagger; DO NOT EDIT.

package visualization_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new visualization service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for visualization service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateVisualization create visualization API
*/
func (a *Client) CreateVisualization(params *CreateVisualizationParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVisualizationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVisualizationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateVisualization",
		Method:             "POST",
		PathPattern:        "/apis/v1/visualizations/{namespace}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateVisualizationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateVisualizationOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

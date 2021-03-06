// Code generated by go-swagger; DO NOT EDIT.

package document_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new document tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for document tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DocumentInfoWithAlias(params *DocumentInfoWithAliasParams) (*DocumentInfoWithAliasOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DocumentInfoWithAlias documents info

  Inquiring the document for the given alias. The alias can be an ID, UUID or a alias name like price_list.xml. For more information on Pricelist download look at our helpcenter.
*/
func (a *Client) DocumentInfoWithAlias(params *DocumentInfoWithAliasParams) (*DocumentInfoWithAliasOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentInfoWithAliasParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "documentInfoWithAlias",
		Method:             "GET",
		PathPattern:        "/document/{alias}",
		ProducesMediaTypes: []string{"application/octet-stream"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentInfoWithAliasReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DocumentInfoWithAliasOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for documentInfoWithAlias: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

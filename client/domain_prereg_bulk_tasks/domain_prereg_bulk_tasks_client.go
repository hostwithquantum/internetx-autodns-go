// Code generated by go-swagger; DO NOT EDIT.

package domain_prereg_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new domain prereg bulk tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for domain prereg bulk tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DomainPreregConfimrs(params *DomainPreregConfimrsParams) (*DomainPreregConfimrsOK, error)

	DomainPreregCreates(params *DomainPreregCreatesParams) (*DomainPreregCreatesOK, error)

	DomainPreregDeletes(params *DomainPreregDeletesParams) (*DomainPreregDeletesOK, error)

	DomainPreregPatches(params *DomainPreregPatchesParams) (*DomainPreregPatchesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DomainPreregConfimrs domains prereg confirm bulk

  Confirming several domain preregistrations with one request.
*/
func (a *Client) DomainPreregConfimrs(params *DomainPreregConfimrsParams) (*DomainPreregConfimrsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDomainPreregConfimrsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "domainPreregConfimrs",
		Method:             "PUT",
		PathPattern:        "/bulk/domainPrereg/_confirm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DomainPreregConfimrsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DomainPreregConfimrsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domainPreregConfimrs: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DomainPreregCreates domains prereg create bulk

  Creating several domain preregistrations with one request.
*/
func (a *Client) DomainPreregCreates(params *DomainPreregCreatesParams) (*DomainPreregCreatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDomainPreregCreatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "domainPreregCreates",
		Method:             "POST",
		PathPattern:        "/bulk/domainPrereg",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DomainPreregCreatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DomainPreregCreatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domainPreregCreates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DomainPreregDeletes domains prereg delete bulk

  Deleting several domain preregistrations with one request.
*/
func (a *Client) DomainPreregDeletes(params *DomainPreregDeletesParams) (*DomainPreregDeletesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDomainPreregDeletesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "domainPreregDeletes",
		Method:             "DELETE",
		PathPattern:        "/bulk/domainPrereg",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DomainPreregDeletesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DomainPreregDeletesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domainPreregDeletes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DomainPreregPatches domains prereg update bulk

  Updating several domain preregistrations with one request.
*/
func (a *Client) DomainPreregPatches(params *DomainPreregPatchesParams) (*DomainPreregPatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDomainPreregPatchesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "domainPreregPatches",
		Method:             "PATCH",
		PathPattern:        "/bulk/domainPrereg",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DomainPreregPatchesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DomainPreregPatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for domainPreregPatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
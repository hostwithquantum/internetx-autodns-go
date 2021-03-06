// Code generated by go-swagger; DO NOT EDIT.

package certificate_bulk_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new certificate bulk tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for certificate bulk tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CertificateCreates(params *CertificateCreatesParams) (*CertificateCreatesOK, error)

	CertificateDeletes(params *CertificateDeletesParams) (*CertificateDeletesOK, error)

	CertificatePatches(params *CertificatePatchesParams) (*CertificatePatchesOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CertificateCreates certificates create bulk

  Creating several new certificates.
*/
func (a *Client) CertificateCreates(params *CertificateCreatesParams) (*CertificateCreatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCertificateCreatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "certificateCreates",
		Method:             "POST",
		PathPattern:        "/bulk/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CertificateCreatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CertificateCreatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificateCreates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CertificateDeletes certificates delete bulk

  Deleting several existing certificates.
*/
func (a *Client) CertificateDeletes(params *CertificateDeletesParams) (*CertificateDeletesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCertificateDeletesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "certificateDeletes",
		Method:             "DELETE",
		PathPattern:        "/bulk/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CertificateDeletesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CertificateDeletesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificateDeletes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  CertificatePatches certificates update bulk

  Updating several existing certificates.
*/
func (a *Client) CertificatePatches(params *CertificatePatchesParams) (*CertificatePatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCertificatePatchesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "certificatePatches",
		Method:             "PATCH",
		PathPattern:        "/bulk/certificate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CertificatePatchesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CertificatePatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for certificatePatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

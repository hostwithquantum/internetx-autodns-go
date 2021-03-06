// Code generated by go-swagger; DO NOT EDIT.

package invoice_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new invoice tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for invoice tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	InvoiceInfo(params *InvoiceInfoParams) (*InvoiceInfoOK, error)

	InvoiceList(params *InvoiceListParams) (*InvoiceListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  InvoiceInfo invoices info

  Inquiring the data for the specified Invoice.
*/
func (a *Client) InvoiceInfo(params *InvoiceInfoParams) (*InvoiceInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvoiceInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InvoiceInfo",
		Method:             "GET",
		PathPattern:        "/invoice/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InvoiceInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvoiceInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InvoiceInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  InvoiceList invoices list

  Inquiring a list of invoices with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: number, payment,  subType, amount, status, type, failed, currency, paid.
*/
func (a *Client) InvoiceList(params *InvoiceListParams) (*InvoiceListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInvoiceListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "InvoiceList",
		Method:             "POST",
		PathPattern:        "/invoice/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &InvoiceListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*InvoiceListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for InvoiceList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

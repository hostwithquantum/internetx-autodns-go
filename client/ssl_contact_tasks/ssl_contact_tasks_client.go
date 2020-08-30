// Code generated by go-swagger; DO NOT EDIT.

package ssl_contact_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ssl contact tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ssl contact tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SslContactCreate(params *SslContactCreateParams) (*SslContactCreateOK, error)

	SslContactDelete(params *SslContactDeleteParams) (*SslContactDeleteOK, error)

	SslContactList(params *SslContactListParams) (*SslContactListOK, error)

	SslContactUpdate(params *SslContactUpdateParams) (*SslContactUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SslContactCreate s s l contact create

  Creating a new SSL contact.
*/
func (a *Client) SslContactCreate(params *SslContactCreateParams) (*SslContactCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSslContactCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sslContactCreate",
		Method:             "POST",
		PathPattern:        "/sslcontact",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SslContactCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SslContactCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sslContactCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SslContactDelete s s l contact delete

  Deleting an existing SSL contact.
*/
func (a *Client) SslContactDelete(params *SslContactDeleteParams) (*SslContactDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSslContactDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sslContactDelete",
		Method:             "DELETE",
		PathPattern:        "/sslcontact/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SslContactDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SslContactDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sslContactDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SslContactList s s l contact list

  Inquiring a list of SSL contacts with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: country, fname, address, city, created, title, lname, phone, organization, state, id, fax, pcode, updated, email.
*/
func (a *Client) SslContactList(params *SslContactListParams) (*SslContactListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSslContactListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sslContactList",
		Method:             "POST",
		PathPattern:        "/sslcontact/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SslContactListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SslContactListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sslContactList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SslContactUpdate s s l contact update

  Updating an existing SSL contact.
*/
func (a *Client) SslContactUpdate(params *SslContactUpdateParams) (*SslContactUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSslContactUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sslContactUpdate",
		Method:             "PUT",
		PathPattern:        "/sslcontact/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SslContactUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SslContactUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sslContactUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

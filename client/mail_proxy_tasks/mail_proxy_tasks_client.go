// Code generated by go-swagger; DO NOT EDIT.

package mail_proxy_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new mail proxy tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for mail proxy tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	MailProxyCreate(params *MailProxyCreateParams) (*MailProxyCreateOK, error)

	MailProxyDelete(params *MailProxyDeleteParams) (*MailProxyDeleteOK, error)

	MailProxyInfo(params *MailProxyInfoParams) (*MailProxyInfoOK, error)

	MailProxyList(params *MailProxyListParams) (*MailProxyListOK, error)

	MailProxyUpdate(params *MailProxyUpdateParams) (*MailProxyUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  MailProxyCreate mails proxy create

  Creating a new mail proxy.
*/
func (a *Client) MailProxyCreate(params *MailProxyCreateParams) (*MailProxyCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyCreate",
		Method:             "POST",
		PathPattern:        "/mailProxy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MailProxyCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mailProxyCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MailProxyDelete mails proxy delete

  Deleting an existing mail proxy.
*/
func (a *Client) MailProxyDelete(params *MailProxyDeleteParams) (*MailProxyDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyDelete",
		Method:             "DELETE",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MailProxyDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mailProxyDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MailProxyInfo mails proxy info

  Inquiring the data for the specified mail proxy.
*/
func (a *Client) MailProxyInfo(params *MailProxyInfoParams) (*MailProxyInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyInfo",
		Method:             "GET",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MailProxyInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mailProxyInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MailProxyList mails proxy list

  Inquiring a list of mail proxies with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: target, admin, protection, created, updated.
*/
func (a *Client) MailProxyList(params *MailProxyListParams) (*MailProxyListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyList",
		Method:             "POST",
		PathPattern:        "/mailProxy/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MailProxyListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mailProxyList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MailProxyUpdate mails proxy update

  Updating an existing mail proxy.
*/
func (a *Client) MailProxyUpdate(params *MailProxyUpdateParams) (*MailProxyUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMailProxyUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "mailProxyUpdate",
		Method:             "PUT",
		PathPattern:        "/mailProxy/{domain}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &MailProxyUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MailProxyUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for mailProxyUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package trusted_application_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new trusted application tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for trusted application tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	TrustedAppCreate(params *TrustedAppCreateParams) (*TrustedAppCreateOK, error)

	TrustedAppDelete(params *TrustedAppDeleteParams) (*TrustedAppDeleteOK, error)

	TrustedAppInfo(params *TrustedAppInfoParams) (*TrustedAppInfoOK, error)

	TrustedAppList(params *TrustedAppListParams) (*TrustedAppListOK, error)

	TrustedAppUpdate(params *TrustedAppUpdateParams) (*TrustedAppUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  TrustedAppCreate trusteds app create

  Creating a new trusted application.
*/
func (a *Client) TrustedAppCreate(params *TrustedAppCreateParams) (*TrustedAppCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrustedAppCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trustedAppCreate",
		Method:             "POST",
		PathPattern:        "/trustedapp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrustedAppCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrustedAppCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trustedAppCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TrustedAppDelete trusteds app delete

  Deleting an existing trusted application.
*/
func (a *Client) TrustedAppDelete(params *TrustedAppDeleteParams) (*TrustedAppDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrustedAppDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trustedAppDelete",
		Method:             "DELETE",
		PathPattern:        "/trustedapp/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrustedAppDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrustedAppDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trustedAppDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TrustedAppInfo trusteds app info

  Inquiring the data for the specified trusted application.
*/
func (a *Client) TrustedAppInfo(params *TrustedAppInfoParams) (*TrustedAppInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrustedAppInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trustedAppInfo",
		Method:             "GET",
		PathPattern:        "/trustedapp/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrustedAppInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrustedAppInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trustedAppInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TrustedAppList trusteds app list

  Inquiring a list of trusted app with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: created, comment, uuid, device, updated, application.
*/
func (a *Client) TrustedAppList(params *TrustedAppListParams) (*TrustedAppListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrustedAppListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trustedAppList",
		Method:             "POST",
		PathPattern:        "/trustedapp/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrustedAppListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrustedAppListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trustedAppList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TrustedAppUpdate trusteds app update

  Updating an existing trusted application.
*/
func (a *Client) TrustedAppUpdate(params *TrustedAppUpdateParams) (*TrustedAppUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTrustedAppUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "trustedAppUpdate",
		Method:             "PUT",
		PathPattern:        "/trustedapp/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TrustedAppUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TrustedAppUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for trustedAppUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

// Code generated by go-swagger; DO NOT EDIT.

package id4me_agent_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new id4me agent tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for id4me agent tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateId4MeAgent(params *CreateId4MeAgentParams) (*CreateId4MeAgentOK, error)

	DeleteId4MeAgent(params *DeleteId4MeAgentParams) (*DeleteId4MeAgentOK, error)

	Id4MeAgentInfo(params *Id4MeAgentInfoParams) (*Id4MeAgentInfoOK, error)

	Id4MeAgentList(params *Id4MeAgentListParams) (*Id4MeAgentListOK, error)

	UpdateId4MeAgent(params *UpdateId4MeAgentParams) (*UpdateId4MeAgentOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateId4MeAgent id4mes agent create

  Creating a new id4me agent. Not available yet.
*/
func (a *Client) CreateId4MeAgent(params *CreateId4MeAgentParams) (*CreateId4MeAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateId4MeAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createId4MeAgent",
		Method:             "POST",
		PathPattern:        "/id4MeAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateId4MeAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateId4MeAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createId4MeAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteId4MeAgent id4mes agent delete

  Deleting an existing id4me agent. Not available yet.
*/
func (a *Client) DeleteId4MeAgent(params *DeleteId4MeAgentParams) (*DeleteId4MeAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteId4MeAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteId4MeAgent",
		Method:             "DELETE",
		PathPattern:        "/id4MeAgent/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteId4MeAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteId4MeAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteId4MeAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Id4MeAgentInfo id4mes agent inf

  Inquiríng the data for the specified name. Not available yet.
*/
func (a *Client) Id4MeAgentInfo(params *Id4MeAgentInfoParams) (*Id4MeAgentInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewId4MeAgentInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "id4MeAgentInfo",
		Method:             "GET",
		PathPattern:        "/id4MeAgent/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Id4MeAgentInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Id4MeAgentInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for id4MeAgentInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  Id4MeAgentList id4mes agent list

  Inquiring a list of id4Me agents with certain details. Not available yet.
*/
func (a *Client) Id4MeAgentList(params *Id4MeAgentListParams) (*Id4MeAgentListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewId4MeAgentListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "id4MeAgentList",
		Method:             "POST",
		PathPattern:        "/id4MeAgent/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &Id4MeAgentListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*Id4MeAgentListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for id4MeAgentList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateId4MeAgent id4mes agent update

  Updating an existing id4me agent. Not available yet.
*/
func (a *Client) UpdateId4MeAgent(params *UpdateId4MeAgentParams) (*UpdateId4MeAgentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateId4MeAgentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateId4MeAgent",
		Method:             "PUT",
		PathPattern:        "/id4MeAgent",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateId4MeAgentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateId4MeAgentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateId4MeAgent: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

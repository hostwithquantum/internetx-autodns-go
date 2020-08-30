// Code generated by go-swagger; DO NOT EDIT.

package subscription_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new subscription tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for subscription tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	SubscriptionCreate(params *SubscriptionCreateParams) (*SubscriptionCreateOK, error)

	SubscriptionDelte(params *SubscriptionDelteParams) (*SubscriptionDelteOK, error)

	SubscriptionList(params *SubscriptionListParams) (*SubscriptionListOK, error)

	SubscriptionUpdate(params *SubscriptionUpdateParams) (*SubscriptionUpdateOK, error)

	SubscriptionUpgrade(params *SubscriptionUpgradeParams) (*SubscriptionUpgradeOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  SubscriptionCreate subscriptions create

  Creating a new subscription contract.
*/
func (a *Client) SubscriptionCreate(params *SubscriptionCreateParams) (*SubscriptionCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscriptionCreate",
		Method:             "POST",
		PathPattern:        "/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscriptionCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscriptionDelte subscriptions delete

  Deleting a subscription contract.
*/
func (a *Client) SubscriptionDelte(params *SubscriptionDelteParams) (*SubscriptionDelteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionDelteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscriptionDelte",
		Method:             "DELETE",
		PathPattern:        "/subscription/{contractId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionDelteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionDelteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscriptionDelte: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscriptionList subscriptions list

  Inquiring a list of the subscription contracts
*/
func (a *Client) SubscriptionList(params *SubscriptionListParams) (*SubscriptionListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscriptionList",
		Method:             "POST",
		PathPattern:        "/subscription/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscriptionList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscriptionUpdate subscriptions update

  Updating a subscription contract.
*/
func (a *Client) SubscriptionUpdate(params *SubscriptionUpdateParams) (*SubscriptionUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscriptionUpdate",
		Method:             "PUT",
		PathPattern:        "/subscription/{contractId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscriptionUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SubscriptionUpgrade subscriptions upgrade

  Upgrading a subscription contract.
*/
func (a *Client) SubscriptionUpgrade(params *SubscriptionUpgradeParams) (*SubscriptionUpgradeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscriptionUpgradeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscriptionUpgrade",
		Method:             "PUT",
		PathPattern:        "/subscription/{contractId}/_upgrade",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscriptionUpgradeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscriptionUpgradeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscriptionUpgrade: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

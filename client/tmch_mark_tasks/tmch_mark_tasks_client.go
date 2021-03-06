// Code generated by go-swagger; DO NOT EDIT.

package tmch_mark_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new tmch mark tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tmch mark tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DocumentCreate(params *DocumentCreateParams) (*DocumentCreateOK, error)

	ImportAndtmchMarkTransfer(params *ImportAndtmchMarkTransferParams) (*ImportAndtmchMarkTransferOK, error)

	TmchMArkDelete(params *TmchMArkDeleteParams) (*TmchMArkDeleteOK, error)

	TmchMarkConfirm(params *TmchMarkConfirmParams) (*TmchMarkConfirmOK, error)

	TmchMarkCreate(params *TmchMarkCreateParams) (*TmchMarkCreateOK, error)

	TmchMarkDocumentCreate(params *TmchMarkDocumentCreateParams) (*TmchMarkDocumentCreateOK, error)

	TmchMarkDocumentDelete(params *TmchMarkDocumentDeleteParams) (*TmchMarkDocumentDeleteOK, error)

	TmchMarkDocumentInfo(params *TmchMarkDocumentInfoParams) (*TmchMarkDocumentInfoOK, error)

	TmchMarkImport(params *TmchMarkImportParams) (*TmchMarkImportOK, error)

	TmchMarkInfo(params *TmchMarkInfoParams) (*TmchMarkInfoOK, error)

	TmchMarkList(params *TmchMarkListParams) (*TmchMarkListOK, error)

	TmchMarkTransfer(params *TmchMarkTransferParams) (*TmchMarkTransferOK, error)

	TmchMarkUpdate(params *TmchMarkUpdateParams) (*TmchMarkUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DocumentCreate tmches mark document upload

  Uploading a specific document to a tmch mark entry.
*/
func (a *Client) DocumentCreate(params *DocumentCreateParams) (*DocumentCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDocumentCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "documentCreate",
		Method:             "PUT",
		PathPattern:        "/tmchMark/{reference}/document/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DocumentCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DocumentCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for documentCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ImportAndtmchMarkTransfer tmches mark transfer request

  Importing an existing TmchMark and starting the transfer.
*/
func (a *Client) ImportAndtmchMarkTransfer(params *ImportAndtmchMarkTransferParams) (*ImportAndtmchMarkTransferOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewImportAndtmchMarkTransferParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "importAndtmchMarkTransfer",
		Method:             "POST",
		PathPattern:        "/tmchMark/_transfer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ImportAndtmchMarkTransferReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ImportAndtmchMarkTransferOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for importAndtmchMarkTransfer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMArkDelete tmches mark delete

  Deleting the TmchMark entry for the given reference.
*/
func (a *Client) TmchMArkDelete(params *TmchMArkDeleteParams) (*TmchMArkDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMArkDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMArkDelete",
		Method:             "DELETE",
		PathPattern:        "/tmchMark/{reference}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMArkDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMArkDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMArkDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkConfirm tmches mark confirm

  Confirming an existing TmchMark.
*/
func (a *Client) TmchMarkConfirm(params *TmchMarkConfirmParams) (*TmchMarkConfirmOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkConfirmParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkConfirm",
		Method:             "POST",
		PathPattern:        "/tmchMark/{reference}/_confirm",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkConfirmReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkConfirmOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkConfirm: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkCreate tmches mark create

  Creating a new TmchMark.
*/
func (a *Client) TmchMarkCreate(params *TmchMarkCreateParams) (*TmchMarkCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkCreate",
		Method:             "POST",
		PathPattern:        "/tmchMark",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkDocumentCreate tmches mark document create

  Assinging an uploaded document to a tmch mark entry.
*/
func (a *Client) TmchMarkDocumentCreate(params *TmchMarkDocumentCreateParams) (*TmchMarkDocumentCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkDocumentCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkDocumentCreate",
		Method:             "POST",
		PathPattern:        "/tmchMark/{reference}/document",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkDocumentCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkDocumentCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkDocumentCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkDocumentDelete tmches mark document delete

  Deleting a single TmchMark document for the given type.
*/
func (a *Client) TmchMarkDocumentDelete(params *TmchMarkDocumentDeleteParams) (*TmchMarkDocumentDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkDocumentDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkDocumentDelete",
		Method:             "DELETE",
		PathPattern:        "/tmchMark/{reference}/document/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkDocumentDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkDocumentDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkDocumentDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkDocumentInfo tmches mark document info

  Inquiring a single TmchMark document for the given type.
*/
func (a *Client) TmchMarkDocumentInfo(params *TmchMarkDocumentInfoParams) (*TmchMarkDocumentInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkDocumentInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkDocumentInfo",
		Method:             "GET",
		PathPattern:        "/tmchMark/{reference}/document/{type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkDocumentInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkDocumentInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkDocumentInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkImport tmches mark import

  Importing an existing TmchMark.
*/
func (a *Client) TmchMarkImport(params *TmchMarkImportParams) (*TmchMarkImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkImportParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkImport",
		Method:             "POST",
		PathPattern:        "/tmchMark/_import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkImportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkImport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkInfo tmches mark info

  -inquiring the TmchMark entry for the given reference.
*/
func (a *Client) TmchMarkInfo(params *TmchMarkInfoParams) (*TmchMarkInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkInfo",
		Method:             "GET",
		PathPattern:        "/tmchMark/{reference}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkList tmches mark list

  Inquiring a list of TmchMark with certain details. The following keys can be used for filtering, ordering and fetching additional data via query parameter: type, name, reference, description, period, renew, status, payable, sent.
*/
func (a *Client) TmchMarkList(params *TmchMarkListParams) (*TmchMarkListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkList",
		Method:             "POST",
		PathPattern:        "/tmchMark/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkTransfer tmches mark transfer request

  Starting the transfer for given reference (External TmchMark).
*/
func (a *Client) TmchMarkTransfer(params *TmchMarkTransferParams) (*TmchMarkTransferOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkTransferParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkTransfer",
		Method:             "PUT",
		PathPattern:        "/tmchMark/{reference}/_transfer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkTransferReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkTransferOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkTransfer: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TmchMarkUpdate tmches mark update

  Updating the TmchMark entry for the given reference.
*/
func (a *Client) TmchMarkUpdate(params *TmchMarkUpdateParams) (*TmchMarkUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTmchMarkUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "tmchMarkUpdate",
		Method:             "PUT",
		PathPattern:        "/tmchMark/{reference}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &TmchMarkUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TmchMarkUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for tmchMarkUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

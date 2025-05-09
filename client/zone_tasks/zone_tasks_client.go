// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new zone tasks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new zone tasks API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new zone tasks API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for zone tasks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ZoneAddDomainSafe(params *ZoneAddDomainSafeParams, opts ...ClientOption) (*ZoneAddDomainSafeOK, error)

	ZoneAxfr(params *ZoneAxfrParams, opts ...ClientOption) (*ZoneAxfrOK, error)

	ZoneBulkPatches(params *ZoneBulkPatchesParams, opts ...ClientOption) (*ZoneBulkPatchesOK, error)

	ZoneCopy(params *ZoneCopyParams, opts ...ClientOption) (*ZoneCopyOK, error)

	ZoneCreate(params *ZoneCreateParams, opts ...ClientOption) (*ZoneCreateOK, error)

	ZoneCreates(params *ZoneCreatesParams, opts ...ClientOption) (*ZoneCreatesOK, error)

	ZoneDelete(params *ZoneDeleteParams, opts ...ClientOption) (*ZoneDeleteOK, error)

	ZoneDeleteDomainSafe(params *ZoneDeleteDomainSafeParams, opts ...ClientOption) (*ZoneDeleteDomainSafeOK, error)

	ZoneDeletes(params *ZoneDeletesParams, opts ...ClientOption) (*ZoneDeletesOK, error)

	ZoneHistoryInfo(params *ZoneHistoryInfoParams, opts ...ClientOption) (*ZoneHistoryInfoOK, error)

	ZoneHistoryList(params *ZoneHistoryListParams, opts ...ClientOption) (*ZoneHistoryListOK, error)

	ZoneImport(params *ZoneImportParams, opts ...ClientOption) (*ZoneImportOK, error)

	ZoneImports(params *ZoneImportsParams, opts ...ClientOption) (*ZoneImportsOK, error)

	ZoneInfo(params *ZoneInfoParams, opts ...ClientOption) (*ZoneInfoOK, error)

	ZoneList(params *ZoneListParams, opts ...ClientOption) (*ZoneListOK, error)

	ZoneMigrate(params *ZoneMigrateParams, opts ...ClientOption) (*ZoneMigrateOK, error)

	ZonePatch(params *ZonePatchParams, opts ...ClientOption) (*ZonePatchOK, error)

	ZoneRestore(params *ZoneRestoreParams, opts ...ClientOption) (*ZoneRestoreOK, error)

	ZoneRestores(params *ZoneRestoresParams, opts ...ClientOption) (*ZoneRestoresOK, error)

	ZoneStream(params *ZoneStreamParams, opts ...ClientOption) (*ZoneStreamOK, error)

	ZoneUpdate(params *ZoneUpdateParams, opts ...ClientOption) (*ZoneUpdateOK, error)

	ZoneUpdateComment(params *ZoneUpdateCommentParams, opts ...ClientOption) (*ZoneUpdateCommentOK, error)

	ZoneUpdateComments(params *ZoneUpdateCommentsParams, opts ...ClientOption) (*ZoneUpdateCommentsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ZoneAddDomainSafe saves object create 0601

Adding the zone to the domain safe
*/
func (a *Client) ZoneAddDomainSafe(params *ZoneAddDomainSafeParams, opts ...ClientOption) (*ZoneAddDomainSafeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneAddDomainSafeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneAddDomainSafe",
		Method:             "PUT",
		PathPattern:        "/zone/{name}/{systemNameServer}/_domainSafe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneAddDomainSafeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneAddDomainSafeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneAddDomainSafe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneAxfr zones axfr 0210

Inquiring the AXFR data for the specified zone.
*/
func (a *Client) ZoneAxfr(params *ZoneAxfrParams, opts ...ClientOption) (*ZoneAxfrOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneAxfrParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneAxfr",
		Method:             "GET",
		PathPattern:        "/zone/{name}/{systemNameServer}/_axfr",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneAxfrReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneAxfrOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneAxfr: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneBulkPatches updates zones 0202001

Updating several existing zones with one request.
*/
func (a *Client) ZoneBulkPatches(params *ZoneBulkPatchesParams, opts ...ClientOption) (*ZoneBulkPatchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneBulkPatchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneBulkPatches",
		Method:             "PATCH",
		PathPattern:        "/bulk/zone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneBulkPatchesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneBulkPatchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneBulkPatches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneCopy zones update 0212

Copying an existing zone.
*/
func (a *Client) ZoneCopy(params *ZoneCopyParams, opts ...ClientOption) (*ZoneCopyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneCopyParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneCopy",
		Method:             "PUT",
		PathPattern:        "/zone/{name}/{systemNameServer}/_copy",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneCopyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneCopyOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneCopy: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneCreate zones create 0201

Creating a zone.
*/
func (a *Client) ZoneCreate(params *ZoneCreateParams, opts ...ClientOption) (*ZoneCreateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneCreate",
		Method:             "POST",
		PathPattern:        "/zone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneCreateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneCreates creates zones 0201

Creating several zones with one request.
*/
func (a *Client) ZoneCreates(params *ZoneCreatesParams, opts ...ClientOption) (*ZoneCreatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneCreatesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneCreates",
		Method:             "POST",
		PathPattern:        "/bulk/zone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneCreatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneCreatesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneCreates: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneDelete zones delete 0203

Deleting an existing zone.
*/
func (a *Client) ZoneDelete(params *ZoneDeleteParams, opts ...ClientOption) (*ZoneDeleteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneDelete",
		Method:             "DELETE",
		PathPattern:        "/zone/{name}/{systemNameServer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneDeleteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneDeleteDomainSafe saves object delete 0603

Deleting the zone from the domain safe
*/
func (a *Client) ZoneDeleteDomainSafe(params *ZoneDeleteDomainSafeParams, opts ...ClientOption) (*ZoneDeleteDomainSafeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneDeleteDomainSafeParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneDeleteDomainSafe",
		Method:             "DELETE",
		PathPattern:        "/zone/{name}/{systemNameServer}/_domainSafe",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneDeleteDomainSafeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneDeleteDomainSafeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneDeleteDomainSafe: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneDeletes deletes zones 0203

Deleting several existing zones with one request.
*/
func (a *Client) ZoneDeletes(params *ZoneDeletesParams, opts ...ClientOption) (*ZoneDeletesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneDeletesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneDeletes",
		Method:             "DELETE",
		PathPattern:        "/bulk/zone",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneDeletesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneDeletesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneDeletes: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneHistoryInfo zones history info 0224

Inquiring the data for the specified log.
*/
func (a *Client) ZoneHistoryInfo(params *ZoneHistoryInfoParams, opts ...ClientOption) (*ZoneHistoryInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneHistoryInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneHistoryInfo",
		Method:             "GET",
		PathPattern:        "/zone/history/{logId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneHistoryInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneHistoryInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneHistoryInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneHistoryList zones history list 0225

Inquiring a list of zones history entries with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: dnssec, created, mainip, secondary1, secondary2, secondary3, secondary4, secondary5, secondary6, secondary7, virtualNameServer, domainsafe, name, comment, updated, action, primary, changed.
*/
func (a *Client) ZoneHistoryList(params *ZoneHistoryListParams, opts ...ClientOption) (*ZoneHistoryListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneHistoryListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneHistoryList",
		Method:             "POST",
		PathPattern:        "/zone/history/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneHistoryListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneHistoryListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneHistoryList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneImport zones import 0204

Importing the specified zone.
*/
func (a *Client) ZoneImport(params *ZoneImportParams, opts ...ClientOption) (*ZoneImportOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneImportParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneImport",
		Method:             "POST",
		PathPattern:        "/zone/{name}/{systemNameServer}/_import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneImportReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneImportOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneImport: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneImports imports zones 0204

Importing several specified zones with one request.
*/
func (a *Client) ZoneImports(params *ZoneImportsParams, opts ...ClientOption) (*ZoneImportsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneImportsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneImports",
		Method:             "POST",
		PathPattern:        "/bulk/zone/_import",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneImportsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneImportsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneImports: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneInfo zones info 0205

Inquiring the data for the specified zone.
*/
func (a *Client) ZoneInfo(params *ZoneInfoParams, opts ...ClientOption) (*ZoneInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneInfo",
		Method:             "GET",
		PathPattern:        "/zone/{name}/{systemNameServer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneInfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneInfo: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneList zones list 0205

Inquiring a list of zones with certain details. The following keys can be used for filtering, ordering and inquiring additional data via query parameter: dnssec, created, mainip, secondary1, secondary2, secondary3, secondary4, secondary5, secondary6, secondary7, virtualNameServer, domainsafe, name, comment, updated, action, primary, changed.
*/
func (a *Client) ZoneList(params *ZoneListParams, opts ...ClientOption) (*ZoneListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneList",
		Method:             "POST",
		PathPattern:        "/zone/_search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneMigrate zones update 0212

Copying an existing zone and updating the domain with the new name servers.
*/
func (a *Client) ZoneMigrate(params *ZoneMigrateParams, opts ...ClientOption) (*ZoneMigrateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneMigrateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneMigrate",
		Method:             "PUT",
		PathPattern:        "/zone/{name}/{systemNameServer}/_migrate",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneMigrateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneMigrateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneMigrate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZonePatch zones update 0202001

Updating an existing zone.
*/
func (a *Client) ZonePatch(params *ZonePatchParams, opts ...ClientOption) (*ZonePatchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZonePatchParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zonePatch",
		Method:             "PATCH",
		PathPattern:        "/zone/{name}/{systemNameServer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZonePatchReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZonePatchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zonePatch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneRestore zones restore 0226

Restoring the specified zone.
*/
func (a *Client) ZoneRestore(params *ZoneRestoreParams, opts ...ClientOption) (*ZoneRestoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneRestoreParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneRestore",
		Method:             "POST",
		PathPattern:        "/zone/{name}/{systemNameServer}/_restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneRestoreReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneRestoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneRestore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneRestores restores zones 0226

Restoring several specified zones with one request.
*/
func (a *Client) ZoneRestores(params *ZoneRestoresParams, opts ...ClientOption) (*ZoneRestoresOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneRestoresParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneRestores",
		Method:             "POST",
		PathPattern:        "/bulk/zone/_restore",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneRestoresReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneRestoresOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneRestores: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneStream zones stream update 0202002

Adding or removing records for any zone with the given name.
*/
func (a *Client) ZoneStream(params *ZoneStreamParams, opts ...ClientOption) (*ZoneStreamOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneStreamParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneStream",
		Method:             "POST",
		PathPattern:        "/zone/{name}/_stream",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneStreamReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneStreamOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneStream: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneUpdate zones update 0202

Updating an existing zone.
*/
func (a *Client) ZoneUpdate(params *ZoneUpdateParams, opts ...ClientOption) (*ZoneUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneUpdate",
		Method:             "PUT",
		PathPattern:        "/zone/{name}/{systemNameServer}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneUpdateComment zones comment update 0202004

Updating an existing zone.
*/
func (a *Client) ZoneUpdateComment(params *ZoneUpdateCommentParams, opts ...ClientOption) (*ZoneUpdateCommentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneUpdateCommentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneUpdateComment",
		Method:             "PUT",
		PathPattern:        "/zone/{name}/{systemNameServer}/_comment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneUpdateCommentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneUpdateCommentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneUpdateComment: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ZoneUpdateComments updates zone comments 0202004

Updating several existing zone comments with one request.
*/
func (a *Client) ZoneUpdateComments(params *ZoneUpdateCommentsParams, opts ...ClientOption) (*ZoneUpdateCommentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewZoneUpdateCommentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "zoneUpdateComments",
		Method:             "PUT",
		PathPattern:        "/bulk/zone/_comment",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ZoneUpdateCommentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ZoneUpdateCommentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for zoneUpdateComments: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}

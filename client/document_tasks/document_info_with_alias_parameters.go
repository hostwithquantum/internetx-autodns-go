// Code generated by go-swagger; DO NOT EDIT.

package document_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewDocumentInfoWithAliasParams creates a new DocumentInfoWithAliasParams object
// with the default values initialized.
func NewDocumentInfoWithAliasParams() *DocumentInfoWithAliasParams {
	var ()
	return &DocumentInfoWithAliasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentInfoWithAliasParamsWithTimeout creates a new DocumentInfoWithAliasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDocumentInfoWithAliasParamsWithTimeout(timeout time.Duration) *DocumentInfoWithAliasParams {
	var ()
	return &DocumentInfoWithAliasParams{

		timeout: timeout,
	}
}

// NewDocumentInfoWithAliasParamsWithContext creates a new DocumentInfoWithAliasParams object
// with the default values initialized, and the ability to set a context for a request
func NewDocumentInfoWithAliasParamsWithContext(ctx context.Context) *DocumentInfoWithAliasParams {
	var ()
	return &DocumentInfoWithAliasParams{

		Context: ctx,
	}
}

// NewDocumentInfoWithAliasParamsWithHTTPClient creates a new DocumentInfoWithAliasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDocumentInfoWithAliasParamsWithHTTPClient(client *http.Client) *DocumentInfoWithAliasParams {
	var ()
	return &DocumentInfoWithAliasParams{
		HTTPClient: client,
	}
}

/*DocumentInfoWithAliasParams contains all the parameters to send to the API endpoint
for the document info with alias operation typically these are written to a http.Request
*/
type DocumentInfoWithAliasParams struct {

	/*XDomainrobotBulkLimit*/
	XDomainrobotBulkLimit *int32
	/*XDomainrobotContext*/
	XDomainrobotContext *int32
	/*XDomainrobotDemo*/
	XDomainrobotDemo *bool
	/*XDomainrobotDomainSafePin*/
	XDomainrobotDomainSafePin *string
	/*XDomainrobotDomainSafeTan*/
	XDomainrobotDomainSafeTan *string
	/*XDomainrobotDomainSafeTransaction*/
	XDomainrobotDomainSafeTransaction *string
	/*XDomainrobotDomainSafeTransactionExpire*/
	XDomainrobotDomainSafeTransactionExpire *strfmt.DateTime
	/*XDomainrobotOwnerContext*/
	XDomainrobotOwnerContext *int32
	/*XDomainrobotOwnerUser*/
	XDomainrobotOwnerUser *string
	/*XDomainrobotSessionID*/
	XDomainrobotSessionID *string
	/*XDomainrobotWS*/
	XDomainrobotWS *string
	/*Alias*/
	Alias string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithTimeout(timeout time.Duration) *DocumentInfoWithAliasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithContext(ctx context.Context) *DocumentInfoWithAliasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithHTTPClient(client *http.Client) *DocumentInfoWithAliasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithXDomainrobotWS(xDomainrobotWS *string) *DocumentInfoWithAliasParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithAlias adds the alias to the document info with alias params
func (o *DocumentInfoWithAliasParams) WithAlias(alias string) *DocumentInfoWithAliasParams {
	o.SetAlias(alias)
	return o
}

// SetAlias adds the alias to the document info with alias params
func (o *DocumentInfoWithAliasParams) SetAlias(alias string) {
	o.Alias = alias
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentInfoWithAliasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDomainrobotBulkLimit != nil {

		// header param X-Domainrobot-Bulk-Limit
		if err := r.SetHeaderParam("X-Domainrobot-Bulk-Limit", swag.FormatInt32(*o.XDomainrobotBulkLimit)); err != nil {
			return err
		}

	}

	if o.XDomainrobotContext != nil {

		// header param X-Domainrobot-Context
		if err := r.SetHeaderParam("X-Domainrobot-Context", swag.FormatInt32(*o.XDomainrobotContext)); err != nil {
			return err
		}

	}

	if o.XDomainrobotDemo != nil {

		// header param X-Domainrobot-Demo
		if err := r.SetHeaderParam("X-Domainrobot-Demo", swag.FormatBool(*o.XDomainrobotDemo)); err != nil {
			return err
		}

	}

	if o.XDomainrobotDomainSafePin != nil {

		// header param X-Domainrobot-Domain-Safe-Pin
		if err := r.SetHeaderParam("X-Domainrobot-Domain-Safe-Pin", *o.XDomainrobotDomainSafePin); err != nil {
			return err
		}

	}

	if o.XDomainrobotDomainSafeTan != nil {

		// header param X-Domainrobot-Domain-Safe-Tan
		if err := r.SetHeaderParam("X-Domainrobot-Domain-Safe-Tan", *o.XDomainrobotDomainSafeTan); err != nil {
			return err
		}

	}

	if o.XDomainrobotDomainSafeTransaction != nil {

		// header param X-Domainrobot-Domain-Safe-Transaction
		if err := r.SetHeaderParam("X-Domainrobot-Domain-Safe-Transaction", *o.XDomainrobotDomainSafeTransaction); err != nil {
			return err
		}

	}

	if o.XDomainrobotDomainSafeTransactionExpire != nil {

		// header param X-Domainrobot-Domain-Safe-Transaction-Expire
		if err := r.SetHeaderParam("X-Domainrobot-Domain-Safe-Transaction-Expire", o.XDomainrobotDomainSafeTransactionExpire.String()); err != nil {
			return err
		}

	}

	if o.XDomainrobotOwnerContext != nil {

		// header param X-Domainrobot-Owner-Context
		if err := r.SetHeaderParam("X-Domainrobot-Owner-Context", swag.FormatInt32(*o.XDomainrobotOwnerContext)); err != nil {
			return err
		}

	}

	if o.XDomainrobotOwnerUser != nil {

		// header param X-Domainrobot-Owner-User
		if err := r.SetHeaderParam("X-Domainrobot-Owner-User", *o.XDomainrobotOwnerUser); err != nil {
			return err
		}

	}

	if o.XDomainrobotSessionID != nil {

		// header param X-Domainrobot-SessionId
		if err := r.SetHeaderParam("X-Domainrobot-SessionId", *o.XDomainrobotSessionID); err != nil {
			return err
		}

	}

	if o.XDomainrobotWS != nil {

		// header param X-Domainrobot-WS
		if err := r.SetHeaderParam("X-Domainrobot-WS", *o.XDomainrobotWS); err != nil {
			return err
		}

	}

	// path param alias
	if err := r.SetPathParam("alias", o.Alias); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

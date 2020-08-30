// Code generated by go-swagger; DO NOT EDIT.

package redirect_tasks

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

// NewRedirectDeleteParams creates a new RedirectDeleteParams object
// with the default values initialized.
func NewRedirectDeleteParams() *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectDeleteParamsWithTimeout creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRedirectDeleteParamsWithTimeout(timeout time.Duration) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		timeout: timeout,
	}
}

// NewRedirectDeleteParamsWithContext creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewRedirectDeleteParamsWithContext(ctx context.Context) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{

		Context: ctx,
	}
}

// NewRedirectDeleteParamsWithHTTPClient creates a new RedirectDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRedirectDeleteParamsWithHTTPClient(client *http.Client) *RedirectDeleteParams {
	var ()
	return &RedirectDeleteParams{
		HTTPClient: client,
	}
}

/*RedirectDeleteParams contains all the parameters to send to the API endpoint
for the redirect delete operation typically these are written to a http.Request
*/
type RedirectDeleteParams struct {

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
	/*Keys
	  If the dns should be provisioned if available.

	*/
	Keys []string
	/*Source*/
	Source string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the redirect delete params
func (o *RedirectDeleteParams) WithTimeout(timeout time.Duration) *RedirectDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect delete params
func (o *RedirectDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect delete params
func (o *RedirectDeleteParams) WithContext(ctx context.Context) *RedirectDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect delete params
func (o *RedirectDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect delete params
func (o *RedirectDeleteParams) WithHTTPClient(client *http.Client) *RedirectDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect delete params
func (o *RedirectDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *RedirectDeleteParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotContext(xDomainrobotContext *int32) *RedirectDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *RedirectDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *RedirectDeleteParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *RedirectDeleteParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *RedirectDeleteParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *RedirectDeleteParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *RedirectDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *RedirectDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *RedirectDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the redirect delete params
func (o *RedirectDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *RedirectDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the redirect delete params
func (o *RedirectDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithKeys adds the keys to the redirect delete params
func (o *RedirectDeleteParams) WithKeys(keys []string) *RedirectDeleteParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the redirect delete params
func (o *RedirectDeleteParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithSource adds the source to the redirect delete params
func (o *RedirectDeleteParams) WithSource(source string) *RedirectDeleteParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the redirect delete params
func (o *RedirectDeleteParams) SetSource(source string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

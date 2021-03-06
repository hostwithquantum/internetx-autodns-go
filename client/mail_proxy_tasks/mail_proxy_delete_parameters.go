// Code generated by go-swagger; DO NOT EDIT.

package mail_proxy_tasks

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

// NewMailProxyDeleteParams creates a new MailProxyDeleteParams object
// with the default values initialized.
func NewMailProxyDeleteParams() *MailProxyDeleteParams {
	var ()
	return &MailProxyDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMailProxyDeleteParamsWithTimeout creates a new MailProxyDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMailProxyDeleteParamsWithTimeout(timeout time.Duration) *MailProxyDeleteParams {
	var ()
	return &MailProxyDeleteParams{

		timeout: timeout,
	}
}

// NewMailProxyDeleteParamsWithContext creates a new MailProxyDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewMailProxyDeleteParamsWithContext(ctx context.Context) *MailProxyDeleteParams {
	var ()
	return &MailProxyDeleteParams{

		Context: ctx,
	}
}

// NewMailProxyDeleteParamsWithHTTPClient creates a new MailProxyDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMailProxyDeleteParamsWithHTTPClient(client *http.Client) *MailProxyDeleteParams {
	var ()
	return &MailProxyDeleteParams{
		HTTPClient: client,
	}
}

/*MailProxyDeleteParams contains all the parameters to send to the API endpoint
for the mail proxy delete operation typically these are written to a http.Request
*/
type MailProxyDeleteParams struct {

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
	/*Domain*/
	Domain string
	/*Keys
	  If the dns should be provisioned if available.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the mail proxy delete params
func (o *MailProxyDeleteParams) WithTimeout(timeout time.Duration) *MailProxyDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mail proxy delete params
func (o *MailProxyDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mail proxy delete params
func (o *MailProxyDeleteParams) WithContext(ctx context.Context) *MailProxyDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mail proxy delete params
func (o *MailProxyDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mail proxy delete params
func (o *MailProxyDeleteParams) WithHTTPClient(client *http.Client) *MailProxyDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mail proxy delete params
func (o *MailProxyDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *MailProxyDeleteParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotContext(xDomainrobotContext *int32) *MailProxyDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *MailProxyDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *MailProxyDeleteParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *MailProxyDeleteParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *MailProxyDeleteParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *MailProxyDeleteParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *MailProxyDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *MailProxyDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *MailProxyDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the mail proxy delete params
func (o *MailProxyDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *MailProxyDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the mail proxy delete params
func (o *MailProxyDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithDomain adds the domain to the mail proxy delete params
func (o *MailProxyDeleteParams) WithDomain(domain string) *MailProxyDeleteParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the mail proxy delete params
func (o *MailProxyDeleteParams) SetDomain(domain string) {
	o.Domain = domain
}

// WithKeys adds the keys to the mail proxy delete params
func (o *MailProxyDeleteParams) WithKeys(keys []string) *MailProxyDeleteParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the mail proxy delete params
func (o *MailProxyDeleteParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *MailProxyDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param domain
	if err := r.SetPathParam("domain", o.Domain); err != nil {
		return err
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewMailProxyInfoParams creates a new MailProxyInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMailProxyInfoParams() *MailProxyInfoParams {
	return &MailProxyInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMailProxyInfoParamsWithTimeout creates a new MailProxyInfoParams object
// with the ability to set a timeout on a request.
func NewMailProxyInfoParamsWithTimeout(timeout time.Duration) *MailProxyInfoParams {
	return &MailProxyInfoParams{
		timeout: timeout,
	}
}

// NewMailProxyInfoParamsWithContext creates a new MailProxyInfoParams object
// with the ability to set a context for a request.
func NewMailProxyInfoParamsWithContext(ctx context.Context) *MailProxyInfoParams {
	return &MailProxyInfoParams{
		Context: ctx,
	}
}

// NewMailProxyInfoParamsWithHTTPClient creates a new MailProxyInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewMailProxyInfoParamsWithHTTPClient(client *http.Client) *MailProxyInfoParams {
	return &MailProxyInfoParams{
		HTTPClient: client,
	}
}

/*
MailProxyInfoParams contains all the parameters to send to the API endpoint

	for the mail proxy info operation.

	Typically these are written to a http.Request.
*/
type MailProxyInfoParams struct {

	// XDomainrobot2FAToken.
	//
	// Format: int32
	XDomainrobot2FAToken *int32

	// XDomainrobotBulkLimit.
	//
	// Format: int32
	XDomainrobotBulkLimit *int32

	// XDomainrobotContext.
	//
	// Format: int32
	XDomainrobotContext *int32

	// XDomainrobotDemo.
	XDomainrobotDemo *bool

	// XDomainrobotDomainSafePin.
	XDomainrobotDomainSafePin *string

	// XDomainrobotDomainSafeTan.
	XDomainrobotDomainSafeTan *string

	// XDomainrobotDomainSafeTransaction.
	XDomainrobotDomainSafeTransaction *string

	// XDomainrobotDomainSafeTransactionExpire.
	//
	// Format: date-time
	XDomainrobotDomainSafeTransactionExpire *strfmt.DateTime

	// XDomainrobotOwnerContext.
	//
	// Format: int32
	XDomainrobotOwnerContext *int32

	// XDomainrobotOwnerUser.
	XDomainrobotOwnerUser *string

	// XDomainrobotSessionID.
	XDomainrobotSessionID *string

	// XDomainrobotWS.
	XDomainrobotWS *string

	// Domain.
	Domain string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the mail proxy info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailProxyInfoParams) WithDefaults() *MailProxyInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the mail proxy info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MailProxyInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the mail proxy info params
func (o *MailProxyInfoParams) WithTimeout(timeout time.Duration) *MailProxyInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the mail proxy info params
func (o *MailProxyInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the mail proxy info params
func (o *MailProxyInfoParams) WithContext(ctx context.Context) *MailProxyInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the mail proxy info params
func (o *MailProxyInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the mail proxy info params
func (o *MailProxyInfoParams) WithHTTPClient(client *http.Client) *MailProxyInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the mail proxy info params
func (o *MailProxyInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *MailProxyInfoParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *MailProxyInfoParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotContext(xDomainrobotContext *int32) *MailProxyInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *MailProxyInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *MailProxyInfoParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *MailProxyInfoParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *MailProxyInfoParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *MailProxyInfoParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *MailProxyInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *MailProxyInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *MailProxyInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the mail proxy info params
func (o *MailProxyInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *MailProxyInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the mail proxy info params
func (o *MailProxyInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithDomain adds the domain to the mail proxy info params
func (o *MailProxyInfoParams) WithDomain(domain string) *MailProxyInfoParams {
	o.SetDomain(domain)
	return o
}

// SetDomain adds the domain to the mail proxy info params
func (o *MailProxyInfoParams) SetDomain(domain string) {
	o.Domain = domain
}

// WriteToRequest writes these params to a swagger request
func (o *MailProxyInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XDomainrobot2FAToken != nil {

		// header param X-Domainrobot-2FA-Token
		if err := r.SetHeaderParam("X-Domainrobot-2FA-Token", swag.FormatInt32(*o.XDomainrobot2FAToken)); err != nil {
			return err
		}
	}

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package session_tasks

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

// NewLogoutParams creates a new LogoutParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLogoutParams() *LogoutParams {
	return &LogoutParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLogoutParamsWithTimeout creates a new LogoutParams object
// with the ability to set a timeout on a request.
func NewLogoutParamsWithTimeout(timeout time.Duration) *LogoutParams {
	return &LogoutParams{
		timeout: timeout,
	}
}

// NewLogoutParamsWithContext creates a new LogoutParams object
// with the ability to set a context for a request.
func NewLogoutParamsWithContext(ctx context.Context) *LogoutParams {
	return &LogoutParams{
		Context: ctx,
	}
}

// NewLogoutParamsWithHTTPClient creates a new LogoutParams object
// with the ability to set a custom HTTPClient for a request.
func NewLogoutParamsWithHTTPClient(client *http.Client) *LogoutParams {
	return &LogoutParams{
		HTTPClient: client,
	}
}

/*
LogoutParams contains all the parameters to send to the API endpoint

	for the logout operation.

	Typically these are written to a http.Request.
*/
type LogoutParams struct {

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogoutParams) WithDefaults() *LogoutParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the logout params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LogoutParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the logout params
func (o *LogoutParams) WithTimeout(timeout time.Duration) *LogoutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the logout params
func (o *LogoutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the logout params
func (o *LogoutParams) WithContext(ctx context.Context) *LogoutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the logout params
func (o *LogoutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the logout params
func (o *LogoutParams) WithHTTPClient(client *http.Client) *LogoutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the logout params
func (o *LogoutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the logout params
func (o *LogoutParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *LogoutParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the logout params
func (o *LogoutParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the logout params
func (o *LogoutParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *LogoutParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the logout params
func (o *LogoutParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the logout params
func (o *LogoutParams) WithXDomainrobotContext(xDomainrobotContext *int32) *LogoutParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the logout params
func (o *LogoutParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the logout params
func (o *LogoutParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *LogoutParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the logout params
func (o *LogoutParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the logout params
func (o *LogoutParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *LogoutParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the logout params
func (o *LogoutParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the logout params
func (o *LogoutParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *LogoutParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the logout params
func (o *LogoutParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the logout params
func (o *LogoutParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *LogoutParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the logout params
func (o *LogoutParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the logout params
func (o *LogoutParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *LogoutParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the logout params
func (o *LogoutParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the logout params
func (o *LogoutParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *LogoutParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the logout params
func (o *LogoutParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the logout params
func (o *LogoutParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *LogoutParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the logout params
func (o *LogoutParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the logout params
func (o *LogoutParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *LogoutParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the logout params
func (o *LogoutParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the logout params
func (o *LogoutParams) WithXDomainrobotWS(xDomainrobotWS *string) *LogoutParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the logout params
func (o *LogoutParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WriteToRequest writes these params to a swagger request
func (o *LogoutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

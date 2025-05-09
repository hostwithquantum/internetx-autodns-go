// Code generated by go-swagger; DO NOT EDIT.

package user_2fa_tasks

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

// NewUser2faDeleteParams creates a new User2faDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUser2faDeleteParams() *User2faDeleteParams {
	return &User2faDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUser2faDeleteParamsWithTimeout creates a new User2faDeleteParams object
// with the ability to set a timeout on a request.
func NewUser2faDeleteParamsWithTimeout(timeout time.Duration) *User2faDeleteParams {
	return &User2faDeleteParams{
		timeout: timeout,
	}
}

// NewUser2faDeleteParamsWithContext creates a new User2faDeleteParams object
// with the ability to set a context for a request.
func NewUser2faDeleteParamsWithContext(ctx context.Context) *User2faDeleteParams {
	return &User2faDeleteParams{
		Context: ctx,
	}
}

// NewUser2faDeleteParamsWithHTTPClient creates a new User2faDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewUser2faDeleteParamsWithHTTPClient(client *http.Client) *User2faDeleteParams {
	return &User2faDeleteParams{
		HTTPClient: client,
	}
}

/*
User2faDeleteParams contains all the parameters to send to the API endpoint

	for the user2fa delete operation.

	Typically these are written to a http.Request.
*/
type User2faDeleteParams struct {

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

// WithDefaults hydrates default values in the user2fa delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *User2faDeleteParams) WithDefaults() *User2faDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the user2fa delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *User2faDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the user2fa delete params
func (o *User2faDeleteParams) WithTimeout(timeout time.Duration) *User2faDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user2fa delete params
func (o *User2faDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user2fa delete params
func (o *User2faDeleteParams) WithContext(ctx context.Context) *User2faDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user2fa delete params
func (o *User2faDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user2fa delete params
func (o *User2faDeleteParams) WithHTTPClient(client *http.Client) *User2faDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user2fa delete params
func (o *User2faDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *User2faDeleteParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *User2faDeleteParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotContext(xDomainrobotContext *int32) *User2faDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *User2faDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *User2faDeleteParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *User2faDeleteParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *User2faDeleteParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *User2faDeleteParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *User2faDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *User2faDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *User2faDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the user2fa delete params
func (o *User2faDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *User2faDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the user2fa delete params
func (o *User2faDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WriteToRequest writes these params to a swagger request
func (o *User2faDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

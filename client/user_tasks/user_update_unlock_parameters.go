// Code generated by go-swagger; DO NOT EDIT.

package user_tasks

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

// NewUserUpdateUnlockParams creates a new UserUpdateUnlockParams object
// with the default values initialized.
func NewUserUpdateUnlockParams() *UserUpdateUnlockParams {
	var ()
	return &UserUpdateUnlockParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUserUpdateUnlockParamsWithTimeout creates a new UserUpdateUnlockParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUserUpdateUnlockParamsWithTimeout(timeout time.Duration) *UserUpdateUnlockParams {
	var ()
	return &UserUpdateUnlockParams{

		timeout: timeout,
	}
}

// NewUserUpdateUnlockParamsWithContext creates a new UserUpdateUnlockParams object
// with the default values initialized, and the ability to set a context for a request
func NewUserUpdateUnlockParamsWithContext(ctx context.Context) *UserUpdateUnlockParams {
	var ()
	return &UserUpdateUnlockParams{

		Context: ctx,
	}
}

// NewUserUpdateUnlockParamsWithHTTPClient creates a new UserUpdateUnlockParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUserUpdateUnlockParamsWithHTTPClient(client *http.Client) *UserUpdateUnlockParams {
	var ()
	return &UserUpdateUnlockParams{
		HTTPClient: client,
	}
}

/*UserUpdateUnlockParams contains all the parameters to send to the API endpoint
for the user update unlock operation typically these are written to a http.Request
*/
type UserUpdateUnlockParams struct {

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
	/*Context*/
	Context string
	/*Keys
	  Specifies whether the children should be unlocked as well. Example : ?keys[]=children

	*/
	Keys []string
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the user update unlock params
func (o *UserUpdateUnlockParams) WithTimeout(timeout time.Duration) *UserUpdateUnlockParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the user update unlock params
func (o *UserUpdateUnlockParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the user update unlock params
func (o *UserUpdateUnlockParams) WithContext(ctx context.Context) *UserUpdateUnlockParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the user update unlock params
func (o *UserUpdateUnlockParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the user update unlock params
func (o *UserUpdateUnlockParams) WithHTTPClient(client *http.Client) *UserUpdateUnlockParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the user update unlock params
func (o *UserUpdateUnlockParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *UserUpdateUnlockParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotContext(xDomainrobotContext *int32) *UserUpdateUnlockParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *UserUpdateUnlockParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *UserUpdateUnlockParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *UserUpdateUnlockParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the user update unlock params
func (o *UserUpdateUnlockParams) WithXDomainrobotWS(xDomainrobotWS *string) *UserUpdateUnlockParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the user update unlock params
func (o *UserUpdateUnlockParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithContext adds the context to the user update unlock params
func (o *UserUpdateUnlockParams) WithContext(context string) *UserUpdateUnlockParams {
	o.SetContext(context)
	return o
}

// SetContext adds the context to the user update unlock params
func (o *UserUpdateUnlockParams) SetContext(context string) {
	o.Context = context
}

// WithKeys adds the keys to the user update unlock params
func (o *UserUpdateUnlockParams) WithKeys(keys []string) *UserUpdateUnlockParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the user update unlock params
func (o *UserUpdateUnlockParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WithName adds the name to the user update unlock params
func (o *UserUpdateUnlockParams) WithName(name string) *UserUpdateUnlockParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the user update unlock params
func (o *UserUpdateUnlockParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *UserUpdateUnlockParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param context
	if err := r.SetPathParam("context", o.Context); err != nil {
		return err
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

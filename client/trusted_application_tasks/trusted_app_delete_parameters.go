// Code generated by go-swagger; DO NOT EDIT.

package trusted_application_tasks

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

// NewTrustedAppDeleteParams creates a new TrustedAppDeleteParams object
// with the default values initialized.
func NewTrustedAppDeleteParams() *TrustedAppDeleteParams {
	var ()
	return &TrustedAppDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTrustedAppDeleteParamsWithTimeout creates a new TrustedAppDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrustedAppDeleteParamsWithTimeout(timeout time.Duration) *TrustedAppDeleteParams {
	var ()
	return &TrustedAppDeleteParams{

		timeout: timeout,
	}
}

// NewTrustedAppDeleteParamsWithContext creates a new TrustedAppDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrustedAppDeleteParamsWithContext(ctx context.Context) *TrustedAppDeleteParams {
	var ()
	return &TrustedAppDeleteParams{

		Context: ctx,
	}
}

// NewTrustedAppDeleteParamsWithHTTPClient creates a new TrustedAppDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrustedAppDeleteParamsWithHTTPClient(client *http.Client) *TrustedAppDeleteParams {
	var ()
	return &TrustedAppDeleteParams{
		HTTPClient: client,
	}
}

/*TrustedAppDeleteParams contains all the parameters to send to the API endpoint
for the trusted app delete operation typically these are written to a http.Request
*/
type TrustedAppDeleteParams struct {

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
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trusted app delete params
func (o *TrustedAppDeleteParams) WithTimeout(timeout time.Duration) *TrustedAppDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trusted app delete params
func (o *TrustedAppDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trusted app delete params
func (o *TrustedAppDeleteParams) WithContext(ctx context.Context) *TrustedAppDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trusted app delete params
func (o *TrustedAppDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trusted app delete params
func (o *TrustedAppDeleteParams) WithHTTPClient(client *http.Client) *TrustedAppDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trusted app delete params
func (o *TrustedAppDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *TrustedAppDeleteParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotContext(xDomainrobotContext *int32) *TrustedAppDeleteParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *TrustedAppDeleteParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *TrustedAppDeleteParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *TrustedAppDeleteParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the trusted app delete params
func (o *TrustedAppDeleteParams) WithXDomainrobotWS(xDomainrobotWS *string) *TrustedAppDeleteParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the trusted app delete params
func (o *TrustedAppDeleteParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithID adds the id to the trusted app delete params
func (o *TrustedAppDeleteParams) WithID(id int32) *TrustedAppDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the trusted app delete params
func (o *TrustedAppDeleteParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *TrustedAppDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
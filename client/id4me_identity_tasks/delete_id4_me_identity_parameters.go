// Code generated by go-swagger; DO NOT EDIT.

package id4me_identity_tasks

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

// NewDeleteId4MeIdentityParams creates a new DeleteId4MeIdentityParams object
// with the default values initialized.
func NewDeleteId4MeIdentityParams() *DeleteId4MeIdentityParams {
	var ()
	return &DeleteId4MeIdentityParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteId4MeIdentityParamsWithTimeout creates a new DeleteId4MeIdentityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteId4MeIdentityParamsWithTimeout(timeout time.Duration) *DeleteId4MeIdentityParams {
	var ()
	return &DeleteId4MeIdentityParams{

		timeout: timeout,
	}
}

// NewDeleteId4MeIdentityParamsWithContext creates a new DeleteId4MeIdentityParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteId4MeIdentityParamsWithContext(ctx context.Context) *DeleteId4MeIdentityParams {
	var ()
	return &DeleteId4MeIdentityParams{

		Context: ctx,
	}
}

// NewDeleteId4MeIdentityParamsWithHTTPClient creates a new DeleteId4MeIdentityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteId4MeIdentityParamsWithHTTPClient(client *http.Client) *DeleteId4MeIdentityParams {
	var ()
	return &DeleteId4MeIdentityParams{
		HTTPClient: client,
	}
}

/*DeleteId4MeIdentityParams contains all the parameters to send to the API endpoint
for the delete id4 me identity operation typically these are written to a http.Request
*/
type DeleteId4MeIdentityParams struct {

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
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithTimeout(timeout time.Duration) *DeleteId4MeIdentityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithContext(ctx context.Context) *DeleteId4MeIdentityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithHTTPClient(client *http.Client) *DeleteId4MeIdentityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithXDomainrobotWS(xDomainrobotWS *string) *DeleteId4MeIdentityParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) WithName(name string) *DeleteId4MeIdentityParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete id4 me identity params
func (o *DeleteId4MeIdentityParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteId4MeIdentityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

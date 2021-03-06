// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

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

// NewDomainDeleteDomainSafeParams creates a new DomainDeleteDomainSafeParams object
// with the default values initialized.
func NewDomainDeleteDomainSafeParams() *DomainDeleteDomainSafeParams {
	var ()
	return &DomainDeleteDomainSafeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainDeleteDomainSafeParamsWithTimeout creates a new DomainDeleteDomainSafeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainDeleteDomainSafeParamsWithTimeout(timeout time.Duration) *DomainDeleteDomainSafeParams {
	var ()
	return &DomainDeleteDomainSafeParams{

		timeout: timeout,
	}
}

// NewDomainDeleteDomainSafeParamsWithContext creates a new DomainDeleteDomainSafeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainDeleteDomainSafeParamsWithContext(ctx context.Context) *DomainDeleteDomainSafeParams {
	var ()
	return &DomainDeleteDomainSafeParams{

		Context: ctx,
	}
}

// NewDomainDeleteDomainSafeParamsWithHTTPClient creates a new DomainDeleteDomainSafeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainDeleteDomainSafeParamsWithHTTPClient(client *http.Client) *DomainDeleteDomainSafeParams {
	var ()
	return &DomainDeleteDomainSafeParams{
		HTTPClient: client,
	}
}

/*DomainDeleteDomainSafeParams contains all the parameters to send to the API endpoint
for the domain delete domain safe operation typically these are written to a http.Request
*/
type DomainDeleteDomainSafeParams struct {

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
	/*SystemNameServer*/
	SystemNameServer string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithTimeout(timeout time.Duration) *DomainDeleteDomainSafeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithContext(ctx context.Context) *DomainDeleteDomainSafeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithHTTPClient(client *http.Client) *DomainDeleteDomainSafeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainDeleteDomainSafeParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithName(name string) *DomainDeleteDomainSafeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetName(name string) {
	o.Name = name
}

// WithSystemNameServer adds the systemNameServer to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) WithSystemNameServer(systemNameServer string) *DomainDeleteDomainSafeParams {
	o.SetSystemNameServer(systemNameServer)
	return o
}

// SetSystemNameServer adds the systemNameServer to the domain delete domain safe params
func (o *DomainDeleteDomainSafeParams) SetSystemNameServer(systemNameServer string) {
	o.SystemNameServer = systemNameServer
}

// WriteToRequest writes these params to a swagger request
func (o *DomainDeleteDomainSafeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param systemNameServer
	if err := r.SetPathParam("systemNameServer", o.SystemNameServer); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

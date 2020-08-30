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

// NewDomainAddDomainSafeParams creates a new DomainAddDomainSafeParams object
// with the default values initialized.
func NewDomainAddDomainSafeParams() *DomainAddDomainSafeParams {
	var ()
	return &DomainAddDomainSafeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainAddDomainSafeParamsWithTimeout creates a new DomainAddDomainSafeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainAddDomainSafeParamsWithTimeout(timeout time.Duration) *DomainAddDomainSafeParams {
	var ()
	return &DomainAddDomainSafeParams{

		timeout: timeout,
	}
}

// NewDomainAddDomainSafeParamsWithContext creates a new DomainAddDomainSafeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainAddDomainSafeParamsWithContext(ctx context.Context) *DomainAddDomainSafeParams {
	var ()
	return &DomainAddDomainSafeParams{

		Context: ctx,
	}
}

// NewDomainAddDomainSafeParamsWithHTTPClient creates a new DomainAddDomainSafeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainAddDomainSafeParamsWithHTTPClient(client *http.Client) *DomainAddDomainSafeParams {
	var ()
	return &DomainAddDomainSafeParams{
		HTTPClient: client,
	}
}

/*DomainAddDomainSafeParams contains all the parameters to send to the API endpoint
for the domain add domain safe operation typically these are written to a http.Request
*/
type DomainAddDomainSafeParams struct {

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

// WithTimeout adds the timeout to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithTimeout(timeout time.Duration) *DomainAddDomainSafeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithContext(ctx context.Context) *DomainAddDomainSafeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithHTTPClient(client *http.Client) *DomainAddDomainSafeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DomainAddDomainSafeParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DomainAddDomainSafeParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DomainAddDomainSafeParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DomainAddDomainSafeParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DomainAddDomainSafeParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainAddDomainSafeParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithName(name string) *DomainAddDomainSafeParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetName(name string) {
	o.Name = name
}

// WithSystemNameServer adds the systemNameServer to the domain add domain safe params
func (o *DomainAddDomainSafeParams) WithSystemNameServer(systemNameServer string) *DomainAddDomainSafeParams {
	o.SetSystemNameServer(systemNameServer)
	return o
}

// SetSystemNameServer adds the systemNameServer to the domain add domain safe params
func (o *DomainAddDomainSafeParams) SetSystemNameServer(systemNameServer string) {
	o.SystemNameServer = systemNameServer
}

// WriteToRequest writes these params to a swagger request
func (o *DomainAddDomainSafeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
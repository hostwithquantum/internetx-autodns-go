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

// NewZoneAxfrParams creates a new ZoneAxfrParams object
// with the default values initialized.
func NewZoneAxfrParams() *ZoneAxfrParams {
	var ()
	return &ZoneAxfrParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewZoneAxfrParamsWithTimeout creates a new ZoneAxfrParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewZoneAxfrParamsWithTimeout(timeout time.Duration) *ZoneAxfrParams {
	var ()
	return &ZoneAxfrParams{

		timeout: timeout,
	}
}

// NewZoneAxfrParamsWithContext creates a new ZoneAxfrParams object
// with the default values initialized, and the ability to set a context for a request
func NewZoneAxfrParamsWithContext(ctx context.Context) *ZoneAxfrParams {
	var ()
	return &ZoneAxfrParams{

		Context: ctx,
	}
}

// NewZoneAxfrParamsWithHTTPClient creates a new ZoneAxfrParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewZoneAxfrParamsWithHTTPClient(client *http.Client) *ZoneAxfrParams {
	var ()
	return &ZoneAxfrParams{
		HTTPClient: client,
	}
}

/*ZoneAxfrParams contains all the parameters to send to the API endpoint
for the zone axfr operation typically these are written to a http.Request
*/
type ZoneAxfrParams struct {

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

// WithTimeout adds the timeout to the zone axfr params
func (o *ZoneAxfrParams) WithTimeout(timeout time.Duration) *ZoneAxfrParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zone axfr params
func (o *ZoneAxfrParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zone axfr params
func (o *ZoneAxfrParams) WithContext(ctx context.Context) *ZoneAxfrParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zone axfr params
func (o *ZoneAxfrParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zone axfr params
func (o *ZoneAxfrParams) WithHTTPClient(client *http.Client) *ZoneAxfrParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zone axfr params
func (o *ZoneAxfrParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *ZoneAxfrParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotContext(xDomainrobotContext *int32) *ZoneAxfrParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *ZoneAxfrParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *ZoneAxfrParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *ZoneAxfrParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *ZoneAxfrParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *ZoneAxfrParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *ZoneAxfrParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ZoneAxfrParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ZoneAxfrParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the zone axfr params
func (o *ZoneAxfrParams) WithXDomainrobotWS(xDomainrobotWS *string) *ZoneAxfrParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the zone axfr params
func (o *ZoneAxfrParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithName adds the name to the zone axfr params
func (o *ZoneAxfrParams) WithName(name string) *ZoneAxfrParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the zone axfr params
func (o *ZoneAxfrParams) SetName(name string) {
	o.Name = name
}

// WithSystemNameServer adds the systemNameServer to the zone axfr params
func (o *ZoneAxfrParams) WithSystemNameServer(systemNameServer string) *ZoneAxfrParams {
	o.SetSystemNameServer(systemNameServer)
	return o
}

// SetSystemNameServer adds the systemNameServer to the zone axfr params
func (o *ZoneAxfrParams) SetSystemNameServer(systemNameServer string) {
	o.SystemNameServer = systemNameServer
}

// WriteToRequest writes these params to a swagger request
func (o *ZoneAxfrParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
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

// NewZoneHistoryInfoParams creates a new ZoneHistoryInfoParams object
// with the default values initialized.
func NewZoneHistoryInfoParams() *ZoneHistoryInfoParams {
	var ()
	return &ZoneHistoryInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewZoneHistoryInfoParamsWithTimeout creates a new ZoneHistoryInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewZoneHistoryInfoParamsWithTimeout(timeout time.Duration) *ZoneHistoryInfoParams {
	var ()
	return &ZoneHistoryInfoParams{

		timeout: timeout,
	}
}

// NewZoneHistoryInfoParamsWithContext creates a new ZoneHistoryInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewZoneHistoryInfoParamsWithContext(ctx context.Context) *ZoneHistoryInfoParams {
	var ()
	return &ZoneHistoryInfoParams{

		Context: ctx,
	}
}

// NewZoneHistoryInfoParamsWithHTTPClient creates a new ZoneHistoryInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewZoneHistoryInfoParamsWithHTTPClient(client *http.Client) *ZoneHistoryInfoParams {
	var ()
	return &ZoneHistoryInfoParams{
		HTTPClient: client,
	}
}

/*ZoneHistoryInfoParams contains all the parameters to send to the API endpoint
for the zone history info operation typically these are written to a http.Request
*/
type ZoneHistoryInfoParams struct {

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
	/*LogID*/
	LogID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the zone history info params
func (o *ZoneHistoryInfoParams) WithTimeout(timeout time.Duration) *ZoneHistoryInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the zone history info params
func (o *ZoneHistoryInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the zone history info params
func (o *ZoneHistoryInfoParams) WithContext(ctx context.Context) *ZoneHistoryInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the zone history info params
func (o *ZoneHistoryInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the zone history info params
func (o *ZoneHistoryInfoParams) WithHTTPClient(client *http.Client) *ZoneHistoryInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the zone history info params
func (o *ZoneHistoryInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *ZoneHistoryInfoParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotContext(xDomainrobotContext *int32) *ZoneHistoryInfoParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *ZoneHistoryInfoParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *ZoneHistoryInfoParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *ZoneHistoryInfoParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the zone history info params
func (o *ZoneHistoryInfoParams) WithXDomainrobotWS(xDomainrobotWS *string) *ZoneHistoryInfoParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the zone history info params
func (o *ZoneHistoryInfoParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithLogID adds the logID to the zone history info params
func (o *ZoneHistoryInfoParams) WithLogID(logID int32) *ZoneHistoryInfoParams {
	o.SetLogID(logID)
	return o
}

// SetLogID adds the logId to the zone history info params
func (o *ZoneHistoryInfoParams) SetLogID(logID int32) {
	o.LogID = logID
}

// WriteToRequest writes these params to a swagger request
func (o *ZoneHistoryInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param logId
	if err := r.SetPathParam("logId", swag.FormatInt32(o.LogID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

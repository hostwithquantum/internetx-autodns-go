// Code generated by go-swagger; DO NOT EDIT.

package subscription_tasks

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

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// NewSubscriptionCreateParams creates a new SubscriptionCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSubscriptionCreateParams() *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSubscriptionCreateParamsWithTimeout creates a new SubscriptionCreateParams object
// with the ability to set a timeout on a request.
func NewSubscriptionCreateParamsWithTimeout(timeout time.Duration) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		timeout: timeout,
	}
}

// NewSubscriptionCreateParamsWithContext creates a new SubscriptionCreateParams object
// with the ability to set a context for a request.
func NewSubscriptionCreateParamsWithContext(ctx context.Context) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		Context: ctx,
	}
}

// NewSubscriptionCreateParamsWithHTTPClient creates a new SubscriptionCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSubscriptionCreateParamsWithHTTPClient(client *http.Client) *SubscriptionCreateParams {
	return &SubscriptionCreateParams{
		HTTPClient: client,
	}
}

/*
SubscriptionCreateParams contains all the parameters to send to the API endpoint

	for the subscription create operation.

	Typically these are written to a http.Request.
*/
type SubscriptionCreateParams struct {

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

	// Subscription.
	Subscription *models.PeriodicBilling

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the subscription create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionCreateParams) WithDefaults() *SubscriptionCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the subscription create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SubscriptionCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the subscription create params
func (o *SubscriptionCreateParams) WithTimeout(timeout time.Duration) *SubscriptionCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscription create params
func (o *SubscriptionCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscription create params
func (o *SubscriptionCreateParams) WithContext(ctx context.Context) *SubscriptionCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscription create params
func (o *SubscriptionCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscription create params
func (o *SubscriptionCreateParams) WithHTTPClient(client *http.Client) *SubscriptionCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscription create params
func (o *SubscriptionCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *SubscriptionCreateParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *SubscriptionCreateParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotContext(xDomainrobotContext *int32) *SubscriptionCreateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *SubscriptionCreateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *SubscriptionCreateParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *SubscriptionCreateParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *SubscriptionCreateParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *SubscriptionCreateParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *SubscriptionCreateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *SubscriptionCreateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *SubscriptionCreateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the subscription create params
func (o *SubscriptionCreateParams) WithXDomainrobotWS(xDomainrobotWS *string) *SubscriptionCreateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the subscription create params
func (o *SubscriptionCreateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithSubscription adds the subscription to the subscription create params
func (o *SubscriptionCreateParams) WithSubscription(subscription *models.PeriodicBilling) *SubscriptionCreateParams {
	o.SetSubscription(subscription)
	return o
}

// SetSubscription adds the subscription to the subscription create params
func (o *SubscriptionCreateParams) SetSubscription(subscription *models.PeriodicBilling) {
	o.Subscription = subscription
}

// WriteToRequest writes these params to a swagger request
func (o *SubscriptionCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Subscription != nil {
		if err := r.SetBodyParam(o.Subscription); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

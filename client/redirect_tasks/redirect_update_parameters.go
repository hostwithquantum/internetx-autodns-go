// Code generated by go-swagger; DO NOT EDIT.

package redirect_tasks

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

// NewRedirectUpdateParams creates a new RedirectUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRedirectUpdateParams() *RedirectUpdateParams {
	return &RedirectUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRedirectUpdateParamsWithTimeout creates a new RedirectUpdateParams object
// with the ability to set a timeout on a request.
func NewRedirectUpdateParamsWithTimeout(timeout time.Duration) *RedirectUpdateParams {
	return &RedirectUpdateParams{
		timeout: timeout,
	}
}

// NewRedirectUpdateParamsWithContext creates a new RedirectUpdateParams object
// with the ability to set a context for a request.
func NewRedirectUpdateParamsWithContext(ctx context.Context) *RedirectUpdateParams {
	return &RedirectUpdateParams{
		Context: ctx,
	}
}

// NewRedirectUpdateParamsWithHTTPClient creates a new RedirectUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewRedirectUpdateParamsWithHTTPClient(client *http.Client) *RedirectUpdateParams {
	return &RedirectUpdateParams{
		HTTPClient: client,
	}
}

/*
RedirectUpdateParams contains all the parameters to send to the API endpoint

	for the redirect update operation.

	Typically these are written to a http.Request.
*/
type RedirectUpdateParams struct {

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

	/* Body.

	   redirect
	*/
	Body *models.Redirect

	// Source.
	Source string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the redirect update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedirectUpdateParams) WithDefaults() *RedirectUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the redirect update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RedirectUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the redirect update params
func (o *RedirectUpdateParams) WithTimeout(timeout time.Duration) *RedirectUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the redirect update params
func (o *RedirectUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the redirect update params
func (o *RedirectUpdateParams) WithContext(ctx context.Context) *RedirectUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the redirect update params
func (o *RedirectUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the redirect update params
func (o *RedirectUpdateParams) WithHTTPClient(client *http.Client) *RedirectUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the redirect update params
func (o *RedirectUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *RedirectUpdateParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *RedirectUpdateParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotContext(xDomainrobotContext *int32) *RedirectUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *RedirectUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *RedirectUpdateParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *RedirectUpdateParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *RedirectUpdateParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *RedirectUpdateParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *RedirectUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *RedirectUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *RedirectUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the redirect update params
func (o *RedirectUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *RedirectUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the redirect update params
func (o *RedirectUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the redirect update params
func (o *RedirectUpdateParams) WithBody(body *models.Redirect) *RedirectUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the redirect update params
func (o *RedirectUpdateParams) SetBody(body *models.Redirect) {
	o.Body = body
}

// WithSource adds the source to the redirect update params
func (o *RedirectUpdateParams) WithSource(source string) *RedirectUpdateParams {
	o.SetSource(source)
	return o
}

// SetSource adds the source to the redirect update params
func (o *RedirectUpdateParams) SetSource(source string) {
	o.Source = source
}

// WriteToRequest writes these params to a swagger request
func (o *RedirectUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

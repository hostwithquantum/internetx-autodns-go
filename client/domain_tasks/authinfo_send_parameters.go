// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

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

// NewAuthinfoSendParams creates a new AuthinfoSendParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAuthinfoSendParams() *AuthinfoSendParams {
	return &AuthinfoSendParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAuthinfoSendParamsWithTimeout creates a new AuthinfoSendParams object
// with the ability to set a timeout on a request.
func NewAuthinfoSendParamsWithTimeout(timeout time.Duration) *AuthinfoSendParams {
	return &AuthinfoSendParams{
		timeout: timeout,
	}
}

// NewAuthinfoSendParamsWithContext creates a new AuthinfoSendParams object
// with the ability to set a context for a request.
func NewAuthinfoSendParamsWithContext(ctx context.Context) *AuthinfoSendParams {
	return &AuthinfoSendParams{
		Context: ctx,
	}
}

// NewAuthinfoSendParamsWithHTTPClient creates a new AuthinfoSendParams object
// with the ability to set a custom HTTPClient for a request.
func NewAuthinfoSendParamsWithHTTPClient(client *http.Client) *AuthinfoSendParams {
	return &AuthinfoSendParams{
		HTTPClient: client,
	}
}

/*
AuthinfoSendParams contains all the parameters to send to the API endpoint

	for the authinfo send operation.

	Typically these are written to a http.Request.
*/
type AuthinfoSendParams struct {

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

	   domain
	*/
	Body *models.Domain

	// Name.
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the authinfo send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthinfoSendParams) WithDefaults() *AuthinfoSendParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the authinfo send params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AuthinfoSendParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the authinfo send params
func (o *AuthinfoSendParams) WithTimeout(timeout time.Duration) *AuthinfoSendParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the authinfo send params
func (o *AuthinfoSendParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the authinfo send params
func (o *AuthinfoSendParams) WithContext(ctx context.Context) *AuthinfoSendParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the authinfo send params
func (o *AuthinfoSendParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the authinfo send params
func (o *AuthinfoSendParams) WithHTTPClient(client *http.Client) *AuthinfoSendParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the authinfo send params
func (o *AuthinfoSendParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *AuthinfoSendParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *AuthinfoSendParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotContext(xDomainrobotContext *int32) *AuthinfoSendParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *AuthinfoSendParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *AuthinfoSendParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *AuthinfoSendParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *AuthinfoSendParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *AuthinfoSendParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *AuthinfoSendParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *AuthinfoSendParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *AuthinfoSendParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the authinfo send params
func (o *AuthinfoSendParams) WithXDomainrobotWS(xDomainrobotWS *string) *AuthinfoSendParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the authinfo send params
func (o *AuthinfoSendParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the authinfo send params
func (o *AuthinfoSendParams) WithBody(body *models.Domain) *AuthinfoSendParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the authinfo send params
func (o *AuthinfoSendParams) SetBody(body *models.Domain) {
	o.Body = body
}

// WithName adds the name to the authinfo send params
func (o *AuthinfoSendParams) WithName(name string) *AuthinfoSendParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the authinfo send params
func (o *AuthinfoSendParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *AuthinfoSendParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

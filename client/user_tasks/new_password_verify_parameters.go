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

// NewNewPasswordVerifyParams creates a new NewPasswordVerifyParams object
// with the default values initialized.
func NewNewPasswordVerifyParams() *NewPasswordVerifyParams {
	var ()
	return &NewPasswordVerifyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewNewPasswordVerifyParamsWithTimeout creates a new NewPasswordVerifyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewNewPasswordVerifyParamsWithTimeout(timeout time.Duration) *NewPasswordVerifyParams {
	var ()
	return &NewPasswordVerifyParams{

		timeout: timeout,
	}
}

// NewNewPasswordVerifyParamsWithContext creates a new NewPasswordVerifyParams object
// with the default values initialized, and the ability to set a context for a request
func NewNewPasswordVerifyParamsWithContext(ctx context.Context) *NewPasswordVerifyParams {
	var ()
	return &NewPasswordVerifyParams{

		Context: ctx,
	}
}

// NewNewPasswordVerifyParamsWithHTTPClient creates a new NewPasswordVerifyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewNewPasswordVerifyParamsWithHTTPClient(client *http.Client) *NewPasswordVerifyParams {
	var ()
	return &NewPasswordVerifyParams{
		HTTPClient: client,
	}
}

/*NewPasswordVerifyParams contains all the parameters to send to the API endpoint
for the new password verify operation typically these are written to a http.Request
*/
type NewPasswordVerifyParams struct {

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
	/*Token*/
	Token string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the new password verify params
func (o *NewPasswordVerifyParams) WithTimeout(timeout time.Duration) *NewPasswordVerifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the new password verify params
func (o *NewPasswordVerifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the new password verify params
func (o *NewPasswordVerifyParams) WithContext(ctx context.Context) *NewPasswordVerifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the new password verify params
func (o *NewPasswordVerifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the new password verify params
func (o *NewPasswordVerifyParams) WithHTTPClient(client *http.Client) *NewPasswordVerifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the new password verify params
func (o *NewPasswordVerifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *NewPasswordVerifyParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotContext(xDomainrobotContext *int32) *NewPasswordVerifyParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *NewPasswordVerifyParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *NewPasswordVerifyParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *NewPasswordVerifyParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the new password verify params
func (o *NewPasswordVerifyParams) WithXDomainrobotWS(xDomainrobotWS *string) *NewPasswordVerifyParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the new password verify params
func (o *NewPasswordVerifyParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithToken adds the token to the new password verify params
func (o *NewPasswordVerifyParams) WithToken(token string) *NewPasswordVerifyParams {
	o.SetToken(token)
	return o
}

// SetToken adds the token to the new password verify params
func (o *NewPasswordVerifyParams) SetToken(token string) {
	o.Token = token
}

// WriteToRequest writes these params to a swagger request
func (o *NewPasswordVerifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param token
	if err := r.SetPathParam("token", o.Token); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

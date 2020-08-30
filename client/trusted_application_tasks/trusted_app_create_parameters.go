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

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// NewTrustedAppCreateParams creates a new TrustedAppCreateParams object
// with the default values initialized.
func NewTrustedAppCreateParams() *TrustedAppCreateParams {
	var ()
	return &TrustedAppCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTrustedAppCreateParamsWithTimeout creates a new TrustedAppCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrustedAppCreateParamsWithTimeout(timeout time.Duration) *TrustedAppCreateParams {
	var ()
	return &TrustedAppCreateParams{

		timeout: timeout,
	}
}

// NewTrustedAppCreateParamsWithContext creates a new TrustedAppCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrustedAppCreateParamsWithContext(ctx context.Context) *TrustedAppCreateParams {
	var ()
	return &TrustedAppCreateParams{

		Context: ctx,
	}
}

// NewTrustedAppCreateParamsWithHTTPClient creates a new TrustedAppCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrustedAppCreateParamsWithHTTPClient(client *http.Client) *TrustedAppCreateParams {
	var ()
	return &TrustedAppCreateParams{
		HTTPClient: client,
	}
}

/*TrustedAppCreateParams contains all the parameters to send to the API endpoint
for the trusted app create operation typically these are written to a http.Request
*/
type TrustedAppCreateParams struct {

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
	/*Body
	  trustedApplication

	*/
	Body *models.TrustedApplication

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trusted app create params
func (o *TrustedAppCreateParams) WithTimeout(timeout time.Duration) *TrustedAppCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trusted app create params
func (o *TrustedAppCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trusted app create params
func (o *TrustedAppCreateParams) WithContext(ctx context.Context) *TrustedAppCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trusted app create params
func (o *TrustedAppCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trusted app create params
func (o *TrustedAppCreateParams) WithHTTPClient(client *http.Client) *TrustedAppCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trusted app create params
func (o *TrustedAppCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *TrustedAppCreateParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotContext(xDomainrobotContext *int32) *TrustedAppCreateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *TrustedAppCreateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *TrustedAppCreateParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *TrustedAppCreateParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *TrustedAppCreateParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *TrustedAppCreateParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *TrustedAppCreateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *TrustedAppCreateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *TrustedAppCreateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the trusted app create params
func (o *TrustedAppCreateParams) WithXDomainrobotWS(xDomainrobotWS *string) *TrustedAppCreateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the trusted app create params
func (o *TrustedAppCreateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the trusted app create params
func (o *TrustedAppCreateParams) WithBody(body *models.TrustedApplication) *TrustedAppCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trusted app create params
func (o *TrustedAppCreateParams) SetBody(body *models.TrustedApplication) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TrustedAppCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

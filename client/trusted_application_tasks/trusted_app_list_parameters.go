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

// NewTrustedAppListParams creates a new TrustedAppListParams object
// with the default values initialized.
func NewTrustedAppListParams() *TrustedAppListParams {
	var ()
	return &TrustedAppListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTrustedAppListParamsWithTimeout creates a new TrustedAppListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTrustedAppListParamsWithTimeout(timeout time.Duration) *TrustedAppListParams {
	var ()
	return &TrustedAppListParams{

		timeout: timeout,
	}
}

// NewTrustedAppListParamsWithContext creates a new TrustedAppListParams object
// with the default values initialized, and the ability to set a context for a request
func NewTrustedAppListParamsWithContext(ctx context.Context) *TrustedAppListParams {
	var ()
	return &TrustedAppListParams{

		Context: ctx,
	}
}

// NewTrustedAppListParamsWithHTTPClient creates a new TrustedAppListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTrustedAppListParamsWithHTTPClient(client *http.Client) *TrustedAppListParams {
	var ()
	return &TrustedAppListParams{
		HTTPClient: client,
	}
}

/*TrustedAppListParams contains all the parameters to send to the API endpoint
for the trusted app list operation typically these are written to a http.Request
*/
type TrustedAppListParams struct {

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
	  query

	*/
	Body *models.Query
	/*Keys
	  The query parameter to inquire additional details.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the trusted app list params
func (o *TrustedAppListParams) WithTimeout(timeout time.Duration) *TrustedAppListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the trusted app list params
func (o *TrustedAppListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the trusted app list params
func (o *TrustedAppListParams) WithContext(ctx context.Context) *TrustedAppListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the trusted app list params
func (o *TrustedAppListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the trusted app list params
func (o *TrustedAppListParams) WithHTTPClient(client *http.Client) *TrustedAppListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the trusted app list params
func (o *TrustedAppListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *TrustedAppListParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotContext(xDomainrobotContext *int32) *TrustedAppListParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *TrustedAppListParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *TrustedAppListParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *TrustedAppListParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *TrustedAppListParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *TrustedAppListParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *TrustedAppListParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *TrustedAppListParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *TrustedAppListParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the trusted app list params
func (o *TrustedAppListParams) WithXDomainrobotWS(xDomainrobotWS *string) *TrustedAppListParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the trusted app list params
func (o *TrustedAppListParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the trusted app list params
func (o *TrustedAppListParams) WithBody(body *models.Query) *TrustedAppListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the trusted app list params
func (o *TrustedAppListParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithKeys adds the keys to the trusted app list params
func (o *TrustedAppListParams) WithKeys(keys []string) *TrustedAppListParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the trusted app list params
func (o *TrustedAppListParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *TrustedAppListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
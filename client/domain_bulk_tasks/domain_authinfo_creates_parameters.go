// Code generated by go-swagger; DO NOT EDIT.

package domain_bulk_tasks

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

// NewDomainAuthinfoCreatesParams creates a new DomainAuthinfoCreatesParams object
// with the default values initialized.
func NewDomainAuthinfoCreatesParams() *DomainAuthinfoCreatesParams {
	var ()
	return &DomainAuthinfoCreatesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainAuthinfoCreatesParamsWithTimeout creates a new DomainAuthinfoCreatesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainAuthinfoCreatesParamsWithTimeout(timeout time.Duration) *DomainAuthinfoCreatesParams {
	var ()
	return &DomainAuthinfoCreatesParams{

		timeout: timeout,
	}
}

// NewDomainAuthinfoCreatesParamsWithContext creates a new DomainAuthinfoCreatesParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainAuthinfoCreatesParamsWithContext(ctx context.Context) *DomainAuthinfoCreatesParams {
	var ()
	return &DomainAuthinfoCreatesParams{

		Context: ctx,
	}
}

// NewDomainAuthinfoCreatesParamsWithHTTPClient creates a new DomainAuthinfoCreatesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainAuthinfoCreatesParamsWithHTTPClient(client *http.Client) *DomainAuthinfoCreatesParams {
	var ()
	return &DomainAuthinfoCreatesParams{
		HTTPClient: client,
	}
}

/*DomainAuthinfoCreatesParams contains all the parameters to send to the API endpoint
for the domain authinfo creates operation typically these are written to a http.Request
*/
type DomainAuthinfoCreatesParams struct {

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
	  domains

	*/
	Body *models.BulkDomainDeleteRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithTimeout(timeout time.Duration) *DomainAuthinfoCreatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithContext(ctx context.Context) *DomainAuthinfoCreatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithHTTPClient(client *http.Client) *DomainAuthinfoCreatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainAuthinfoCreatesParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) WithBody(body *models.BulkDomainDeleteRequest) *DomainAuthinfoCreatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain authinfo creates params
func (o *DomainAuthinfoCreatesParams) SetBody(body *models.BulkDomainDeleteRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *DomainAuthinfoCreatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

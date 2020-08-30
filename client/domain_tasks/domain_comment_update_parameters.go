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

// NewDomainCommentUpdateParams creates a new DomainCommentUpdateParams object
// with the default values initialized.
func NewDomainCommentUpdateParams() *DomainCommentUpdateParams {
	var ()
	return &DomainCommentUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDomainCommentUpdateParamsWithTimeout creates a new DomainCommentUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDomainCommentUpdateParamsWithTimeout(timeout time.Duration) *DomainCommentUpdateParams {
	var ()
	return &DomainCommentUpdateParams{

		timeout: timeout,
	}
}

// NewDomainCommentUpdateParamsWithContext creates a new DomainCommentUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDomainCommentUpdateParamsWithContext(ctx context.Context) *DomainCommentUpdateParams {
	var ()
	return &DomainCommentUpdateParams{

		Context: ctx,
	}
}

// NewDomainCommentUpdateParamsWithHTTPClient creates a new DomainCommentUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDomainCommentUpdateParamsWithHTTPClient(client *http.Client) *DomainCommentUpdateParams {
	var ()
	return &DomainCommentUpdateParams{
		HTTPClient: client,
	}
}

/*DomainCommentUpdateParams contains all the parameters to send to the API endpoint
for the domain comment update operation typically these are written to a http.Request
*/
type DomainCommentUpdateParams struct {

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
	  domain

	*/
	Body *models.Domain
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the domain comment update params
func (o *DomainCommentUpdateParams) WithTimeout(timeout time.Duration) *DomainCommentUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain comment update params
func (o *DomainCommentUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain comment update params
func (o *DomainCommentUpdateParams) WithContext(ctx context.Context) *DomainCommentUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain comment update params
func (o *DomainCommentUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain comment update params
func (o *DomainCommentUpdateParams) WithHTTPClient(client *http.Client) *DomainCommentUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain comment update params
func (o *DomainCommentUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DomainCommentUpdateParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DomainCommentUpdateParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DomainCommentUpdateParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DomainCommentUpdateParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DomainCommentUpdateParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the domain comment update params
func (o *DomainCommentUpdateParams) WithXDomainrobotWS(xDomainrobotWS *string) *DomainCommentUpdateParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the domain comment update params
func (o *DomainCommentUpdateParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the domain comment update params
func (o *DomainCommentUpdateParams) WithBody(body *models.Domain) *DomainCommentUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain comment update params
func (o *DomainCommentUpdateParams) SetBody(body *models.Domain) {
	o.Body = body
}

// WithName adds the name to the domain comment update params
func (o *DomainCommentUpdateParams) WithName(name string) *DomainCommentUpdateParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the domain comment update params
func (o *DomainCommentUpdateParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *DomainCommentUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package object_user_assignment_tasks

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

// NewObjectUserAssignmentParams creates a new ObjectUserAssignmentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewObjectUserAssignmentParams() *ObjectUserAssignmentParams {
	return &ObjectUserAssignmentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewObjectUserAssignmentParamsWithTimeout creates a new ObjectUserAssignmentParams object
// with the ability to set a timeout on a request.
func NewObjectUserAssignmentParamsWithTimeout(timeout time.Duration) *ObjectUserAssignmentParams {
	return &ObjectUserAssignmentParams{
		timeout: timeout,
	}
}

// NewObjectUserAssignmentParamsWithContext creates a new ObjectUserAssignmentParams object
// with the ability to set a context for a request.
func NewObjectUserAssignmentParamsWithContext(ctx context.Context) *ObjectUserAssignmentParams {
	return &ObjectUserAssignmentParams{
		Context: ctx,
	}
}

// NewObjectUserAssignmentParamsWithHTTPClient creates a new ObjectUserAssignmentParams object
// with the ability to set a custom HTTPClient for a request.
func NewObjectUserAssignmentParamsWithHTTPClient(client *http.Client) *ObjectUserAssignmentParams {
	return &ObjectUserAssignmentParams{
		HTTPClient: client,
	}
}

/*
ObjectUserAssignmentParams contains all the parameters to send to the API endpoint

	for the object user assignment operation.

	Typically these are written to a http.Request.
*/
type ObjectUserAssignmentParams struct {

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

	   object
	*/
	Body *models.ObjectUserAssignment

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the object user assignment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectUserAssignmentParams) WithDefaults() *ObjectUserAssignmentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the object user assignment params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ObjectUserAssignmentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the object user assignment params
func (o *ObjectUserAssignmentParams) WithTimeout(timeout time.Duration) *ObjectUserAssignmentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the object user assignment params
func (o *ObjectUserAssignmentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the object user assignment params
func (o *ObjectUserAssignmentParams) WithContext(ctx context.Context) *ObjectUserAssignmentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the object user assignment params
func (o *ObjectUserAssignmentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the object user assignment params
func (o *ObjectUserAssignmentParams) WithHTTPClient(client *http.Client) *ObjectUserAssignmentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the object user assignment params
func (o *ObjectUserAssignmentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *ObjectUserAssignmentParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *ObjectUserAssignmentParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotContext(xDomainrobotContext *int32) *ObjectUserAssignmentParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *ObjectUserAssignmentParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *ObjectUserAssignmentParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *ObjectUserAssignmentParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the object user assignment params
func (o *ObjectUserAssignmentParams) WithXDomainrobotWS(xDomainrobotWS *string) *ObjectUserAssignmentParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the object user assignment params
func (o *ObjectUserAssignmentParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the object user assignment params
func (o *ObjectUserAssignmentParams) WithBody(body *models.ObjectUserAssignment) *ObjectUserAssignmentParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the object user assignment params
func (o *ObjectUserAssignmentParams) SetBody(body *models.ObjectUserAssignment) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ObjectUserAssignmentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

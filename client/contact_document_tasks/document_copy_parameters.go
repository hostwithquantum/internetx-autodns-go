// Code generated by go-swagger; DO NOT EDIT.

package contact_document_tasks

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

// NewDocumentCopyParams creates a new DocumentCopyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDocumentCopyParams() *DocumentCopyParams {
	return &DocumentCopyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDocumentCopyParamsWithTimeout creates a new DocumentCopyParams object
// with the ability to set a timeout on a request.
func NewDocumentCopyParamsWithTimeout(timeout time.Duration) *DocumentCopyParams {
	return &DocumentCopyParams{
		timeout: timeout,
	}
}

// NewDocumentCopyParamsWithContext creates a new DocumentCopyParams object
// with the ability to set a context for a request.
func NewDocumentCopyParamsWithContext(ctx context.Context) *DocumentCopyParams {
	return &DocumentCopyParams{
		Context: ctx,
	}
}

// NewDocumentCopyParamsWithHTTPClient creates a new DocumentCopyParams object
// with the ability to set a custom HTTPClient for a request.
func NewDocumentCopyParamsWithHTTPClient(client *http.Client) *DocumentCopyParams {
	return &DocumentCopyParams{
		HTTPClient: client,
	}
}

/*
DocumentCopyParams contains all the parameters to send to the API endpoint

	for the document copy operation.

	Typically these are written to a http.Request.
*/
type DocumentCopyParams struct {

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

	   contactDocument
	*/
	Body *models.ContactDocument

	// File.
	File *string

	// ID.
	//
	// Format: int64
	ID int64

	// Type.
	Type string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the document copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentCopyParams) WithDefaults() *DocumentCopyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the document copy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DocumentCopyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the document copy params
func (o *DocumentCopyParams) WithTimeout(timeout time.Duration) *DocumentCopyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the document copy params
func (o *DocumentCopyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the document copy params
func (o *DocumentCopyParams) WithContext(ctx context.Context) *DocumentCopyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the document copy params
func (o *DocumentCopyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the document copy params
func (o *DocumentCopyParams) WithHTTPClient(client *http.Client) *DocumentCopyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the document copy params
func (o *DocumentCopyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the document copy params
func (o *DocumentCopyParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *DocumentCopyParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the document copy params
func (o *DocumentCopyParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *DocumentCopyParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotContext(xDomainrobotContext *int32) *DocumentCopyParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *DocumentCopyParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *DocumentCopyParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *DocumentCopyParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *DocumentCopyParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *DocumentCopyParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *DocumentCopyParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *DocumentCopyParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *DocumentCopyParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the document copy params
func (o *DocumentCopyParams) WithXDomainrobotWS(xDomainrobotWS *string) *DocumentCopyParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the document copy params
func (o *DocumentCopyParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the document copy params
func (o *DocumentCopyParams) WithBody(body *models.ContactDocument) *DocumentCopyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the document copy params
func (o *DocumentCopyParams) SetBody(body *models.ContactDocument) {
	o.Body = body
}

// WithFile adds the file to the document copy params
func (o *DocumentCopyParams) WithFile(file *string) *DocumentCopyParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the document copy params
func (o *DocumentCopyParams) SetFile(file *string) {
	o.File = file
}

// WithID adds the id to the document copy params
func (o *DocumentCopyParams) WithID(id int64) *DocumentCopyParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the document copy params
func (o *DocumentCopyParams) SetID(id int64) {
	o.ID = id
}

// WithType adds the typeVar to the document copy params
func (o *DocumentCopyParams) WithType(typeVar string) *DocumentCopyParams {
	o.SetType(typeVar)
	return o
}

// SetType adds the type to the document copy params
func (o *DocumentCopyParams) SetType(typeVar string) {
	o.Type = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *DocumentCopyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.File != nil {

		// form param file
		var frFile string
		if o.File != nil {
			frFile = *o.File
		}
		fFile := frFile
		if fFile != "" {
			if err := r.SetFormParam("file", fFile); err != nil {
				return err
			}
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param type
	if err := r.SetPathParam("type", o.Type); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

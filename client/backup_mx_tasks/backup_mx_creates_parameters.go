// Code generated by go-swagger; DO NOT EDIT.

package backup_mx_tasks

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

// NewBackupMxCreatesParams creates a new BackupMxCreatesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBackupMxCreatesParams() *BackupMxCreatesParams {
	return &BackupMxCreatesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBackupMxCreatesParamsWithTimeout creates a new BackupMxCreatesParams object
// with the ability to set a timeout on a request.
func NewBackupMxCreatesParamsWithTimeout(timeout time.Duration) *BackupMxCreatesParams {
	return &BackupMxCreatesParams{
		timeout: timeout,
	}
}

// NewBackupMxCreatesParamsWithContext creates a new BackupMxCreatesParams object
// with the ability to set a context for a request.
func NewBackupMxCreatesParamsWithContext(ctx context.Context) *BackupMxCreatesParams {
	return &BackupMxCreatesParams{
		Context: ctx,
	}
}

// NewBackupMxCreatesParamsWithHTTPClient creates a new BackupMxCreatesParams object
// with the ability to set a custom HTTPClient for a request.
func NewBackupMxCreatesParamsWithHTTPClient(client *http.Client) *BackupMxCreatesParams {
	return &BackupMxCreatesParams{
		HTTPClient: client,
	}
}

/*
BackupMxCreatesParams contains all the parameters to send to the API endpoint

	for the backup mx creates operation.

	Typically these are written to a http.Request.
*/
type BackupMxCreatesParams struct {

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

	   backupMxs
	*/
	Body *models.BulkBackupMxPostRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the backup mx creates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupMxCreatesParams) WithDefaults() *BackupMxCreatesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the backup mx creates params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BackupMxCreatesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the backup mx creates params
func (o *BackupMxCreatesParams) WithTimeout(timeout time.Duration) *BackupMxCreatesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the backup mx creates params
func (o *BackupMxCreatesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the backup mx creates params
func (o *BackupMxCreatesParams) WithContext(ctx context.Context) *BackupMxCreatesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the backup mx creates params
func (o *BackupMxCreatesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the backup mx creates params
func (o *BackupMxCreatesParams) WithHTTPClient(client *http.Client) *BackupMxCreatesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the backup mx creates params
func (o *BackupMxCreatesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXDomainrobot2FAToken adds the xDomainrobot2FAToken to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobot2FAToken(xDomainrobot2FAToken *int32) *BackupMxCreatesParams {
	o.SetXDomainrobot2FAToken(xDomainrobot2FAToken)
	return o
}

// SetXDomainrobot2FAToken adds the xDomainrobot2FAToken to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobot2FAToken(xDomainrobot2FAToken *int32) {
	o.XDomainrobot2FAToken = xDomainrobot2FAToken
}

// WithXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) *BackupMxCreatesParams {
	o.SetXDomainrobotBulkLimit(xDomainrobotBulkLimit)
	return o
}

// SetXDomainrobotBulkLimit adds the xDomainrobotBulkLimit to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotBulkLimit(xDomainrobotBulkLimit *int32) {
	o.XDomainrobotBulkLimit = xDomainrobotBulkLimit
}

// WithXDomainrobotContext adds the xDomainrobotContext to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotContext(xDomainrobotContext *int32) *BackupMxCreatesParams {
	o.SetXDomainrobotContext(xDomainrobotContext)
	return o
}

// SetXDomainrobotContext adds the xDomainrobotContext to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotContext(xDomainrobotContext *int32) {
	o.XDomainrobotContext = xDomainrobotContext
}

// WithXDomainrobotDemo adds the xDomainrobotDemo to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotDemo(xDomainrobotDemo *bool) *BackupMxCreatesParams {
	o.SetXDomainrobotDemo(xDomainrobotDemo)
	return o
}

// SetXDomainrobotDemo adds the xDomainrobotDemo to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotDemo(xDomainrobotDemo *bool) {
	o.XDomainrobotDemo = xDomainrobotDemo
}

// WithXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) *BackupMxCreatesParams {
	o.SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin)
	return o
}

// SetXDomainrobotDomainSafePin adds the xDomainrobotDomainSafePin to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotDomainSafePin(xDomainrobotDomainSafePin *string) {
	o.XDomainrobotDomainSafePin = xDomainrobotDomainSafePin
}

// WithXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) *BackupMxCreatesParams {
	o.SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan)
	return o
}

// SetXDomainrobotDomainSafeTan adds the xDomainrobotDomainSafeTan to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotDomainSafeTan(xDomainrobotDomainSafeTan *string) {
	o.XDomainrobotDomainSafeTan = xDomainrobotDomainSafeTan
}

// WithXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) *BackupMxCreatesParams {
	o.SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction)
	return o
}

// SetXDomainrobotDomainSafeTransaction adds the xDomainrobotDomainSafeTransaction to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotDomainSafeTransaction(xDomainrobotDomainSafeTransaction *string) {
	o.XDomainrobotDomainSafeTransaction = xDomainrobotDomainSafeTransaction
}

// WithXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) *BackupMxCreatesParams {
	o.SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire)
	return o
}

// SetXDomainrobotDomainSafeTransactionExpire adds the xDomainrobotDomainSafeTransactionExpire to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotDomainSafeTransactionExpire(xDomainrobotDomainSafeTransactionExpire *strfmt.DateTime) {
	o.XDomainrobotDomainSafeTransactionExpire = xDomainrobotDomainSafeTransactionExpire
}

// WithXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) *BackupMxCreatesParams {
	o.SetXDomainrobotOwnerContext(xDomainrobotOwnerContext)
	return o
}

// SetXDomainrobotOwnerContext adds the xDomainrobotOwnerContext to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotOwnerContext(xDomainrobotOwnerContext *int32) {
	o.XDomainrobotOwnerContext = xDomainrobotOwnerContext
}

// WithXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) *BackupMxCreatesParams {
	o.SetXDomainrobotOwnerUser(xDomainrobotOwnerUser)
	return o
}

// SetXDomainrobotOwnerUser adds the xDomainrobotOwnerUser to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotOwnerUser(xDomainrobotOwnerUser *string) {
	o.XDomainrobotOwnerUser = xDomainrobotOwnerUser
}

// WithXDomainrobotSessionID adds the xDomainrobotSessionID to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotSessionID(xDomainrobotSessionID *string) *BackupMxCreatesParams {
	o.SetXDomainrobotSessionID(xDomainrobotSessionID)
	return o
}

// SetXDomainrobotSessionID adds the xDomainrobotSessionId to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotSessionID(xDomainrobotSessionID *string) {
	o.XDomainrobotSessionID = xDomainrobotSessionID
}

// WithXDomainrobotWS adds the xDomainrobotWS to the backup mx creates params
func (o *BackupMxCreatesParams) WithXDomainrobotWS(xDomainrobotWS *string) *BackupMxCreatesParams {
	o.SetXDomainrobotWS(xDomainrobotWS)
	return o
}

// SetXDomainrobotWS adds the xDomainrobotWS to the backup mx creates params
func (o *BackupMxCreatesParams) SetXDomainrobotWS(xDomainrobotWS *string) {
	o.XDomainrobotWS = xDomainrobotWS
}

// WithBody adds the body to the backup mx creates params
func (o *BackupMxCreatesParams) WithBody(body *models.BulkBackupMxPostRequest) *BackupMxCreatesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the backup mx creates params
func (o *BackupMxCreatesParams) SetBody(body *models.BulkBackupMxPostRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BackupMxCreatesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

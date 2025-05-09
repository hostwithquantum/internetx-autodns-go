// Code generated by go-swagger; DO NOT EDIT.

package tmch_mark_tasks

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
)

// NewTmchMarkConfirmParams creates a new TmchMarkConfirmParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTmchMarkConfirmParams() *TmchMarkConfirmParams {
	return &TmchMarkConfirmParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkConfirmParamsWithTimeout creates a new TmchMarkConfirmParams object
// with the ability to set a timeout on a request.
func NewTmchMarkConfirmParamsWithTimeout(timeout time.Duration) *TmchMarkConfirmParams {
	return &TmchMarkConfirmParams{
		timeout: timeout,
	}
}

// NewTmchMarkConfirmParamsWithContext creates a new TmchMarkConfirmParams object
// with the ability to set a context for a request.
func NewTmchMarkConfirmParamsWithContext(ctx context.Context) *TmchMarkConfirmParams {
	return &TmchMarkConfirmParams{
		Context: ctx,
	}
}

// NewTmchMarkConfirmParamsWithHTTPClient creates a new TmchMarkConfirmParams object
// with the ability to set a custom HTTPClient for a request.
func NewTmchMarkConfirmParamsWithHTTPClient(client *http.Client) *TmchMarkConfirmParams {
	return &TmchMarkConfirmParams{
		HTTPClient: client,
	}
}

/*
TmchMarkConfirmParams contains all the parameters to send to the API endpoint

	for the tmch mark confirm operation.

	Typically these are written to a http.Request.
*/
type TmchMarkConfirmParams struct {

	// Reference.
	Reference string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the tmch mark confirm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TmchMarkConfirmParams) WithDefaults() *TmchMarkConfirmParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the tmch mark confirm params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TmchMarkConfirmParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the tmch mark confirm params
func (o *TmchMarkConfirmParams) WithTimeout(timeout time.Duration) *TmchMarkConfirmParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tmch mark confirm params
func (o *TmchMarkConfirmParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tmch mark confirm params
func (o *TmchMarkConfirmParams) WithContext(ctx context.Context) *TmchMarkConfirmParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tmch mark confirm params
func (o *TmchMarkConfirmParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tmch mark confirm params
func (o *TmchMarkConfirmParams) WithHTTPClient(client *http.Client) *TmchMarkConfirmParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tmch mark confirm params
func (o *TmchMarkConfirmParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReference adds the reference to the tmch mark confirm params
func (o *TmchMarkConfirmParams) WithReference(reference string) *TmchMarkConfirmParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the tmch mark confirm params
func (o *TmchMarkConfirmParams) SetReference(reference string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkConfirmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

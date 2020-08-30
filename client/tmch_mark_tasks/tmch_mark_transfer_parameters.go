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

// NewTmchMarkTransferParams creates a new TmchMarkTransferParams object
// with the default values initialized.
func NewTmchMarkTransferParams() *TmchMarkTransferParams {
	var ()
	return &TmchMarkTransferParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkTransferParamsWithTimeout creates a new TmchMarkTransferParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTmchMarkTransferParamsWithTimeout(timeout time.Duration) *TmchMarkTransferParams {
	var ()
	return &TmchMarkTransferParams{

		timeout: timeout,
	}
}

// NewTmchMarkTransferParamsWithContext creates a new TmchMarkTransferParams object
// with the default values initialized, and the ability to set a context for a request
func NewTmchMarkTransferParamsWithContext(ctx context.Context) *TmchMarkTransferParams {
	var ()
	return &TmchMarkTransferParams{

		Context: ctx,
	}
}

// NewTmchMarkTransferParamsWithHTTPClient creates a new TmchMarkTransferParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTmchMarkTransferParamsWithHTTPClient(client *http.Client) *TmchMarkTransferParams {
	var ()
	return &TmchMarkTransferParams{
		HTTPClient: client,
	}
}

/*TmchMarkTransferParams contains all the parameters to send to the API endpoint
for the tmch mark transfer operation typically these are written to a http.Request
*/
type TmchMarkTransferParams struct {

	/*Reference*/
	Reference string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tmch mark transfer params
func (o *TmchMarkTransferParams) WithTimeout(timeout time.Duration) *TmchMarkTransferParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tmch mark transfer params
func (o *TmchMarkTransferParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tmch mark transfer params
func (o *TmchMarkTransferParams) WithContext(ctx context.Context) *TmchMarkTransferParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tmch mark transfer params
func (o *TmchMarkTransferParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tmch mark transfer params
func (o *TmchMarkTransferParams) WithHTTPClient(client *http.Client) *TmchMarkTransferParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tmch mark transfer params
func (o *TmchMarkTransferParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReference adds the reference to the tmch mark transfer params
func (o *TmchMarkTransferParams) WithReference(reference string) *TmchMarkTransferParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the tmch mark transfer params
func (o *TmchMarkTransferParams) SetReference(reference string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkTransferParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

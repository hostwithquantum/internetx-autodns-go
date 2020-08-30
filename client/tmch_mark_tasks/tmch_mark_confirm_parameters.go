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

// NewTmchMarkConfirmParams creates a new TmchMarkConfirmParams object
// with the default values initialized.
func NewTmchMarkConfirmParams() *TmchMarkConfirmParams {
	var ()
	return &TmchMarkConfirmParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkConfirmParamsWithTimeout creates a new TmchMarkConfirmParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTmchMarkConfirmParamsWithTimeout(timeout time.Duration) *TmchMarkConfirmParams {
	var ()
	return &TmchMarkConfirmParams{

		timeout: timeout,
	}
}

// NewTmchMarkConfirmParamsWithContext creates a new TmchMarkConfirmParams object
// with the default values initialized, and the ability to set a context for a request
func NewTmchMarkConfirmParamsWithContext(ctx context.Context) *TmchMarkConfirmParams {
	var ()
	return &TmchMarkConfirmParams{

		Context: ctx,
	}
}

// NewTmchMarkConfirmParamsWithHTTPClient creates a new TmchMarkConfirmParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTmchMarkConfirmParamsWithHTTPClient(client *http.Client) *TmchMarkConfirmParams {
	var ()
	return &TmchMarkConfirmParams{
		HTTPClient: client,
	}
}

/*TmchMarkConfirmParams contains all the parameters to send to the API endpoint
for the tmch mark confirm operation typically these are written to a http.Request
*/
type TmchMarkConfirmParams struct {

	/*Body
	  reference

	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
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

// WithBody adds the body to the tmch mark confirm params
func (o *TmchMarkConfirmParams) WithBody(body string) *TmchMarkConfirmParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tmch mark confirm params
func (o *TmchMarkConfirmParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkConfirmParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

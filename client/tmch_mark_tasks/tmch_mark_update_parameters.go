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

// NewTmchMarkUpdateParams creates a new TmchMarkUpdateParams object
// with the default values initialized.
func NewTmchMarkUpdateParams() *TmchMarkUpdateParams {
	var ()
	return &TmchMarkUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkUpdateParamsWithTimeout creates a new TmchMarkUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTmchMarkUpdateParamsWithTimeout(timeout time.Duration) *TmchMarkUpdateParams {
	var ()
	return &TmchMarkUpdateParams{

		timeout: timeout,
	}
}

// NewTmchMarkUpdateParamsWithContext creates a new TmchMarkUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTmchMarkUpdateParamsWithContext(ctx context.Context) *TmchMarkUpdateParams {
	var ()
	return &TmchMarkUpdateParams{

		Context: ctx,
	}
}

// NewTmchMarkUpdateParamsWithHTTPClient creates a new TmchMarkUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTmchMarkUpdateParamsWithHTTPClient(client *http.Client) *TmchMarkUpdateParams {
	var ()
	return &TmchMarkUpdateParams{
		HTTPClient: client,
	}
}

/*TmchMarkUpdateParams contains all the parameters to send to the API endpoint
for the tmch mark update operation typically these are written to a http.Request
*/
type TmchMarkUpdateParams struct {

	/*Body
	  reference

	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tmch mark update params
func (o *TmchMarkUpdateParams) WithTimeout(timeout time.Duration) *TmchMarkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tmch mark update params
func (o *TmchMarkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tmch mark update params
func (o *TmchMarkUpdateParams) WithContext(ctx context.Context) *TmchMarkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tmch mark update params
func (o *TmchMarkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tmch mark update params
func (o *TmchMarkUpdateParams) WithHTTPClient(client *http.Client) *TmchMarkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tmch mark update params
func (o *TmchMarkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tmch mark update params
func (o *TmchMarkUpdateParams) WithBody(body string) *TmchMarkUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tmch mark update params
func (o *TmchMarkUpdateParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

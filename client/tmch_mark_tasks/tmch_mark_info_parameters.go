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

// NewTmchMarkInfoParams creates a new TmchMarkInfoParams object
// with the default values initialized.
func NewTmchMarkInfoParams() *TmchMarkInfoParams {
	var ()
	return &TmchMarkInfoParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkInfoParamsWithTimeout creates a new TmchMarkInfoParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTmchMarkInfoParamsWithTimeout(timeout time.Duration) *TmchMarkInfoParams {
	var ()
	return &TmchMarkInfoParams{

		timeout: timeout,
	}
}

// NewTmchMarkInfoParamsWithContext creates a new TmchMarkInfoParams object
// with the default values initialized, and the ability to set a context for a request
func NewTmchMarkInfoParamsWithContext(ctx context.Context) *TmchMarkInfoParams {
	var ()
	return &TmchMarkInfoParams{

		Context: ctx,
	}
}

// NewTmchMarkInfoParamsWithHTTPClient creates a new TmchMarkInfoParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTmchMarkInfoParamsWithHTTPClient(client *http.Client) *TmchMarkInfoParams {
	var ()
	return &TmchMarkInfoParams{
		HTTPClient: client,
	}
}

/*TmchMarkInfoParams contains all the parameters to send to the API endpoint
for the tmch mark info operation typically these are written to a http.Request
*/
type TmchMarkInfoParams struct {

	/*Body
	  reference

	*/
	Body string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tmch mark info params
func (o *TmchMarkInfoParams) WithTimeout(timeout time.Duration) *TmchMarkInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tmch mark info params
func (o *TmchMarkInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tmch mark info params
func (o *TmchMarkInfoParams) WithContext(ctx context.Context) *TmchMarkInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tmch mark info params
func (o *TmchMarkInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tmch mark info params
func (o *TmchMarkInfoParams) WithHTTPClient(client *http.Client) *TmchMarkInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tmch mark info params
func (o *TmchMarkInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tmch mark info params
func (o *TmchMarkInfoParams) WithBody(body string) *TmchMarkInfoParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tmch mark info params
func (o *TmchMarkInfoParams) SetBody(body string) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

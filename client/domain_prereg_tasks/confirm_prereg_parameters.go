// Code generated by go-swagger; DO NOT EDIT.

package domain_prereg_tasks

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

// NewConfirmPreregParams creates a new ConfirmPreregParams object
// with the default values initialized.
func NewConfirmPreregParams() *ConfirmPreregParams {
	var ()
	return &ConfirmPreregParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewConfirmPreregParamsWithTimeout creates a new ConfirmPreregParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewConfirmPreregParamsWithTimeout(timeout time.Duration) *ConfirmPreregParams {
	var ()
	return &ConfirmPreregParams{

		timeout: timeout,
	}
}

// NewConfirmPreregParamsWithContext creates a new ConfirmPreregParams object
// with the default values initialized, and the ability to set a context for a request
func NewConfirmPreregParamsWithContext(ctx context.Context) *ConfirmPreregParams {
	var ()
	return &ConfirmPreregParams{

		Context: ctx,
	}
}

// NewConfirmPreregParamsWithHTTPClient creates a new ConfirmPreregParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewConfirmPreregParamsWithHTTPClient(client *http.Client) *ConfirmPreregParams {
	var ()
	return &ConfirmPreregParams{
		HTTPClient: client,
	}
}

/*ConfirmPreregParams contains all the parameters to send to the API endpoint
for the confirm prereg operation typically these are written to a http.Request
*/
type ConfirmPreregParams struct {

	/*Reference*/
	Reference string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the confirm prereg params
func (o *ConfirmPreregParams) WithTimeout(timeout time.Duration) *ConfirmPreregParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the confirm prereg params
func (o *ConfirmPreregParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the confirm prereg params
func (o *ConfirmPreregParams) WithContext(ctx context.Context) *ConfirmPreregParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the confirm prereg params
func (o *ConfirmPreregParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the confirm prereg params
func (o *ConfirmPreregParams) WithHTTPClient(client *http.Client) *ConfirmPreregParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the confirm prereg params
func (o *ConfirmPreregParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReference adds the reference to the confirm prereg params
func (o *ConfirmPreregParams) WithReference(reference string) *ConfirmPreregParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the confirm prereg params
func (o *ConfirmPreregParams) SetReference(reference string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *ConfirmPreregParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
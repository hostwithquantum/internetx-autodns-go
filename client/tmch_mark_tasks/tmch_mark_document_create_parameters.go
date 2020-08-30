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
	"github.com/go-openapi/swag"

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// NewTmchMarkDocumentCreateParams creates a new TmchMarkDocumentCreateParams object
// with the default values initialized.
func NewTmchMarkDocumentCreateParams() *TmchMarkDocumentCreateParams {
	var ()
	return &TmchMarkDocumentCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTmchMarkDocumentCreateParamsWithTimeout creates a new TmchMarkDocumentCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTmchMarkDocumentCreateParamsWithTimeout(timeout time.Duration) *TmchMarkDocumentCreateParams {
	var ()
	return &TmchMarkDocumentCreateParams{

		timeout: timeout,
	}
}

// NewTmchMarkDocumentCreateParamsWithContext creates a new TmchMarkDocumentCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewTmchMarkDocumentCreateParamsWithContext(ctx context.Context) *TmchMarkDocumentCreateParams {
	var ()
	return &TmchMarkDocumentCreateParams{

		Context: ctx,
	}
}

// NewTmchMarkDocumentCreateParamsWithHTTPClient creates a new TmchMarkDocumentCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTmchMarkDocumentCreateParamsWithHTTPClient(client *http.Client) *TmchMarkDocumentCreateParams {
	var ()
	return &TmchMarkDocumentCreateParams{
		HTTPClient: client,
	}
}

/*TmchMarkDocumentCreateParams contains all the parameters to send to the API endpoint
for the tmch mark document create operation typically these are written to a http.Request
*/
type TmchMarkDocumentCreateParams struct {

	/*Body
	  tmchMark

	*/
	Body *models.TmchMarkDocument
	/*Keys
	  The query parameter to enable assignment replacment.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) WithTimeout(timeout time.Duration) *TmchMarkDocumentCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) WithContext(ctx context.Context) *TmchMarkDocumentCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) WithHTTPClient(client *http.Client) *TmchMarkDocumentCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) WithBody(body *models.TmchMarkDocument) *TmchMarkDocumentCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) SetBody(body *models.TmchMarkDocument) {
	o.Body = body
}

// WithKeys adds the keys to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) WithKeys(keys []string) *TmchMarkDocumentCreateParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the tmch mark document create params
func (o *TmchMarkDocumentCreateParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *TmchMarkDocumentCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	valuesKeys := o.Keys

	joinedKeys := swag.JoinByFormat(valuesKeys, "multi")
	// query array param keys
	if err := r.SetQueryParam("keys", joinedKeys...); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

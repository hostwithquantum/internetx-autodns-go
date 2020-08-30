// Code generated by go-swagger; DO NOT EDIT.

package domain_safe_tasks

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

// NewSafeObjectListParams creates a new SafeObjectListParams object
// with the default values initialized.
func NewSafeObjectListParams() *SafeObjectListParams {
	var ()
	return &SafeObjectListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSafeObjectListParamsWithTimeout creates a new SafeObjectListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSafeObjectListParamsWithTimeout(timeout time.Duration) *SafeObjectListParams {
	var ()
	return &SafeObjectListParams{

		timeout: timeout,
	}
}

// NewSafeObjectListParamsWithContext creates a new SafeObjectListParams object
// with the default values initialized, and the ability to set a context for a request
func NewSafeObjectListParamsWithContext(ctx context.Context) *SafeObjectListParams {
	var ()
	return &SafeObjectListParams{

		Context: ctx,
	}
}

// NewSafeObjectListParamsWithHTTPClient creates a new SafeObjectListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSafeObjectListParamsWithHTTPClient(client *http.Client) *SafeObjectListParams {
	var ()
	return &SafeObjectListParams{
		HTTPClient: client,
	}
}

/*SafeObjectListParams contains all the parameters to send to the API endpoint
for the safe object list operation typically these are written to a http.Request
*/
type SafeObjectListParams struct {

	/*Body
	  query

	*/
	Body *models.Query
	/*Keys
	  The query parameter to inquire additional details.

	*/
	Keys []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the safe object list params
func (o *SafeObjectListParams) WithTimeout(timeout time.Duration) *SafeObjectListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the safe object list params
func (o *SafeObjectListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the safe object list params
func (o *SafeObjectListParams) WithContext(ctx context.Context) *SafeObjectListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the safe object list params
func (o *SafeObjectListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the safe object list params
func (o *SafeObjectListParams) WithHTTPClient(client *http.Client) *SafeObjectListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the safe object list params
func (o *SafeObjectListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the safe object list params
func (o *SafeObjectListParams) WithBody(body *models.Query) *SafeObjectListParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the safe object list params
func (o *SafeObjectListParams) SetBody(body *models.Query) {
	o.Body = body
}

// WithKeys adds the keys to the safe object list params
func (o *SafeObjectListParams) WithKeys(keys []string) *SafeObjectListParams {
	o.SetKeys(keys)
	return o
}

// SetKeys adds the keys to the safe object list params
func (o *SafeObjectListParams) SetKeys(keys []string) {
	o.Keys = keys
}

// WriteToRequest writes these params to a swagger request
func (o *SafeObjectListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
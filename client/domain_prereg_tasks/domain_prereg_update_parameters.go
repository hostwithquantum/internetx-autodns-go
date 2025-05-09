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

	"github.com/hostwithquantum/internetx-autodns-go/models"
)

// NewDomainPreregUpdateParams creates a new DomainPreregUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDomainPreregUpdateParams() *DomainPreregUpdateParams {
	return &DomainPreregUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDomainPreregUpdateParamsWithTimeout creates a new DomainPreregUpdateParams object
// with the ability to set a timeout on a request.
func NewDomainPreregUpdateParamsWithTimeout(timeout time.Duration) *DomainPreregUpdateParams {
	return &DomainPreregUpdateParams{
		timeout: timeout,
	}
}

// NewDomainPreregUpdateParamsWithContext creates a new DomainPreregUpdateParams object
// with the ability to set a context for a request.
func NewDomainPreregUpdateParamsWithContext(ctx context.Context) *DomainPreregUpdateParams {
	return &DomainPreregUpdateParams{
		Context: ctx,
	}
}

// NewDomainPreregUpdateParamsWithHTTPClient creates a new DomainPreregUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDomainPreregUpdateParamsWithHTTPClient(client *http.Client) *DomainPreregUpdateParams {
	return &DomainPreregUpdateParams{
		HTTPClient: client,
	}
}

/*
DomainPreregUpdateParams contains all the parameters to send to the API endpoint

	for the domain prereg update operation.

	Typically these are written to a http.Request.
*/
type DomainPreregUpdateParams struct {

	/* Body.

	   domainPrereg
	*/
	Body *models.DomainPrereg

	// Reference.
	Reference string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the domain prereg update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPreregUpdateParams) WithDefaults() *DomainPreregUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the domain prereg update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DomainPreregUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the domain prereg update params
func (o *DomainPreregUpdateParams) WithTimeout(timeout time.Duration) *DomainPreregUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the domain prereg update params
func (o *DomainPreregUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the domain prereg update params
func (o *DomainPreregUpdateParams) WithContext(ctx context.Context) *DomainPreregUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the domain prereg update params
func (o *DomainPreregUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the domain prereg update params
func (o *DomainPreregUpdateParams) WithHTTPClient(client *http.Client) *DomainPreregUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the domain prereg update params
func (o *DomainPreregUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the domain prereg update params
func (o *DomainPreregUpdateParams) WithBody(body *models.DomainPrereg) *DomainPreregUpdateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the domain prereg update params
func (o *DomainPreregUpdateParams) SetBody(body *models.DomainPrereg) {
	o.Body = body
}

// WithReference adds the reference to the domain prereg update params
func (o *DomainPreregUpdateParams) WithReference(reference string) *DomainPreregUpdateParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the domain prereg update params
func (o *DomainPreregUpdateParams) SetReference(reference string) {
	o.Reference = reference
}

// WriteToRequest writes these params to a swagger request
func (o *DomainPreregUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

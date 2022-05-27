// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewAddPostgresExporterParams creates a new AddPostgresExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAddPostgresExporterParams() *AddPostgresExporterParams {
	return &AddPostgresExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAddPostgresExporterParamsWithTimeout creates a new AddPostgresExporterParams object
// with the ability to set a timeout on a request.
func NewAddPostgresExporterParamsWithTimeout(timeout time.Duration) *AddPostgresExporterParams {
	return &AddPostgresExporterParams{
		timeout: timeout,
	}
}

// NewAddPostgresExporterParamsWithContext creates a new AddPostgresExporterParams object
// with the ability to set a context for a request.
func NewAddPostgresExporterParamsWithContext(ctx context.Context) *AddPostgresExporterParams {
	return &AddPostgresExporterParams{
		Context: ctx,
	}
}

// NewAddPostgresExporterParamsWithHTTPClient creates a new AddPostgresExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewAddPostgresExporterParamsWithHTTPClient(client *http.Client) *AddPostgresExporterParams {
	return &AddPostgresExporterParams{
		HTTPClient: client,
	}
}

/* AddPostgresExporterParams contains all the parameters to send to the API endpoint
   for the add postgres exporter operation.

   Typically these are written to a http.Request.
*/
type AddPostgresExporterParams struct {
	// Body.
	Body AddPostgresExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the add postgres exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPostgresExporterParams) WithDefaults() *AddPostgresExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the add postgres exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AddPostgresExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the add postgres exporter params
func (o *AddPostgresExporterParams) WithTimeout(timeout time.Duration) *AddPostgresExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add postgres exporter params
func (o *AddPostgresExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add postgres exporter params
func (o *AddPostgresExporterParams) WithContext(ctx context.Context) *AddPostgresExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add postgres exporter params
func (o *AddPostgresExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add postgres exporter params
func (o *AddPostgresExporterParams) WithHTTPClient(client *http.Client) *AddPostgresExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add postgres exporter params
func (o *AddPostgresExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add postgres exporter params
func (o *AddPostgresExporterParams) WithBody(body AddPostgresExporterBody) *AddPostgresExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add postgres exporter params
func (o *AddPostgresExporterParams) SetBody(body AddPostgresExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddPostgresExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

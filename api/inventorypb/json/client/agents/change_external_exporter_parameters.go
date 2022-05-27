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

// NewChangeExternalExporterParams creates a new ChangeExternalExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeExternalExporterParams() *ChangeExternalExporterParams {
	return &ChangeExternalExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeExternalExporterParamsWithTimeout creates a new ChangeExternalExporterParams object
// with the ability to set a timeout on a request.
func NewChangeExternalExporterParamsWithTimeout(timeout time.Duration) *ChangeExternalExporterParams {
	return &ChangeExternalExporterParams{
		timeout: timeout,
	}
}

// NewChangeExternalExporterParamsWithContext creates a new ChangeExternalExporterParams object
// with the ability to set a context for a request.
func NewChangeExternalExporterParamsWithContext(ctx context.Context) *ChangeExternalExporterParams {
	return &ChangeExternalExporterParams{
		Context: ctx,
	}
}

// NewChangeExternalExporterParamsWithHTTPClient creates a new ChangeExternalExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeExternalExporterParamsWithHTTPClient(client *http.Client) *ChangeExternalExporterParams {
	return &ChangeExternalExporterParams{
		HTTPClient: client,
	}
}

/* ChangeExternalExporterParams contains all the parameters to send to the API endpoint
   for the change external exporter operation.

   Typically these are written to a http.Request.
*/
type ChangeExternalExporterParams struct {
	// Body.
	Body ChangeExternalExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change external exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeExternalExporterParams) WithDefaults() *ChangeExternalExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change external exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeExternalExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change external exporter params
func (o *ChangeExternalExporterParams) WithTimeout(timeout time.Duration) *ChangeExternalExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change external exporter params
func (o *ChangeExternalExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change external exporter params
func (o *ChangeExternalExporterParams) WithContext(ctx context.Context) *ChangeExternalExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change external exporter params
func (o *ChangeExternalExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change external exporter params
func (o *ChangeExternalExporterParams) WithHTTPClient(client *http.Client) *ChangeExternalExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change external exporter params
func (o *ChangeExternalExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change external exporter params
func (o *ChangeExternalExporterParams) WithBody(body ChangeExternalExporterBody) *ChangeExternalExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change external exporter params
func (o *ChangeExternalExporterParams) SetBody(body ChangeExternalExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeExternalExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

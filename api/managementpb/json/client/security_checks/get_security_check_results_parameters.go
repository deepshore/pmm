// Code generated by go-swagger; DO NOT EDIT.

package security_checks

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

// NewGetSecurityCheckResultsParams creates a new GetSecurityCheckResultsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetSecurityCheckResultsParams() *GetSecurityCheckResultsParams {
	return &GetSecurityCheckResultsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetSecurityCheckResultsParamsWithTimeout creates a new GetSecurityCheckResultsParams object
// with the ability to set a timeout on a request.
func NewGetSecurityCheckResultsParamsWithTimeout(timeout time.Duration) *GetSecurityCheckResultsParams {
	return &GetSecurityCheckResultsParams{
		timeout: timeout,
	}
}

// NewGetSecurityCheckResultsParamsWithContext creates a new GetSecurityCheckResultsParams object
// with the ability to set a context for a request.
func NewGetSecurityCheckResultsParamsWithContext(ctx context.Context) *GetSecurityCheckResultsParams {
	return &GetSecurityCheckResultsParams{
		Context: ctx,
	}
}

// NewGetSecurityCheckResultsParamsWithHTTPClient creates a new GetSecurityCheckResultsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetSecurityCheckResultsParamsWithHTTPClient(client *http.Client) *GetSecurityCheckResultsParams {
	return &GetSecurityCheckResultsParams{
		HTTPClient: client,
	}
}

/* GetSecurityCheckResultsParams contains all the parameters to send to the API endpoint
   for the get security check results operation.

   Typically these are written to a http.Request.
*/
type GetSecurityCheckResultsParams struct {
	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get security check results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityCheckResultsParams) WithDefaults() *GetSecurityCheckResultsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get security check results params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetSecurityCheckResultsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get security check results params
func (o *GetSecurityCheckResultsParams) WithTimeout(timeout time.Duration) *GetSecurityCheckResultsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get security check results params
func (o *GetSecurityCheckResultsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get security check results params
func (o *GetSecurityCheckResultsParams) WithContext(ctx context.Context) *GetSecurityCheckResultsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get security check results params
func (o *GetSecurityCheckResultsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get security check results params
func (o *GetSecurityCheckResultsParams) WithHTTPClient(client *http.Client) *GetSecurityCheckResultsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get security check results params
func (o *GetSecurityCheckResultsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get security check results params
func (o *GetSecurityCheckResultsParams) WithBody(body interface{}) *GetSecurityCheckResultsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get security check results params
func (o *GetSecurityCheckResultsParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetSecurityCheckResultsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

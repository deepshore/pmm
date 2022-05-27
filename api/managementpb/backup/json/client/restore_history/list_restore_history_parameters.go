// Code generated by go-swagger; DO NOT EDIT.

package restore_history

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

// NewListRestoreHistoryParams creates a new ListRestoreHistoryParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRestoreHistoryParams() *ListRestoreHistoryParams {
	return &ListRestoreHistoryParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRestoreHistoryParamsWithTimeout creates a new ListRestoreHistoryParams object
// with the ability to set a timeout on a request.
func NewListRestoreHistoryParamsWithTimeout(timeout time.Duration) *ListRestoreHistoryParams {
	return &ListRestoreHistoryParams{
		timeout: timeout,
	}
}

// NewListRestoreHistoryParamsWithContext creates a new ListRestoreHistoryParams object
// with the ability to set a context for a request.
func NewListRestoreHistoryParamsWithContext(ctx context.Context) *ListRestoreHistoryParams {
	return &ListRestoreHistoryParams{
		Context: ctx,
	}
}

// NewListRestoreHistoryParamsWithHTTPClient creates a new ListRestoreHistoryParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRestoreHistoryParamsWithHTTPClient(client *http.Client) *ListRestoreHistoryParams {
	return &ListRestoreHistoryParams{
		HTTPClient: client,
	}
}

/* ListRestoreHistoryParams contains all the parameters to send to the API endpoint
   for the list restore history operation.

   Typically these are written to a http.Request.
*/
type ListRestoreHistoryParams struct {
	// Body.
	Body interface{}

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list restore history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRestoreHistoryParams) WithDefaults() *ListRestoreHistoryParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list restore history params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRestoreHistoryParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list restore history params
func (o *ListRestoreHistoryParams) WithTimeout(timeout time.Duration) *ListRestoreHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list restore history params
func (o *ListRestoreHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list restore history params
func (o *ListRestoreHistoryParams) WithContext(ctx context.Context) *ListRestoreHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list restore history params
func (o *ListRestoreHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list restore history params
func (o *ListRestoreHistoryParams) WithHTTPClient(client *http.Client) *ListRestoreHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list restore history params
func (o *ListRestoreHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list restore history params
func (o *ListRestoreHistoryParams) WithBody(body interface{}) *ListRestoreHistoryParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list restore history params
func (o *ListRestoreHistoryParams) SetBody(body interface{}) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListRestoreHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

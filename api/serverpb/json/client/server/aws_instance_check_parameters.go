// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewAWSInstanceCheckParams creates a new AWSInstanceCheckParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAWSInstanceCheckParams() *AWSInstanceCheckParams {
	return &AWSInstanceCheckParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAWSInstanceCheckParamsWithTimeout creates a new AWSInstanceCheckParams object
// with the ability to set a timeout on a request.
func NewAWSInstanceCheckParamsWithTimeout(timeout time.Duration) *AWSInstanceCheckParams {
	return &AWSInstanceCheckParams{
		timeout: timeout,
	}
}

// NewAWSInstanceCheckParamsWithContext creates a new AWSInstanceCheckParams object
// with the ability to set a context for a request.
func NewAWSInstanceCheckParamsWithContext(ctx context.Context) *AWSInstanceCheckParams {
	return &AWSInstanceCheckParams{
		Context: ctx,
	}
}

// NewAWSInstanceCheckParamsWithHTTPClient creates a new AWSInstanceCheckParams object
// with the ability to set a custom HTTPClient for a request.
func NewAWSInstanceCheckParamsWithHTTPClient(client *http.Client) *AWSInstanceCheckParams {
	return &AWSInstanceCheckParams{
		HTTPClient: client,
	}
}

/* AWSInstanceCheckParams contains all the parameters to send to the API endpoint
   for the AWS instance check operation.

   Typically these are written to a http.Request.
*/
type AWSInstanceCheckParams struct {
	// Body.
	Body AWSInstanceCheckBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the AWS instance check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AWSInstanceCheckParams) WithDefaults() *AWSInstanceCheckParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the AWS instance check params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AWSInstanceCheckParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the AWS instance check params
func (o *AWSInstanceCheckParams) WithTimeout(timeout time.Duration) *AWSInstanceCheckParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the AWS instance check params
func (o *AWSInstanceCheckParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the AWS instance check params
func (o *AWSInstanceCheckParams) WithContext(ctx context.Context) *AWSInstanceCheckParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the AWS instance check params
func (o *AWSInstanceCheckParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the AWS instance check params
func (o *AWSInstanceCheckParams) WithHTTPClient(client *http.Client) *AWSInstanceCheckParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the AWS instance check params
func (o *AWSInstanceCheckParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the AWS instance check params
func (o *AWSInstanceCheckParams) WithBody(body AWSInstanceCheckBody) *AWSInstanceCheckParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the AWS instance check params
func (o *AWSInstanceCheckParams) SetBody(body AWSInstanceCheckBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AWSInstanceCheckParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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

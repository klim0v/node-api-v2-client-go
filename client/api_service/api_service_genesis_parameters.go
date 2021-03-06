// Code generated by go-swagger; DO NOT EDIT.

package api_service

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

// NewAPIServiceGenesisParams creates a new APIServiceGenesisParams object
// with the default values initialized.
func NewAPIServiceGenesisParams() *APIServiceGenesisParams {

	return &APIServiceGenesisParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceGenesisParamsWithTimeout creates a new APIServiceGenesisParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceGenesisParamsWithTimeout(timeout time.Duration) *APIServiceGenesisParams {

	return &APIServiceGenesisParams{

		timeout: timeout,
	}
}

// NewAPIServiceGenesisParamsWithContext creates a new APIServiceGenesisParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceGenesisParamsWithContext(ctx context.Context) *APIServiceGenesisParams {

	return &APIServiceGenesisParams{

		Context: ctx,
	}
}

// NewAPIServiceGenesisParamsWithHTTPClient creates a new APIServiceGenesisParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceGenesisParamsWithHTTPClient(client *http.Client) *APIServiceGenesisParams {

	return &APIServiceGenesisParams{
		HTTPClient: client,
	}
}

/*APIServiceGenesisParams contains all the parameters to send to the API endpoint
for the Api service genesis operation typically these are written to a http.Request
*/
type APIServiceGenesisParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service genesis params
func (o *APIServiceGenesisParams) WithTimeout(timeout time.Duration) *APIServiceGenesisParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service genesis params
func (o *APIServiceGenesisParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service genesis params
func (o *APIServiceGenesisParams) WithContext(ctx context.Context) *APIServiceGenesisParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service genesis params
func (o *APIServiceGenesisParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service genesis params
func (o *APIServiceGenesisParams) WithHTTPClient(client *http.Client) *APIServiceGenesisParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service genesis params
func (o *APIServiceGenesisParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceGenesisParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

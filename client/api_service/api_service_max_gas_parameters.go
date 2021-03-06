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

// NewAPIServiceMaxGasParams creates a new APIServiceMaxGasParams object
// with the default values initialized.
func NewAPIServiceMaxGasParams() *APIServiceMaxGasParams {
	var ()
	return &APIServiceMaxGasParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceMaxGasParamsWithTimeout creates a new APIServiceMaxGasParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceMaxGasParamsWithTimeout(timeout time.Duration) *APIServiceMaxGasParams {
	var ()
	return &APIServiceMaxGasParams{

		timeout: timeout,
	}
}

// NewAPIServiceMaxGasParamsWithContext creates a new APIServiceMaxGasParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceMaxGasParamsWithContext(ctx context.Context) *APIServiceMaxGasParams {
	var ()
	return &APIServiceMaxGasParams{

		Context: ctx,
	}
}

// NewAPIServiceMaxGasParamsWithHTTPClient creates a new APIServiceMaxGasParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceMaxGasParamsWithHTTPClient(client *http.Client) *APIServiceMaxGasParams {
	var ()
	return &APIServiceMaxGasParams{
		HTTPClient: client,
	}
}

/*APIServiceMaxGasParams contains all the parameters to send to the API endpoint
for the Api service max gas operation typically these are written to a http.Request
*/
type APIServiceMaxGasParams struct {

	/*Height*/
	Height *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service max gas params
func (o *APIServiceMaxGasParams) WithTimeout(timeout time.Duration) *APIServiceMaxGasParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service max gas params
func (o *APIServiceMaxGasParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service max gas params
func (o *APIServiceMaxGasParams) WithContext(ctx context.Context) *APIServiceMaxGasParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service max gas params
func (o *APIServiceMaxGasParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service max gas params
func (o *APIServiceMaxGasParams) WithHTTPClient(client *http.Client) *APIServiceMaxGasParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service max gas params
func (o *APIServiceMaxGasParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the Api service max gas params
func (o *APIServiceMaxGasParams) WithHeight(height *string) *APIServiceMaxGasParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service max gas params
func (o *APIServiceMaxGasParams) SetHeight(height *string) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceMaxGasParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Height != nil {

		// query param height
		var qrHeight string
		if o.Height != nil {
			qrHeight = *o.Height
		}
		qHeight := qrHeight
		if qHeight != "" {
			if err := r.SetQueryParam("height", qHeight); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

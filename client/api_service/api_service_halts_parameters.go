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

// NewAPIServiceHaltsParams creates a new APIServiceHaltsParams object
// with the default values initialized.
func NewAPIServiceHaltsParams() *APIServiceHaltsParams {
	var ()
	return &APIServiceHaltsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceHaltsParamsWithTimeout creates a new APIServiceHaltsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceHaltsParamsWithTimeout(timeout time.Duration) *APIServiceHaltsParams {
	var ()
	return &APIServiceHaltsParams{

		timeout: timeout,
	}
}

// NewAPIServiceHaltsParamsWithContext creates a new APIServiceHaltsParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceHaltsParamsWithContext(ctx context.Context) *APIServiceHaltsParams {
	var ()
	return &APIServiceHaltsParams{

		Context: ctx,
	}
}

// NewAPIServiceHaltsParamsWithHTTPClient creates a new APIServiceHaltsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceHaltsParamsWithHTTPClient(client *http.Client) *APIServiceHaltsParams {
	var ()
	return &APIServiceHaltsParams{
		HTTPClient: client,
	}
}

/*APIServiceHaltsParams contains all the parameters to send to the API endpoint
for the Api service halts operation typically these are written to a http.Request
*/
type APIServiceHaltsParams struct {

	/*Height*/
	Height *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service halts params
func (o *APIServiceHaltsParams) WithTimeout(timeout time.Duration) *APIServiceHaltsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service halts params
func (o *APIServiceHaltsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service halts params
func (o *APIServiceHaltsParams) WithContext(ctx context.Context) *APIServiceHaltsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service halts params
func (o *APIServiceHaltsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service halts params
func (o *APIServiceHaltsParams) WithHTTPClient(client *http.Client) *APIServiceHaltsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service halts params
func (o *APIServiceHaltsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHeight adds the height to the Api service halts params
func (o *APIServiceHaltsParams) WithHeight(height *string) *APIServiceHaltsParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service halts params
func (o *APIServiceHaltsParams) SetHeight(height *string) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceHaltsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

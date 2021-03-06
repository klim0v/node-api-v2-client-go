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

// NewAPIServiceTransactionParams creates a new APIServiceTransactionParams object
// with the default values initialized.
func NewAPIServiceTransactionParams() *APIServiceTransactionParams {
	var ()
	return &APIServiceTransactionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceTransactionParamsWithTimeout creates a new APIServiceTransactionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceTransactionParamsWithTimeout(timeout time.Duration) *APIServiceTransactionParams {
	var ()
	return &APIServiceTransactionParams{

		timeout: timeout,
	}
}

// NewAPIServiceTransactionParamsWithContext creates a new APIServiceTransactionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceTransactionParamsWithContext(ctx context.Context) *APIServiceTransactionParams {
	var ()
	return &APIServiceTransactionParams{

		Context: ctx,
	}
}

// NewAPIServiceTransactionParamsWithHTTPClient creates a new APIServiceTransactionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceTransactionParamsWithHTTPClient(client *http.Client) *APIServiceTransactionParams {
	var ()
	return &APIServiceTransactionParams{
		HTTPClient: client,
	}
}

/*APIServiceTransactionParams contains all the parameters to send to the API endpoint
for the Api service transaction operation typically these are written to a http.Request
*/
type APIServiceTransactionParams struct {

	/*Hash*/
	Hash string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service transaction params
func (o *APIServiceTransactionParams) WithTimeout(timeout time.Duration) *APIServiceTransactionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service transaction params
func (o *APIServiceTransactionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service transaction params
func (o *APIServiceTransactionParams) WithContext(ctx context.Context) *APIServiceTransactionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service transaction params
func (o *APIServiceTransactionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service transaction params
func (o *APIServiceTransactionParams) WithHTTPClient(client *http.Client) *APIServiceTransactionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service transaction params
func (o *APIServiceTransactionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithHash adds the hash to the Api service transaction params
func (o *APIServiceTransactionParams) WithHash(hash string) *APIServiceTransactionParams {
	o.SetHash(hash)
	return o
}

// SetHash adds the hash to the Api service transaction params
func (o *APIServiceTransactionParams) SetHash(hash string) {
	o.Hash = hash
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceTransactionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param hash
	if err := r.SetPathParam("hash", o.Hash); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

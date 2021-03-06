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

// NewAPIServiceFrozenParams creates a new APIServiceFrozenParams object
// with the default values initialized.
func NewAPIServiceFrozenParams() *APIServiceFrozenParams {
	var ()
	return &APIServiceFrozenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceFrozenParamsWithTimeout creates a new APIServiceFrozenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceFrozenParamsWithTimeout(timeout time.Duration) *APIServiceFrozenParams {
	var ()
	return &APIServiceFrozenParams{

		timeout: timeout,
	}
}

// NewAPIServiceFrozenParamsWithContext creates a new APIServiceFrozenParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceFrozenParamsWithContext(ctx context.Context) *APIServiceFrozenParams {
	var ()
	return &APIServiceFrozenParams{

		Context: ctx,
	}
}

// NewAPIServiceFrozenParamsWithHTTPClient creates a new APIServiceFrozenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceFrozenParamsWithHTTPClient(client *http.Client) *APIServiceFrozenParams {
	var ()
	return &APIServiceFrozenParams{
		HTTPClient: client,
	}
}

/*APIServiceFrozenParams contains all the parameters to send to the API endpoint
for the Api service frozen operation typically these are written to a http.Request
*/
type APIServiceFrozenParams struct {

	/*Address*/
	Address *string
	/*Coin*/
	Coin *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service frozen params
func (o *APIServiceFrozenParams) WithTimeout(timeout time.Duration) *APIServiceFrozenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service frozen params
func (o *APIServiceFrozenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service frozen params
func (o *APIServiceFrozenParams) WithContext(ctx context.Context) *APIServiceFrozenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service frozen params
func (o *APIServiceFrozenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service frozen params
func (o *APIServiceFrozenParams) WithHTTPClient(client *http.Client) *APIServiceFrozenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service frozen params
func (o *APIServiceFrozenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddress adds the address to the Api service frozen params
func (o *APIServiceFrozenParams) WithAddress(address *string) *APIServiceFrozenParams {
	o.SetAddress(address)
	return o
}

// SetAddress adds the address to the Api service frozen params
func (o *APIServiceFrozenParams) SetAddress(address *string) {
	o.Address = address
}

// WithCoin adds the coin to the Api service frozen params
func (o *APIServiceFrozenParams) WithCoin(coin *string) *APIServiceFrozenParams {
	o.SetCoin(coin)
	return o
}

// SetCoin adds the coin to the Api service frozen params
func (o *APIServiceFrozenParams) SetCoin(coin *string) {
	o.Coin = coin
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceFrozenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Address != nil {

		// query param address
		var qrAddress string
		if o.Address != nil {
			qrAddress = *o.Address
		}
		qAddress := qrAddress
		if qAddress != "" {
			if err := r.SetQueryParam("address", qAddress); err != nil {
				return err
			}
		}

	}

	if o.Coin != nil {

		// query param coin
		var qrCoin string
		if o.Coin != nil {
			qrCoin = *o.Coin
		}
		qCoin := qrCoin
		if qCoin != "" {
			if err := r.SetQueryParam("coin", qCoin); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/swag"
)

// NewAPIServiceAddressesParams creates a new APIServiceAddressesParams object
// with the default values initialized.
func NewAPIServiceAddressesParams() *APIServiceAddressesParams {
	var ()
	return &APIServiceAddressesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAPIServiceAddressesParamsWithTimeout creates a new APIServiceAddressesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAPIServiceAddressesParamsWithTimeout(timeout time.Duration) *APIServiceAddressesParams {
	var ()
	return &APIServiceAddressesParams{

		timeout: timeout,
	}
}

// NewAPIServiceAddressesParamsWithContext creates a new APIServiceAddressesParams object
// with the default values initialized, and the ability to set a context for a request
func NewAPIServiceAddressesParamsWithContext(ctx context.Context) *APIServiceAddressesParams {
	var ()
	return &APIServiceAddressesParams{

		Context: ctx,
	}
}

// NewAPIServiceAddressesParamsWithHTTPClient creates a new APIServiceAddressesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAPIServiceAddressesParamsWithHTTPClient(client *http.Client) *APIServiceAddressesParams {
	var ()
	return &APIServiceAddressesParams{
		HTTPClient: client,
	}
}

/*APIServiceAddressesParams contains all the parameters to send to the API endpoint
for the Api service addresses operation typically these are written to a http.Request
*/
type APIServiceAddressesParams struct {

	/*Addresses*/
	Addresses []string
	/*Delegated*/
	Delegated *bool
	/*Height*/
	Height *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the Api service addresses params
func (o *APIServiceAddressesParams) WithTimeout(timeout time.Duration) *APIServiceAddressesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the Api service addresses params
func (o *APIServiceAddressesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the Api service addresses params
func (o *APIServiceAddressesParams) WithContext(ctx context.Context) *APIServiceAddressesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the Api service addresses params
func (o *APIServiceAddressesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the Api service addresses params
func (o *APIServiceAddressesParams) WithHTTPClient(client *http.Client) *APIServiceAddressesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the Api service addresses params
func (o *APIServiceAddressesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddresses adds the addresses to the Api service addresses params
func (o *APIServiceAddressesParams) WithAddresses(addresses []string) *APIServiceAddressesParams {
	o.SetAddresses(addresses)
	return o
}

// SetAddresses adds the addresses to the Api service addresses params
func (o *APIServiceAddressesParams) SetAddresses(addresses []string) {
	o.Addresses = addresses
}

// WithDelegated adds the delegated to the Api service addresses params
func (o *APIServiceAddressesParams) WithDelegated(delegated *bool) *APIServiceAddressesParams {
	o.SetDelegated(delegated)
	return o
}

// SetDelegated adds the delegated to the Api service addresses params
func (o *APIServiceAddressesParams) SetDelegated(delegated *bool) {
	o.Delegated = delegated
}

// WithHeight adds the height to the Api service addresses params
func (o *APIServiceAddressesParams) WithHeight(height *string) *APIServiceAddressesParams {
	o.SetHeight(height)
	return o
}

// SetHeight adds the height to the Api service addresses params
func (o *APIServiceAddressesParams) SetHeight(height *string) {
	o.Height = height
}

// WriteToRequest writes these params to a swagger request
func (o *APIServiceAddressesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAddresses := o.Addresses

	joinedAddresses := swag.JoinByFormat(valuesAddresses, "multi")
	// query array param addresses
	if err := r.SetQueryParam("addresses", joinedAddresses...); err != nil {
		return err
	}

	if o.Delegated != nil {

		// query param delegated
		var qrDelegated bool
		if o.Delegated != nil {
			qrDelegated = *o.Delegated
		}
		qDelegated := swag.FormatBool(qrDelegated)
		if qDelegated != "" {
			if err := r.SetQueryParam("delegated", qDelegated); err != nil {
				return err
			}
		}

	}

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

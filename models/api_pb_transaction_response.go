// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbTransactionResponse api pb transaction response
//
// swagger:model api_pbTransactionResponse
type APIPbTransactionResponse struct {

	// code
	Code string `json:"code,omitempty"`

	// data
	Data interface{} `json:"data,omitempty"`

	// from
	From string `json:"from,omitempty"`

	// gas
	Gas string `json:"gas,omitempty"`

	// gas coin
	GasCoin string `json:"gas_coin,omitempty"`

	// gas price
	GasPrice string `json:"gas_price,omitempty"`

	// hash
	Hash string `json:"hash,omitempty"`

	// height
	Height string `json:"height,omitempty"`

	// index
	Index string `json:"index,omitempty"`

	// log
	Log string `json:"log,omitempty"`

	// nonce
	Nonce string `json:"nonce,omitempty"`

	// payload
	// Format: byte
	Payload strfmt.Base64 `json:"payload,omitempty"`

	// raw tx
	RawTx string `json:"raw_tx,omitempty"`

	// tags
	Tags map[string]string `json:"tags,omitempty"`

	// type
	Type string `json:"type,omitempty"`
}

// Validate validates this api pb transaction response
func (m *APIPbTransactionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPbTransactionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbTransactionResponse) UnmarshalBinary(b []byte) error {
	var res APIPbTransactionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

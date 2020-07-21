// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbCandidateResponseStake api pb candidate response stake
//
// swagger:model api_pbCandidateResponseStake
type APIPbCandidateResponseStake struct {

	// bip value
	BipValue string `json:"bip_value,omitempty"`

	// coin
	Coin *APIPbCoin `json:"coin,omitempty"`

	// owner
	Owner string `json:"owner,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this api pb candidate response stake
func (m *APIPbCandidateResponseStake) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCoin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPbCandidateResponseStake) validateCoin(formats strfmt.Registry) error {

	if swag.IsZero(m.Coin) { // not required
		return nil
	}

	if m.Coin != nil {
		if err := m.Coin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("coin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPbCandidateResponseStake) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbCandidateResponseStake) UnmarshalBinary(b []byte) error {
	var res APIPbCandidateResponseStake
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbHaltsResponse api pb halts response
//
// swagger:model api_pbHaltsResponse
type APIPbHaltsResponse struct {

	// votes
	Votes []string `json:"votes"`
}

// Validate validates this api pb halts response
func (m *APIPbHaltsResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIPbHaltsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbHaltsResponse) UnmarshalBinary(b []byte) error {
	var res APIPbHaltsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIPbValidatorsResponse api pb validators response
//
// swagger:model api_pbValidatorsResponse
type APIPbValidatorsResponse struct {

	// validators
	Validators []*APIPbValidatorsResponseResult `json:"validators"`
}

// Validate validates this api pb validators response
func (m *APIPbValidatorsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateValidators(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPbValidatorsResponse) validateValidators(formats strfmt.Registry) error {

	if swag.IsZero(m.Validators) { // not required
		return nil
	}

	for i := 0; i < len(m.Validators); i++ {
		if swag.IsZero(m.Validators[i]) { // not required
			continue
		}

		if m.Validators[i] != nil {
			if err := m.Validators[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("validators" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPbValidatorsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbValidatorsResponse) UnmarshalBinary(b []byte) error {
	var res APIPbValidatorsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

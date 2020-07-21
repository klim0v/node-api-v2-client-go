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

// APIPbBlockResponseEvidence api pb block response evidence
//
// swagger:model api_pbBlockResponseEvidence
type APIPbBlockResponseEvidence struct {

	// evidence
	Evidence []*BlockResponseEvidenceEvidence `json:"evidence"`
}

// Validate validates this api pb block response evidence
func (m *APIPbBlockResponseEvidence) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEvidence(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIPbBlockResponseEvidence) validateEvidence(formats strfmt.Registry) error {

	if swag.IsZero(m.Evidence) { // not required
		return nil
	}

	for i := 0; i < len(m.Evidence); i++ {
		if swag.IsZero(m.Evidence[i]) { // not required
			continue
		}

		if m.Evidence[i] != nil {
			if err := m.Evidence[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("evidence" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIPbBlockResponseEvidence) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIPbBlockResponseEvidence) UnmarshalBinary(b []byte) error {
	var res APIPbBlockResponseEvidence
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// BlockRequestField block request field
//
// swagger:model BlockRequestField
type BlockRequestField string

const (

	// BlockRequestFieldTransactions captures enum value "transactions"
	BlockRequestFieldTransactions BlockRequestField = "transactions"

	// BlockRequestFieldMissed captures enum value "missed"
	BlockRequestFieldMissed BlockRequestField = "missed"

	// BlockRequestFieldBlockReward captures enum value "block_reward"
	BlockRequestFieldBlockReward BlockRequestField = "block_reward"

	// BlockRequestFieldSize captures enum value "size"
	BlockRequestFieldSize BlockRequestField = "size"

	// BlockRequestFieldProposer captures enum value "proposer"
	BlockRequestFieldProposer BlockRequestField = "proposer"

	// BlockRequestFieldValidators captures enum value "validators"
	BlockRequestFieldValidators BlockRequestField = "validators"

	// BlockRequestFieldEvidence captures enum value "evidence"
	BlockRequestFieldEvidence BlockRequestField = "evidence"
)

// for schema
var blockRequestFieldEnum []interface{}

func init() {
	var res []BlockRequestField
	if err := json.Unmarshal([]byte(`["transactions","missed","block_reward","size","proposer","validators","evidence"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		blockRequestFieldEnum = append(blockRequestFieldEnum, v)
	}
}

func (m BlockRequestField) validateBlockRequestFieldEnum(path, location string, value BlockRequestField) error {
	if err := validate.EnumCase(path, location, value, blockRequestFieldEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this block request field
func (m BlockRequestField) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateBlockRequestFieldEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

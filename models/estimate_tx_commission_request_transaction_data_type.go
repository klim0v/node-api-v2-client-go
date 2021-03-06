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

// EstimateTxCommissionRequestTransactionDataType estimate tx commission request transaction data type
//
// swagger:model EstimateTxCommissionRequestTransactionDataType
type EstimateTxCommissionRequestTransactionDataType string

const (

	// EstimateTxCommissionRequestTransactionDataTypeNr captures enum value "_"
	EstimateTxCommissionRequestTransactionDataTypeNr EstimateTxCommissionRequestTransactionDataType = "_"

	// EstimateTxCommissionRequestTransactionDataTypeSend captures enum value "Send"
	EstimateTxCommissionRequestTransactionDataTypeSend EstimateTxCommissionRequestTransactionDataType = "Send"

	// EstimateTxCommissionRequestTransactionDataTypeSellCoin captures enum value "SellCoin"
	EstimateTxCommissionRequestTransactionDataTypeSellCoin EstimateTxCommissionRequestTransactionDataType = "SellCoin"

	// EstimateTxCommissionRequestTransactionDataTypeSellAllCoin captures enum value "SellAllCoin"
	EstimateTxCommissionRequestTransactionDataTypeSellAllCoin EstimateTxCommissionRequestTransactionDataType = "SellAllCoin"

	// EstimateTxCommissionRequestTransactionDataTypeBuyCoin captures enum value "BuyCoin"
	EstimateTxCommissionRequestTransactionDataTypeBuyCoin EstimateTxCommissionRequestTransactionDataType = "BuyCoin"

	// EstimateTxCommissionRequestTransactionDataTypeCreateCoin captures enum value "CreateCoin"
	EstimateTxCommissionRequestTransactionDataTypeCreateCoin EstimateTxCommissionRequestTransactionDataType = "CreateCoin"

	// EstimateTxCommissionRequestTransactionDataTypeDeclareCandidacy captures enum value "DeclareCandidacy"
	EstimateTxCommissionRequestTransactionDataTypeDeclareCandidacy EstimateTxCommissionRequestTransactionDataType = "DeclareCandidacy"

	// EstimateTxCommissionRequestTransactionDataTypeDelegate captures enum value "Delegate"
	EstimateTxCommissionRequestTransactionDataTypeDelegate EstimateTxCommissionRequestTransactionDataType = "Delegate"

	// EstimateTxCommissionRequestTransactionDataTypeUnbond captures enum value "Unbond"
	EstimateTxCommissionRequestTransactionDataTypeUnbond EstimateTxCommissionRequestTransactionDataType = "Unbond"

	// EstimateTxCommissionRequestTransactionDataTypeRedeemCheck captures enum value "RedeemCheck"
	EstimateTxCommissionRequestTransactionDataTypeRedeemCheck EstimateTxCommissionRequestTransactionDataType = "RedeemCheck"

	// EstimateTxCommissionRequestTransactionDataTypeSetCandidateOnline captures enum value "SetCandidateOnline"
	EstimateTxCommissionRequestTransactionDataTypeSetCandidateOnline EstimateTxCommissionRequestTransactionDataType = "SetCandidateOnline"

	// EstimateTxCommissionRequestTransactionDataTypeSetCandidateOffline captures enum value "SetCandidateOffline"
	EstimateTxCommissionRequestTransactionDataTypeSetCandidateOffline EstimateTxCommissionRequestTransactionDataType = "SetCandidateOffline"

	// EstimateTxCommissionRequestTransactionDataTypeCreateMultisig captures enum value "CreateMultisig"
	EstimateTxCommissionRequestTransactionDataTypeCreateMultisig EstimateTxCommissionRequestTransactionDataType = "CreateMultisig"

	// EstimateTxCommissionRequestTransactionDataTypeMultisend captures enum value "Multisend"
	EstimateTxCommissionRequestTransactionDataTypeMultisend EstimateTxCommissionRequestTransactionDataType = "Multisend"

	// EstimateTxCommissionRequestTransactionDataTypeEditCandidate captures enum value "EditCandidate"
	EstimateTxCommissionRequestTransactionDataTypeEditCandidate EstimateTxCommissionRequestTransactionDataType = "EditCandidate"

	// EstimateTxCommissionRequestTransactionDataTypeRecreateCoin captures enum value "RecreateCoin"
	EstimateTxCommissionRequestTransactionDataTypeRecreateCoin EstimateTxCommissionRequestTransactionDataType = "RecreateCoin"

	// EstimateTxCommissionRequestTransactionDataTypeChangeOwner captures enum value "ChangeOwner"
	EstimateTxCommissionRequestTransactionDataTypeChangeOwner EstimateTxCommissionRequestTransactionDataType = "ChangeOwner"
)

// for schema
var estimateTxCommissionRequestTransactionDataTypeEnum []interface{}

func init() {
	var res []EstimateTxCommissionRequestTransactionDataType
	if err := json.Unmarshal([]byte(`["_","Send","SellCoin","SellAllCoin","BuyCoin","CreateCoin","DeclareCandidacy","Delegate","Unbond","RedeemCheck","SetCandidateOnline","SetCandidateOffline","CreateMultisig","Multisend","EditCandidate","RecreateCoin","ChangeOwner"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		estimateTxCommissionRequestTransactionDataTypeEnum = append(estimateTxCommissionRequestTransactionDataTypeEnum, v)
	}
}

func (m EstimateTxCommissionRequestTransactionDataType) validateEstimateTxCommissionRequestTransactionDataTypeEnum(path, location string, value EstimateTxCommissionRequestTransactionDataType) error {
	if err := validate.EnumCase(path, location, value, estimateTxCommissionRequestTransactionDataTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this estimate tx commission request transaction data type
func (m EstimateTxCommissionRequestTransactionDataType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateEstimateTxCommissionRequestTransactionDataTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

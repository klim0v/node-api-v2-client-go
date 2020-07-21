// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NodeInfoProtocolVersion node info protocol version
//
// swagger:model NodeInfoProtocolVersion
type NodeInfoProtocolVersion struct {

	// app
	App string `json:"app,omitempty"`

	// block
	Block string `json:"block,omitempty"`

	// p2p
	P2p string `json:"p2p,omitempty"`
}

// Validate validates this node info protocol version
func (m *NodeInfoProtocolVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NodeInfoProtocolVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NodeInfoProtocolVersion) UnmarshalBinary(b []byte) error {
	var res NodeInfoProtocolVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseVersion Version response
//
// swagger:model Response.Version
type ResponseVersion struct {

	// branch
	Branch string `json:"branch,omitempty"`

	// build time
	BuildTime string `json:"build_time,omitempty"`

	// commit
	Commit string `json:"commit,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this response version
func (m *ResponseVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseVersion) UnmarshalBinary(b []byte) error {
	var res ResponseVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

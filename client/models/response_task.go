// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ResponseTask Task Response
//
// swagger:model Response.Task
type ResponseTask struct {

	// id
	ID string `json:"id,omitempty"`

	// status
	Status string `json:"status,omitempty"`
}

// Validate validates this response task
func (m *ResponseTask) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ResponseTask) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResponseTask) UnmarshalBinary(b []byte) error {
	var res ResponseTask
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

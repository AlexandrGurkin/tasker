// Code generated by go-swagger; DO NOT EDIT.

package exec_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewGetExecuteTaskAgentParams creates a new GetExecuteTaskAgentParams object
// no default values defined in spec.
func NewGetExecuteTaskAgentParams() GetExecuteTaskAgentParams {

	return GetExecuteTaskAgentParams{}
}

// GetExecuteTaskAgentParams contains all the bound params for the get execute task agent operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetExecuteTaskAgent
type GetExecuteTaskAgentParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Required: true
	  In: path
	*/
	Agent string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetExecuteTaskAgentParams() beforehand.
func (o *GetExecuteTaskAgentParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rAgent, rhkAgent, _ := route.Params.GetOK("agent")
	if err := o.bindAgent(rAgent, rhkAgent, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindAgent binds and validates parameter Agent from path.
func (o *GetExecuteTaskAgentParams) bindAgent(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Agent = raw

	return nil
}

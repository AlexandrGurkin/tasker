// Code generated by go-swagger; DO NOT EDIT.

package set_result

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/AlexandrGurkin/tasker/client/models"
)

// NewPostSetResultIDParams creates a new PostSetResultIDParams object
// with the default values initialized.
func NewPostSetResultIDParams() *PostSetResultIDParams {
	var ()
	return &PostSetResultIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostSetResultIDParamsWithTimeout creates a new PostSetResultIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostSetResultIDParamsWithTimeout(timeout time.Duration) *PostSetResultIDParams {
	var ()
	return &PostSetResultIDParams{

		timeout: timeout,
	}
}

// NewPostSetResultIDParamsWithContext creates a new PostSetResultIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostSetResultIDParamsWithContext(ctx context.Context) *PostSetResultIDParams {
	var ()
	return &PostSetResultIDParams{

		Context: ctx,
	}
}

// NewPostSetResultIDParamsWithHTTPClient creates a new PostSetResultIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostSetResultIDParamsWithHTTPClient(client *http.Client) *PostSetResultIDParams {
	var ()
	return &PostSetResultIDParams{
		HTTPClient: client,
	}
}

/*PostSetResultIDParams contains all the parameters to send to the API endpoint
for the post set result ID operation typically these are written to a http.Request
*/
type PostSetResultIDParams struct {

	/*Body*/
	Body *models.ResponseTask
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post set result ID params
func (o *PostSetResultIDParams) WithTimeout(timeout time.Duration) *PostSetResultIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post set result ID params
func (o *PostSetResultIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post set result ID params
func (o *PostSetResultIDParams) WithContext(ctx context.Context) *PostSetResultIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post set result ID params
func (o *PostSetResultIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post set result ID params
func (o *PostSetResultIDParams) WithHTTPClient(client *http.Client) *PostSetResultIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post set result ID params
func (o *PostSetResultIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post set result ID params
func (o *PostSetResultIDParams) WithBody(body *models.ResponseTask) *PostSetResultIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post set result ID params
func (o *PostSetResultIDParams) SetBody(body *models.ResponseTask) {
	o.Body = body
}

// WithID adds the id to the post set result ID params
func (o *PostSetResultIDParams) WithID(id string) *PostSetResultIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post set result ID params
func (o *PostSetResultIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostSetResultIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
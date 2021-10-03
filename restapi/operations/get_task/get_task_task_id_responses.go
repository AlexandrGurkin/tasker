// Code generated by go-swagger; DO NOT EDIT.

package get_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/AlexandrGurkin/tasker/models"
)

// GetTaskTaskIDOKCode is the HTTP code returned for type GetTaskTaskIDOK
const GetTaskTaskIDOKCode int = 200

/*GetTaskTaskIDOK OK

swagger:response getTaskTaskIdOK
*/
type GetTaskTaskIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseTask `json:"body,omitempty"`
}

// NewGetTaskTaskIDOK creates GetTaskTaskIDOK with default headers values
func NewGetTaskTaskIDOK() *GetTaskTaskIDOK {

	return &GetTaskTaskIDOK{}
}

// WithPayload adds the payload to the get task task Id o k response
func (o *GetTaskTaskIDOK) WithPayload(payload *models.ResponseTask) *GetTaskTaskIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task task Id o k response
func (o *GetTaskTaskIDOK) SetPayload(payload *models.ResponseTask) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskTaskIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetTaskTaskIDBadRequestCode is the HTTP code returned for type GetTaskTaskIDBadRequest
const GetTaskTaskIDBadRequestCode int = 400

/*GetTaskTaskIDBadRequest Bad Request

swagger:response getTaskTaskIdBadRequest
*/
type GetTaskTaskIDBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ResponseError `json:"body,omitempty"`
}

// NewGetTaskTaskIDBadRequest creates GetTaskTaskIDBadRequest with default headers values
func NewGetTaskTaskIDBadRequest() *GetTaskTaskIDBadRequest {

	return &GetTaskTaskIDBadRequest{}
}

// WithPayload adds the payload to the get task task Id bad request response
func (o *GetTaskTaskIDBadRequest) WithPayload(payload *models.ResponseError) *GetTaskTaskIDBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get task task Id bad request response
func (o *GetTaskTaskIDBadRequest) SetPayload(payload *models.ResponseError) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTaskTaskIDBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
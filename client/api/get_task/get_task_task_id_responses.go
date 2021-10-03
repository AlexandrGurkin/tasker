// Code generated by go-swagger; DO NOT EDIT.

package get_task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/AlexandrGurkin/tasker/client/models"
)

// GetTaskTaskIDReader is a Reader for the GetTaskTaskID structure.
type GetTaskTaskIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskTaskIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTaskTaskIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTaskTaskIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTaskTaskIDOK creates a GetTaskTaskIDOK with default headers values
func NewGetTaskTaskIDOK() *GetTaskTaskIDOK {
	return &GetTaskTaskIDOK{}
}

/*GetTaskTaskIDOK handles this case with default header values.

OK
*/
type GetTaskTaskIDOK struct {
	Payload *models.ResponseTask
}

func (o *GetTaskTaskIDOK) Error() string {
	return fmt.Sprintf("[GET /task/{task_id}][%d] getTaskTaskIdOK  %+v", 200, o.Payload)
}

func (o *GetTaskTaskIDOK) GetPayload() *models.ResponseTask {
	return o.Payload
}

func (o *GetTaskTaskIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseTask)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskTaskIDBadRequest creates a GetTaskTaskIDBadRequest with default headers values
func NewGetTaskTaskIDBadRequest() *GetTaskTaskIDBadRequest {
	return &GetTaskTaskIDBadRequest{}
}

/*GetTaskTaskIDBadRequest handles this case with default header values.

Bad Request
*/
type GetTaskTaskIDBadRequest struct {
	Payload *models.ResponseError
}

func (o *GetTaskTaskIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /task/{task_id}][%d] getTaskTaskIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetTaskTaskIDBadRequest) GetPayload() *models.ResponseError {
	return o.Payload
}

func (o *GetTaskTaskIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResponseError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
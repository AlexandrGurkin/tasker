package handlers

import (
	"time"

	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/get_task"
	"github.com/go-openapi/runtime/middleware"
)

type GetTaskHandler struct {
	TaskManager *internal.TaskManager
}

func (h GetTaskHandler) Handle(params get_task.GetTaskTaskIDParams) middleware.Responder {
	if ch, ok := h.TaskManager.Tasks[params.TaskID]; ok {
		select {
		case <-time.After(time.Minute):
			return get_task.NewGetTaskTaskIDBadRequest().WithPayload(&models.ResponseError{
				Code:    10,
				Message: "Timeout",
			})
		case res := <-ch:
			return get_task.NewGetTaskTaskIDOK().WithPayload(&res)

		}
	} else {
		return get_task.NewGetTaskTaskIDBadRequest().WithPayload(&models.ResponseError{
			Code:    101,
			Message: "Unknown task id",
		})
	}
}

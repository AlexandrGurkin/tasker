package handlers

import (
	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/create_task"
	"github.com/go-openapi/runtime/middleware"
)

type CreateTaskHandler struct {
	TaskManager *internal.TaskManager
}

func (h CreateTaskHandler) Handle(params create_task.PostTaskParams) middleware.Responder {
	resChan := make(chan models.ResponseTask)
	h.TaskManager.Tasks[params.Body.ID] = resChan

	if queue, ok := h.TaskManager.VMQueue[params.Body.Agent]; ok {
		queue <- *params.Body
	} else {
		ch := make(chan models.RequestTask, 100)
		h.TaskManager.VMQueue[params.Body.Agent] = ch
		ch <- *params.Body
	}

	return create_task.NewPostTaskNoContent()
}

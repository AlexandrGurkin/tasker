package handlers

import (
	"time"

	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/exec_task"
	"github.com/go-openapi/runtime/middleware"
)

type GetToExecHandler struct {
	TaskManager *internal.TaskManager
}

func (h GetToExecHandler) Handle(params exec_task.GetExecuteTaskAgentParams) middleware.Responder {
	if ch, ok := h.TaskManager.VMQueue[params.Agent]; ok {
		select {
		case <-time.After(time.Minute):
			return exec_task.NewGetExecuteTaskAgentBadRequest().WithPayload(&models.ResponseError{
				Code:    12,
				Message: "Timeout",
			})
		case res := <-ch:
			return exec_task.NewGetExecuteTaskAgentOK().WithPayload(&res)

		}
	} else {
		return exec_task.NewGetExecuteTaskAgentBadRequest().WithPayload(&models.ResponseError{
			Code:    120,
			Message: "Unknown agent",
		})
	}
}

package handlers

import (
	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/restapi/operations/set_result"
	"github.com/go-openapi/runtime/middleware"
)

type SetResHandler struct {
	TaskManager *internal.TaskManager
}

func (h SetResHandler) Handle(params set_result.PostSetResultIDParams) middleware.Responder {
	if ch, ok := h.TaskManager.Tasks[params.ID]; ok {
		ch <- *params.Body
	}

	return set_result.NewPostSetResultIDNoContent()
}

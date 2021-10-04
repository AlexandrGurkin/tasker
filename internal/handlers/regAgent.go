package handlers

import (
	"time"

	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/reg_agent"
	"github.com/go-openapi/runtime/middleware"
)

type RegAgentHandler struct {
	TaskManager  *internal.TaskManager
	AgentManager *internal.AgentManager
}

func (h RegAgentHandler) Handle(params reg_agent.PostAgentParams) middleware.Responder {
	id := params.Body.ID
	if _, ok := h.TaskManager.VMQueue[id]; !ok {
		ch := make(chan models.RequestTask, 100)
		h.TaskManager.VMQueue[id] = ch
	}
	if wd, ok := h.AgentManager.Wdogs[id]; ok {
		wd.Reset()
	} else {
		h.AgentManager.Wdogs[params.Body.ID] = internal.NewWD(2*time.Second, func() {
			h.AgentManager.Set(id, []string{})
		})
	}

	h.AgentManager.Set(params.Body.ID, params.Body.Nets)

	return reg_agent.NewPostAgentNoContent()
}

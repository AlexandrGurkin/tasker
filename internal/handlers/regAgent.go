package handlers

import (
	"github.com/AlexandrGurkin/tasker/restapi/operations/reg_agent"
	"github.com/go-openapi/runtime/middleware"
)

type RegAgentHandler struct {
}

func (h RegAgentHandler) Handle(params reg_agent.PostAgentParams) middleware.Responder {

	return reg_agent.NewPostAgentNoContent()
}

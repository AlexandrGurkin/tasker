package handlers

import (
	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/address"
	"github.com/go-openapi/runtime/middleware"
)

type GetAgentHandler struct {
	AgentManager *internal.AgentManager
}

func (h GetAgentHandler) Handle(params address.GetAddressAgentParams) middleware.Responder {
	res, err := h.AgentManager.Get(params.Agent)
	if err != nil {
		address.NewGetAddressAgentBadRequest().WithPayload(&models.ResponseError{
			Code:    33,
			Message: err.Error(),
		})
	}

	return address.NewGetAddressAgentOK().WithPayload(&models.RequestAgent{
		ID:   params.Agent,
		Nets: res,
	})
}

package handlers

import (
	"github.com/AlexandrGurkin/tasker/internal/ver"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi/operations/version"
	"github.com/go-openapi/runtime/middleware"
)

type VersionHandler struct {
}

func (vh VersionHandler) Handle(_ version.GetVersionParams) middleware.Responder {
	return version.NewGetVersionOK().WithPayload(&models.ResponseVersion{
		Version:   ver.GetVersion(),
		Branch:    ver.GetBranch(),
		Commit:    ver.GetCommit(),
		BuildTime: ver.GetBuildTime(),
	})
}

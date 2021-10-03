// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/AlexandrGurkin/common/middlewares"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/AlexandrGurkin/tasker/restapi/operations"
	"github.com/AlexandrGurkin/tasker/restapi/operations/create_task"
	"github.com/AlexandrGurkin/tasker/restapi/operations/exec_task"
	"github.com/AlexandrGurkin/tasker/restapi/operations/get_task"
	"github.com/AlexandrGurkin/tasker/restapi/operations/reg_agent"
	"github.com/AlexandrGurkin/tasker/restapi/operations/set_result"
	"github.com/AlexandrGurkin/tasker/restapi/operations/version"
)

//go:generate swagger generate server --target ../../work --name TemplateForHTTPServer --spec ../api/swagger.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.TemplateForHTTPServerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.TemplateForHTTPServerAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	if api.ExecTaskGetExecuteTaskAgentHandler == nil {
		api.ExecTaskGetExecuteTaskAgentHandler = exec_task.GetExecuteTaskAgentHandlerFunc(func(params exec_task.GetExecuteTaskAgentParams) middleware.Responder {
			return middleware.NotImplemented("operation exec_task.GetExecuteTaskAgent has not yet been implemented")
		})
	}
	if api.GetTaskGetTaskTaskIDHandler == nil {
		api.GetTaskGetTaskTaskIDHandler = get_task.GetTaskTaskIDHandlerFunc(func(params get_task.GetTaskTaskIDParams) middleware.Responder {
			return middleware.NotImplemented("operation get_task.GetTaskTaskID has not yet been implemented")
		})
	}
	if api.VersionGetVersionHandler == nil {
		api.VersionGetVersionHandler = version.GetVersionHandlerFunc(func(params version.GetVersionParams) middleware.Responder {
			return middleware.NotImplemented("operation version.GetVersion has not yet been implemented")
		})
	}
	if api.RegAgentPostAgentHandler == nil {
		api.RegAgentPostAgentHandler = reg_agent.PostAgentHandlerFunc(func(params reg_agent.PostAgentParams) middleware.Responder {
			return middleware.NotImplemented("operation reg_agent.PostAgent has not yet been implemented")
		})
	}
	if api.SetResultPostSetResultIDHandler == nil {
		api.SetResultPostSetResultIDHandler = set_result.PostSetResultIDHandlerFunc(func(params set_result.PostSetResultIDParams) middleware.Responder {
			return middleware.NotImplemented("operation set_result.PostSetResultID has not yet been implemented")
		})
	}
	if api.CreateTaskPostTaskHandler == nil {
		api.CreateTaskPostTaskHandler = create_task.PostTaskHandlerFunc(func(params create_task.PostTaskParams) middleware.Responder {
			return middleware.NotImplemented("operation create_task.PostTask has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

func setupGlobalMiddleware(handler http.Handler) http.Handler {
	if mc.Pprof {
		handler = middlewares.PprofHandler(handler)
	}
	return middlewares.RequestLog(handler, mc)
}

var mc middlewares.MiddlewareConfig

func SetMiddlewareConfig(c middlewares.MiddlewareConfig) {
	mc = c
}

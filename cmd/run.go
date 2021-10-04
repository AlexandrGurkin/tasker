package cmd

import (
	"os"
	"time"

	"github.com/AlexandrGurkin/common/consts"
	"github.com/AlexandrGurkin/common/middlewares"
	"github.com/AlexandrGurkin/common/xlog"
	"github.com/AlexandrGurkin/common/xlog/xzerolog"
	"github.com/AlexandrGurkin/tasker/internal"
	"github.com/AlexandrGurkin/tasker/internal/handlers"
	"github.com/AlexandrGurkin/tasker/models"
	"github.com/AlexandrGurkin/tasker/restapi"
	"github.com/AlexandrGurkin/tasker/restapi/operations"
	"github.com/go-openapi/loads"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "run callback",
	Long:  `run callback`,
	Run:   runMain,
}

func init() {
	RootCmd.AddCommand(runCmd)
}

func runMain(cmd *cobra.Command, args []string) {
	logger := xzerolog.NewXZerolog(xlog.LoggerCfg{
		Level: "trace",
		Out:   os.Stdout,
	})

	var err error
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		logger.Fatal(err.Error())
	}

	api := operations.NewTemplateForHTTPServerAPI(swaggerSpec)
	server := restapi.NewServer(api)

	server.Host = "0.0.0.0"
	server.Port = 8077
	restapi.SetMiddlewareConfig(middlewares.MiddlewareConfig{Logger: logger, Pprof: true})
	api.Logger = func(s string, i ...interface{}) {
		logger.WithXFields(xlog.Fields{consts.FieldModule: "swagger_api_logger"}).
			Infof(s, i...)
	}

	tm := internal.TaskManager{
		Tasks:   make(map[string]chan models.ResponseTask),
		VMQueue: make(map[string]chan models.RequestTask),
	}

	am := internal.NewAgentManager()

	api.VersionGetVersionHandler = handlers.VersionHandler{}
	api.CreateTaskPostTaskHandler = handlers.CreateTaskHandler{&tm}
	api.GetTaskGetTaskTaskIDHandler = handlers.GetTaskHandler{&tm}
	api.RegAgentPostAgentHandler = handlers.RegAgentHandler{&tm, am}
	api.ExecTaskGetExecuteTaskAgentHandler = handlers.GetToExecHandler{&tm}
	api.SetResultPostSetResultIDHandler = handlers.SetResHandler{&tm}
	api.AddressGetAddressAgentHandler = handlers.GetAgentHandler{am}

	server.ConfigureAPI()
	server.KeepAlive = 10 * time.Second

	if err = server.Serve(); err != nil {
		logger.Fatal(err.Error())
	}
}

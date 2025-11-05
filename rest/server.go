package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/config"
	"github.com/mostafejur21/task_manager_backend/rest/handlers/tasks"
	"github.com/mostafejur21/task_manager_backend/rest/middlewares"
	"go.uber.org/zap"
)

type Server struct {
	cnf         *config.Config
	taskHandler *tasks.Handler
	loggers     *zap.SugaredLogger
}

func NewServer(
	cnf *config.Config,
	taskHandler *tasks.Handler,
	logger *zap.SugaredLogger,
) *Server {
	return &Server{
		cnf:         cnf,
		taskHandler: taskHandler,
		loggers:     logger,
	}
}

func (s *Server) Start() {
	manager := middlewares.NewManager()
	manager.Use(middlewares.Preflight, middlewares.Cors, middlewares.Logger)

	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	s.taskHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(s.cnf.Port)
	fmt.Println("Starting the server at port: ", addr)

	if err := http.ListenAndServe(addr, wrappedMux); err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
}

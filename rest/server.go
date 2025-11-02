package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/config"
	"github.com/mostafejur21/task_manager_backend/rest/handlers/tasks"
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
	mux := http.NewServeMux()

	s.taskHandler.RegisterRoutes(mux)

	addr := ":" + strconv.Itoa(s.cnf.Port)
	fmt.Println("Starting the server at port: ", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		fmt.Println("Error starting server: ", err)
		os.Exit(1)
	}
}

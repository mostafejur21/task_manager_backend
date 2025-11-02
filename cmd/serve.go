package cmd

import (
	"github.com/mostafejur21/task_manager_backend/config"
	"github.com/mostafejur21/task_manager_backend/infra/db"
	"github.com/mostafejur21/task_manager_backend/repo"
	"github.com/mostafejur21/task_manager_backend/rest"
	"github.com/mostafejur21/task_manager_backend/rest/handlers/tasks"
	"github.com/mostafejur21/task_manager_backend/task"
	"go.uber.org/zap"
)

func Serve() {
	// 1. load the config
	cnf := config.GetConfig()

	logger := zap.Must(zap.NewProduction()).Sugar()

	//2. load the db connection
	dbConn, err := db.NewDBConnection(cnf.DB)
	if err != nil {
		logger.Fatal(err)
	}

	//3. repo
	taskRepo := repo.NewTaskRepo(dbConn)

	// 4. domain
	taskSvc := task.NewService(taskRepo)

	//5. handler
	taskHandler := tasks.NewHandler(taskSvc)

	server := rest.NewServer(cnf, taskHandler, logger)

	server.Start()
}

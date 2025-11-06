package cmd

import (
	"github.com/mostafejur21/task_manager_backend/config"
	"github.com/mostafejur21/task_manager_backend/infra/db"
	"github.com/mostafejur21/task_manager_backend/repo"
	"github.com/mostafejur21/task_manager_backend/rest"
	taskHandlers "github.com/mostafejur21/task_manager_backend/rest/handlers/tasks"
	userHandlers "github.com/mostafejur21/task_manager_backend/rest/handlers/user"
	"github.com/mostafejur21/task_manager_backend/task"
	"github.com/mostafejur21/task_manager_backend/user"
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
	userRepo := repo.NewUserRepo(dbConn)

	// 4. domain
	taskSvc := task.NewService(taskRepo)
	userSvc := user.NewService(userRepo)

	//5. handler
	taskHandler := taskHandlers.NewHandler(taskSvc)
	userHandler := userHandlers.NewHandler(userSvc)

	server := rest.NewServer(cnf, taskHandler, userHandler, logger)

	server.Start()
}

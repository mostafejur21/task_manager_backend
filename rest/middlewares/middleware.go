package middlewares

import (
	"github.com/mostafejur21/task_manager_backend/config"
)

type Middlewares struct {
	cnf *config.Config
}

func NewMiddleware(cnf *config.Config) *Middlewares {
	return &Middlewares{
		cnf: cnf,
	}
}

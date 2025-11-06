package user

import (
	"github.com/mostafejur21/task_manager_backend/domain"
	userHandler "github.com/mostafejur21/task_manager_backend/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}

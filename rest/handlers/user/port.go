package user

import "github.com/mostafejur21/task_manager_backend/domain"

type Service interface {
	Create(domain.User) (*domain.User, error)
	Find(email, password string) (*domain.User, error)
}

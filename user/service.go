package user

import "github.com/mostafejur21/task_manager_backend/domain"

type service struct {
	userRepo UserRepo
}

func NewService(userRepo UserRepo) Service {
	return &service{userRepo: userRepo}
}

func (srv *service) Create(u domain.User) (*domain.User, error) {
	return srv.userRepo.Create(u)
}

func (srv *service) Find(email, password string) (*domain.User, error) {
	return srv.Find(email, password)
}

package task

import "github.com/mostafejur21/task_manager_backend/domain"

type service struct {
	taskRepo TaskRepo
}

func NewService(taskRepo TaskRepo) Service {
	return &service{taskRepo: taskRepo}
}


func (srv *service) Create(t domain.Task) (*domain.Task, error) {
	return srv.taskRepo.Create(t)
}

func (srv *service) Get(id int) (*domain.Task, error) {
	return srv.taskRepo.Get(id)
}

func (srv *service) List(page, limit int64) ([]*domain.Task, error) {
	return srv.taskRepo.List(page, limit)
}

func (srv *service) Delete(id int) error {
	return srv.taskRepo.Delete(id)
}
func (srv *service) Update(t domain.Task) (*domain.Task, error) {
	return srv.taskRepo.Update(t)
}

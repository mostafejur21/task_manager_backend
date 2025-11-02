package tasks

import "github.com/mostafejur21/task_manager_backend/domain"

type Service interface {
	Create(domain.Task) (*domain.Task, error)
	Get(id int) (*domain.Task, error)
	List(page, limit int64) ([]*domain.Task, error)
	Delete(id int) error
	Update(domain.Task) (*domain.Task, error)
	// Count() (int64, error)
}

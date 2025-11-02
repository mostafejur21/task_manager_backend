package task

import (
	"github.com/mostafejur21/task_manager_backend/domain"
	taskHandler "github.com/mostafejur21/task_manager_backend/rest/handlers/tasks"
)

type Service interface {
	taskHandler.Service
}

type TaskRepo interface {
	Create(domain.Task) (*domain.Task, error)
	Get(id int) (*domain.Task, error)
	GetByStatus(status string) ([]*domain.Task, error)
	List(page, limit int64) ([]*domain.Task, error)
	Delete(id int) error
	Update(domain.Task) (*domain.Task, error)
}

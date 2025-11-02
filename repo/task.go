package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mostafejur21/task_manager_backend/domain"
	"github.com/mostafejur21/task_manager_backend/task"
)

type TaskRepo interface {
	task.TaskRepo
}

type taskRepo struct {
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(t domain.Task) (*domain.Task, error) {
	query := `
		INSERT INTO tasks (
		title,
		description
	) VALUES (
		$1,
		$2
	)
	RETURNING id
	`

	row := r.db.QueryRow(query, t.Title, t.Description)
	if err := row.Scan(&t.ID); err != nil {
		fmt.Println("DB Insert error: ", err)
		return nil, err
	}
	return &t, nil
}

func (r *taskRepo) Get(id int) (*domain.Task, error) {
	return nil, nil
}

func (r *taskRepo) List(page, limit int64) ([]*domain.Task, error) {
	return nil, nil
}

func (r *taskRepo) Delete(id int) error {
	return nil
}
func (r *taskRepo) Update(domain.Task) (*domain.Task, error) {
	return nil, nil
}

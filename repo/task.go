package repo

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

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
	var task domain.Task
	query := `
	SELECT
		id,
		title,
		description,
		status,
		created_at
	FROM tasks
	WHERE id=$1
`
	if err := r.db.Get(&task, query, id); err != nil {
		if err == sql.ErrNoRows {
			// let handler return 404 on nil
			return nil, nil
		}
		return nil, err
	}
	return &task, nil
}

func (r *taskRepo) GetByStatus(status string, page, limit int64) ([]*domain.Task, error) {
	offset := (page - 1) * limit
	var task []*domain.Task
	query := `
	SELECT
		id,
		title,
		description,
		status,
		created_at
	FROM tasks
	WHERE status IS NOT NULL
		AND LOWER(TRIM(status)) = LOWER(TRIM($1))
	ORDER BY id ASC
	LIMIT $2 OFFSET $3
`
	if err := r.db.Select(&task, query, status, limit, offset); err != nil {
		return nil, err
	}

	return task, nil
}

func (r *taskRepo) List(page, limit int64) ([]*domain.Task, error) {
	offset := (page - 1) * limit
	var taskList []*domain.Task

	query := `
	SELECT
		id,
		title,
		description,
		status,
		created_at
	FROM tasks
	LIMIT $1 OFFSET $2
	`

	if err := r.db.Select(&taskList, query, limit, offset); err != nil {
		return nil, err
	}

	return taskList, nil
}

func (r *taskRepo) Delete(id int) error {
	query := `
	DELETE FROM tasks WHERE id=$1
`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *taskRepo) Update(t domain.Task) (*domain.Task, error) {
	setClauses := []string{}
	args := []interface{}{}
	i := 1

	if t.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf("title=$%d", i))
		args = append(args, t.Title)
		i++
	}

	if t.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description=$%d", i))
		args = append(args, t.Description)
		i++

	}
	if t.Status != nil {
		setClauses = append(setClauses, fmt.Sprintf("status=$%d", i))
		args = append(args, t.Status)
		i++
	}
	if len(setClauses) == 0 {
		return nil, errors.New("no fields to update")
	}

	query := fmt.Sprintf(
		"UPDATE tasks SET %s WHERE id=$%d", strings.Join(setClauses, ","), i,
	)
	args = append(args, t.ID)
	_, err := r.db.Exec(query, args...)

	return &t, err
}

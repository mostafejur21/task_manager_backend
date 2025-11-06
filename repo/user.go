package repo

import (
	"github.com/jmoiron/sqlx"
	"github.com/mostafejur21/task_manager_backend/domain"
	"github.com/mostafejur21/task_manager_backend/user"
)

type UserRepo interface {
	user.UserRepo
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{db: db}
}

func (r *userRepo) Create(domain.User) (*domain.User, error) {
	return nil, nil
}

func (r *userRepo) Find(email, password string) (*domain.User, error) {
	return nil, nil
}

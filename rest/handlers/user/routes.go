package user

import (
	"net/http"

	"github.com/mostafejur21/task_manager_backend/rest/middlewares"
)

func (h Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /register", manager.With(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
	mux.Handle("POST /login", manager.With(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))
}

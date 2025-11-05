package tasks

import (
	"net/http"

	"github.com/mostafejur21/task_manager_backend/rest/middlewares"
)

func (h Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	mux.Handle("POST /tasks", manager.With(http.HandlerFunc(h.CreateTask)))
	mux.Handle("GET /tasks/{id}", manager.With(http.HandlerFunc(h.GetTasksById)))
	mux.Handle("GET /tasks", manager.With(http.HandlerFunc(h.GetTasks)))
	mux.Handle("PATCH /tasks/{id}",manager.With(http.HandlerFunc( h.UpdateTask)))
	mux.Handle("DELETE /tasks/{id}",manager.With(http.HandlerFunc( h.DeleteTask)))
}

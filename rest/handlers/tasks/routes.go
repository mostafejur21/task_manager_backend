package tasks

import "net/http"

func (h Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /tasks", h.CreateTask)
}

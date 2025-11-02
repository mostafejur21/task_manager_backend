package tasks

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mostafejur21/task_manager_backend/domain"
	"github.com/mostafejur21/task_manager_backend/utils"
)

type TaskPayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req TaskPayload

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		fmt.Println(err)
		utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
	}

	createTask, err := h.srv.Create(domain.Task{
		Title: req.Title,
		Description: req.Description,
		Status: "in-progress",
	})

	if err != nil {
		fmt.Println(err)
		utils.WriteJsonError(w, http.StatusInternalServerError, "Internal Server error")
	}

	utils.JsonResponse(w, http.StatusCreated, createTask)
}

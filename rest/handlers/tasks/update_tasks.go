package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/domain"
	"github.com/mostafejur21/task_manager_backend/utils"
)

type TaskUpdatePayload struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.PathValue("id")

	tId, err := strconv.Atoi(taskId)
	if err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "invalid task id")
	}

	var req TaskUpdatePayload
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JsonResponse(w, http.StatusBadRequest, "invalid request body")
	}

	_, err = h.srv.Update(domain.Task{
		ID:          tId,
		Title:       req.Title,
		Description: req.Description,
		Status:      &req.Status,
	})

	if err != nil {
		_ = utils.WriteJsonError(w, http.StatusInternalServerError, "Internal server error")
		return
	}
	_ = utils.JsonResponse(w, http.StatusOK, "Successfully updated the task")
}

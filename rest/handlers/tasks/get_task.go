package tasks

import (
	"net/http"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/utils"
)

func (h *Handler) GetTasksById(w http.ResponseWriter, r *http.Request) {
	// get the id from url
	taskId := r.PathValue("id")

	// convert id into integer
	tId, err := strconv.Atoi(taskId)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	task, err := h.srv.Get(tId)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, "internal server error")
		return
	}

	if task == nil {
		utils.WriteJsonError(w, http.StatusNotFound, "task not found")
		return
	}

	utils.JsonResponse(w, http.StatusOK, task)
}

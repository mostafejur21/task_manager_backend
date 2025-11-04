package tasks

import (
	"net/http"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/utils"
)

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskId := r.PathValue("id")

	tId, err := strconv.Atoi(taskId)
	if err != nil {
		utils.WriteJsonError(w, http.StatusBadRequest, "invalid task id")
		return
	}

	err = h.srv.Delete(tId)
	if err != nil {
		utils.WriteJsonError(w, http.StatusInternalServerError, "Internal server error")
	}
	utils.JsonResponse(w, http.StatusOK, "task deleted successfully")
}

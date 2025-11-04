package tasks

import (
	"net/http"
	"strconv"

	"github.com/mostafejur21/task_manager_backend/utils"
)

func (h *Handler) GetTasks(w http.ResponseWriter, r *http.Request) {
	reqQuery := r.URL.Query()

	pageAsstr := reqQuery.Get("page")
	limitAsstr := reqQuery.Get("limit")
	status := reqQuery.Get("status")

	page, _ := strconv.Atoi(pageAsstr)
	limit, _ := strconv.Atoi(limitAsstr)

	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	if status != "" {
		taskList, err := h.srv.GetByStatus(status, int64(page), int64(limit))
		if err != nil {
			utils.WriteJsonError(w, http.StatusInternalServerError, "Internal server error")
		}
		utils.JsonResponse(w, http.StatusOK, taskList)
	} else {

		taskList, err := h.srv.List(int64(page), int64(limit))
		if err != nil {
			utils.WriteJsonError(w, http.StatusInternalServerError, "Internal server error")
		}

		utils.JsonResponse(w, http.StatusOK, taskList)
	}
}

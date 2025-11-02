package rest

import (
	"net/http"

	"github.com/mostafejur21/task_manager_backend/utils"
)

func (app *Server) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	utils.WriteJsonError(w, http.StatusInternalServerError, "the server encountered a problem")
}

func (app *Server) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	app.loggers.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")

	utils.WriteJsonError(w, http.StatusForbidden, "forbidden")
}

func (app *Server) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Warnf("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	utils.WriteJsonError(w, http.StatusBadRequest, err.Error())
}

func (app *Server) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Errorf("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	utils.WriteJsonError(w, http.StatusConflict, err.Error())
}

func (app *Server) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Warnf("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	utils.WriteJsonError(w, http.StatusNotFound, "not found")
}

func (app *Server) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Warnf("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	utils.WriteJsonError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *Server) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.loggers.Warnf("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)

	utils.WriteJsonError(w, http.StatusUnauthorized, "unauthorized")
}

func (app *Server) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	app.loggers.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	utils.WriteJsonError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}

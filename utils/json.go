package utils

import (
	"encoding/json"
	"net/http"
)

func ReadJson(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578 //1mb
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func writeJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Success bool   `json:"success"`
		Status  int    `json:"status"`
		Error   string `json:"error"`
	}

	return writeJson(w, status, &envelope{
		Success: false,
		Status:  status,
		Error:   message,
	})
}

func JsonResponse(w http.ResponseWriter, status int, data any) error {
	type envelope struct {
		Success bool `json:"success"`
		Status  int  `json:"status"`
		Data    any  `json:"data"`
	}
	return writeJson(w, status, &envelope{Success: true, Status: status, Data: data})

}

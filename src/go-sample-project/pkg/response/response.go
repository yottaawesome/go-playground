// Package response provides helper functions for writing JSON HTTP responses.
// It lives in pkg/ because it is generic enough to be imported by other
// projects—unlike the internal/ packages which are private to this module.
package response

import (
	"encoding/json"
	"net/http"
)

// ErrorBody is the standard shape for error responses.
type ErrorBody struct {
	Error string `json:"error"`
}

// JSON marshals v to JSON and writes it with the given HTTP status code.
// If marshalling fails it falls back to a 500 Internal Server Error.
func JSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(v); err != nil {
		// At this point headers are already sent, so we can only log.
		http.Error(w, `{"error":"internal server error"}`, http.StatusInternalServerError)
	}
}

// Error writes a JSON error response with the given status and message.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, ErrorBody{Error: message})
}

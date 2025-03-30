package handlers

import (
	"encoding/json"
	"net/http"
)

// Standard API response structure
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// writes a JSON response
func RespondWithJSON(w http.ResponseWriter, code int, payload any) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// writes an error response
func ErrorResponse(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, Response{
		Status:  "error",
		Message: message,
	})
}

// writes a success response
func SuccessResponse(w http.ResponseWriter, code int, message string, data any) {
	RespondWithJSON(w, code, Response{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

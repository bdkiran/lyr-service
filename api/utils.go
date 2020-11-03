package api

import (
	"encoding/json"
	"net/http"
)

type apiResponse struct {
	Message     string `json:"message"`
	Description string `json:"description"`
}

func sendResponse(message string, description string, statusCode int, w http.ResponseWriter) {
	responsePayload := apiResponse{
		Message:     message,
		Description: description,
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(responsePayload)
}

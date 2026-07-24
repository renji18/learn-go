package util

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type payloadResponse struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func SendJson(w http.ResponseWriter, statusCode int, message string, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	payload := payloadResponse{
		Message: message,
		Data:    data,
	}

	err := json.NewEncoder(w).Encode(payload)
	if err != nil {
		return fmt.Errorf("Error writing json")
	}

	return nil
}

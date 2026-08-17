package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type ResponsePayload struct {
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// Set cookie helper
func SetCookie(w http.ResponseWriter, name, token string, expiry time.Time) {
	http.SetCookie(w, &http.Cookie{
		Name:     name,
		Value:    token,
		Expires:  expiry,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	})
}

// Returns true if no body else false
func ValidateBody(w http.ResponseWriter, reqBody io.ReadCloser) bool {
	if reqBody == nil {
		SendJson(w, 400, "Body not provided", nil)
		return true
	}

	return false
}

// Get paramter and return true if found, else false
func GetParam(w http.ResponseWriter, r *http.Request, search string) (string, bool) {
	if found := r.PathValue(search); found == "" {
		SendJson(w, 400, "Path variable "+search+" not found", nil)
		return "", false
	} else {
		return found, true
	}
}

// Write json response, Returns true if error parsing json else false
func SendJson(w http.ResponseWriter, statusCode int, message string, data any) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(statusCode)

	payload := ResponsePayload{
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(payload); err != nil {
		w.Write([]byte(message))
	}
}

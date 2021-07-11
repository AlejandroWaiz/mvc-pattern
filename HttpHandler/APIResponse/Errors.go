package apiresponse

import (
	"encoding/json"
	"net/http"
)

var (
	ErrBadRequest  = WebError{StatusCode: http.StatusBadRequest, Type: "api_error", Message: "Cannot process current request"}
	ErrInvalidJSON = WebError{StatusCode: http.StatusBadRequest, Type: "invalid_json", Message: "Invalid or malformed JSON"}
)

type WebError struct {
	StatusCode int    `json:"-"`
	Type       string `json:"type"`
	Message    string `json:"message,omitempty"`
}

func (we WebError) Send(w http.ResponseWriter) error {
	statusCode := we.StatusCode
	if statusCode == 0 {
		statusCode = http.StatusBadRequest

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(we)
}

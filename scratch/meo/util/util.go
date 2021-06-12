package util

import (
	"encoding/json"
	"net/http"
)

func RespondOkay(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func RespondWithError(w http.ResponseWriter, statusCode int, message string) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	b := map[string]string{
		"error": message,
	}

	json.NewEncoder(w).Encode(b)

}

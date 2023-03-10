package data

import (
	"encoding/json"
	"io"
	"net/http"
)

func FromJSON(v any, d io.ReadCloser) error {
	err := json.NewDecoder(d).Decode(v)
	return err
}

func ToJSON(w http.ResponseWriter, v any, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	err := json.NewEncoder(w).Encode(v)
	return err
}

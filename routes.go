package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Message struct {
	Message string `json:"message"`
}

func routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(&Message{
			Message: "API is running!",
		})
	})

	return r
}

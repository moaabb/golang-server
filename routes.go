package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/moaabb/golang-server/data"
)

type Message struct {
	Message string `json:"message"`
}

func routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		data.ToJSON(w, &Message{
			Message: fmt.Sprintf("API is running on port %s!", cfg.PORT),
		}, http.StatusOK)
	})

	return r
}

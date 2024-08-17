package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	server := chi.NewRouter()
	server.Use(middleware.Recoverer)

	server.Get("/", app.Hello)

	return server
}

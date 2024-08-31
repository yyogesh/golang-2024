package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()

	// register middleware
	router.Use(middleware.Recoverer)

	// define routes
	router.Get("/", app.Home)

	return router
}

package main

import (
	"bank_application/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func (app *application) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world from %s", app.Domain)

	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "active",
		Message: "Hello, world!",
		Version: "1.0.0",
	}

	out, err := json.Marshal(payload)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm %s, and I'm healthy!", app.Domain)
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie

	abc := models.Movie{}

	movies = append(movies, abc)

	out, err := json.Marshal(movies)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

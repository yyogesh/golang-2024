package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplate = "./templates/"

type TemplateData struct {
	IP   string
	Data map[string]any
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, world from %s", app.Domain)
	w.Header().Set("Content-Type", "text/html")
	app.render(w, "home.page.gohtml", &TemplateData{})
}

func (app *application) render(w http.ResponseWriter, tmpl string, data *TemplateData) error {
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplate, tmpl))

	if err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return err
	}

	data.IP = "198.158.938.11"

	err = parsedTemplate.Execute(w, data)

	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return err
	}

	return nil
}

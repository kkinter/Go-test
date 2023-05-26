package main

import (
	"html/template"
	"net/http"
	"path"
)

var pathToTemplates = "./templates/"

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	_ = app.render(w, r, "home.page.gohtml", &TemplateData{})
}

type TemplateData struct {
	IP   string
	data map[string]any
}

func (app *application) render(w http.ResponseWriter, r *http.Request, tmpl string, data *TemplateData) error {
	// parse the template from disk.
	parsedTemplate, err := template.ParseFiles(path.Join(pathToTemplates, tmpl))
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return err
	}

	data.IP = app.ipFromContext(r.Context())

	// execute the template, passing it data, if any
	err = parsedTemplate.Execute(w, data)
	if err != nil {
		return err
	}

	return nil
}

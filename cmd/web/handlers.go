package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	tmplFiles := []string{
		"./ui/html/base.tmpl",
		"./ui/html/partials/nav.tmpl",
		"./ui/html/pages/home.tmpl",
	}
	templateSet, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	err = templateSet.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.serverError(w, r, err)
	}
}

func (app *Application) getSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprintf(w, "Display snippet with id %d\n", id)
}

func (app *Application) getSnippetForm(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Display the form to create a snippet\n"))
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, "Create/save a new snippet")
}

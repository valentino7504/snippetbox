package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/valentino7504/snippetbox/internals/models"
)

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	// snippets, err := app.snippets.Latest()
	// if err != nil {
	// 	app.serverError(w, r, err)
	// 	return
	// }
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
	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			http.NotFound(w, r)
		} else {
			app.serverError(w, r, err)
		}
		return
	}
	_, _ = fmt.Fprintf(w, "Display snippet with id %d: %+v\n", id, snippet)
}

func (app *Application) getSnippetForm(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Display the form to create a snippet\n"))
}

func (app *Application) createSnippet(w http.ResponseWriter, r *http.Request) {
	title := "Nigeria Anthem thing"
	content := "Nigeria We Hail Thee"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	// w.WriteHeader(http.StatusCreated)
	// _, _ = fmt.Fprintf(w, "Create/save a new snippet")
	// Redirect the user to the view of the snippet
	http.Redirect(w, r, fmt.Sprintf("/snippets/%d", id), http.StatusSeeOther)
}

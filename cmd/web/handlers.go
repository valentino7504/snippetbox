package main

import (
	"errors"
	"fmt"
	_ "html/template"
	"net/http"
	"strconv"

	"github.com/valentino7504/snippetbox/internal/models"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
}

func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	data := app.newTemplateData(r)
	data.Snippets = snippets
	app.render(w, r, http.StatusOK, "home.tmpl", data)
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
	data := app.newTemplateData(r)
	data.Snippet = snippet
	app.render(w, r, http.StatusOK, "view.tmpl", data)
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

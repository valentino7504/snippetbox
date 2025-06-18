package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	_, _ = w.Write([]byte("Welcome to Snippetbox!"))
}

func getSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	_, _ = fmt.Fprintf(w, "Display snippet with id %d\n", id)
}

func getSnippetForm(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Display the form to create a snippet\n"))
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	_, _ = fmt.Fprintf(w, "Create/save a new snippet")
}

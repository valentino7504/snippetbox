package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method, uri, trace := r.Method, r.URL.RequestURI(), string(debug.Stack())
	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, status int, page string, data templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)

	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	w.WriteHeader(status)
	_, err = buf.WriteTo(w)
	if err != nil {
		app.serverError(w, r, err)
		return
	}
}

func newLogger(debugFlag *bool) *slog.Logger {
	var logLevel slog.Level
	if *debugFlag {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
	return logger
}

func (app *Application) newTemplateData(r *http.Request) templateData {
	return templateData{CurrentYear: time.Now().Year()}
}

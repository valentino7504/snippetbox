package main

import (
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
)

func NewLogger(debugFlag *bool) *slog.Logger {
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

func (app *Application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	method, uri, trace := r.Method, r.URL.RequestURI(), string(debug.Stack())
	app.logger.Error(err.Error(), "method", method, "uri", uri, "trace", trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

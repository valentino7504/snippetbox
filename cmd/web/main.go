package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", "4000", "HTTP network address")
	debug := flag.Bool("debug", false, "Debug mode on/off")
	flag.Parse()

	var logLevel slog.Level
	if *debug {
		logLevel = slog.LevelDebug
	} else {
		logLevel = slog.LevelInfo
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))

	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))

	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	app := &Application{
		logger: logger,
	}
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippets/{id}", app.getSnippet)
	mux.HandleFunc("GET snippets/create", app.getSnippetForm)
	mux.HandleFunc("POST snippets/create", app.createSnippet)

	logger.Info("Starting a new server", slog.Any("port", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), mux)
	logger.Error(err.Error())
	os.Exit(1)
}

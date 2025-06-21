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

	app := &Application{
		logger: NewLogger(debug),
	}

	app.logger.Info("Starting a new server", slog.Any("port", *port))
	err := http.ListenAndServe(fmt.Sprintf(":%s", *port), app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

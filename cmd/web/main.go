package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/valentino7504/snippetbox/internals/models"
)

type Application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	port := flag.String("port", "4000", "HTTP network address")
	debug := flag.Bool("debug", false, "Debug mode on/off")
	dbUrl := flag.String("db", os.Getenv("SNIPPETBOX_DB_URL"), "MySQL data source name")

	logger := NewLogger(debug)

	flag.Parse()

	db, err := openDB(*dbUrl)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &Application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	defer func(db *sql.DB) {
		if err := db.Close(); err != nil {
			app.logger.Warn(err.Error())
		}
	}(db)
	app.logger.Info("Starting a new server", slog.Any("port", *port))
	err = http.ListenAndServe(fmt.Sprintf(":%s", *port), app.routes())
	app.logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		_ = db.Close()
		return nil, err
	}
	return db, nil
}

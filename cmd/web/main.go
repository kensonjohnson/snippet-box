package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Take in command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "user:password@tcp/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Setup logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Setup DB pool
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	// Dependency injection
	app := &application{
		logger: logger,
	}

	// Start webserver
	logger.Info(fmt.Sprintf("starting server at http://localhost%s", *addr))
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

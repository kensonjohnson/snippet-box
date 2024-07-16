package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	// Take in command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Setup logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Dependency injection
	app := &application{
		logger: logger,
	}

	// Start webserver
	logger.Info(fmt.Sprintf("starting server at http://localhost%s", *addr))
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

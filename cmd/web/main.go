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

	// Create servemux
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Define routes
	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetView)
	mux.HandleFunc("GET /snippet/create", app.snippetCreate)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)

	// Start webserver
	logger.Info(fmt.Sprintf("starting server at http://localhost%s", *addr))
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

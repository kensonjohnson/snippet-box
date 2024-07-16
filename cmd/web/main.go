package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Take in command line flags
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// Create servemux
	mux := http.NewServeMux()

	// Serve static files
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	// Define routes
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	// Start webserver
	log.Printf("starting server at http://localhost%s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}

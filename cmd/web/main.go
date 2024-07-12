package main

import (
	"log"
	"net/http"
)

func main() {
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
	log.Print("starting server at http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

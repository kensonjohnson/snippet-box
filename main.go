package main

import (
	"log"
	"net/http"
)

func main() {
	// Create the servemux
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Start the server
	log.Print("Server starting at http://localhost:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

// Define a home handler function which write a byte slice
// containing "Hello World" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a form to create a snippet..."))
}

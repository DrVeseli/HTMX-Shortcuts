package main

import "net/http"

func routes() http.Handler {
	// Create a new ServeMux
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/style", styleHandler)
	return mux
}

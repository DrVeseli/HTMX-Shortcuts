package main

import (
	"log/slog"
	"net/http"
)

// Handler function for the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {

		slog.Info("Error executing template: %v", "Error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func styleHandler(w http.ResponseWriter, r *http.Request) {

}

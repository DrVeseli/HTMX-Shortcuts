package main

import (
	"log/slog"
	"net/http"
)

// Handler function for the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := tmpl.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		slog.Error("Error executing template: %v", "Error:", err)
	}
}

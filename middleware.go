package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

// Middleware to log requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Middleware to recover from panics
func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				slog.Info("Panic recovered: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Internal Server Error")
			}
		}()
		next.ServeHTTP(w, r)
	})
}

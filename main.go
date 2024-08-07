package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"text/template"
	"time"
)

var (
	tmpl *template.Template
	err  error
)

func init() {
	// Parse all .html files in the "src" directory
	tmpl, err = template.ParseFS(os.DirFS("src"), "*.html")
	if err != nil {
		slog.Error("Error parsing templates: %v", "Error:", err)
	}
}

func main() {
	mux := routes()

	// Wrap the mux with middleware
	handler := loggingMiddleware(recoveryMiddleware(mux))

	// Create a new HTTP server
	server := &http.Server{
		Addr:         ":3600",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	// Start the server in a separate goroutine
	go func() {
		slog.Info("Starting server on :3600")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Could not listen on %s: %v", server.Addr, err)
		}
	}()

	// Create a channel to listen for OS signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received
	<-stop

	// Create a context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown the server gracefully
	slog.Info("Shutting down server...")
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Could not gracefully shutdown the server: %v", "Error:", err)
	}
	slog.Info("Server stopped")
}

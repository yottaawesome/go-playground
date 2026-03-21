// Package main is the entry point for the Book Server API.
// It wires together the service layer, HTTP handlers, and router,
// then starts the HTTP server.
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/example/go-sample-project/internal/handler"
	"github.com/example/go-sample-project/internal/service"
)

func main() {
	// Configure zerolog for pretty console output during development.
	log.Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).
		With().
		Timestamp().
		Logger()

	// Read the port from the environment, defaulting to 8080.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// --- Dependency Wiring ---
	// Create the service (business logic) layer, then inject it into the handler.
	bookSvc := service.NewBookService()
	bookHandler := handler.NewBookHandler(bookSvc)

	// --- Router Setup ---
	r := chi.NewRouter()

	// Built-in chi middleware for request logging and panic recovery.
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a request timeout to prevent long-running requests.
	r.Use(middleware.Timeout(30 * time.Second))

	// Health-check endpoint — useful for load balancers and k8s probes.
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, `{"status":"ok"}`)
	})

	// Mount book routes under the /api/v1/books namespace.
	r.Route("/api/v1/books", func(r chi.Router) {
		r.Get("/", bookHandler.ListBooks)       // GET    /api/v1/books
		r.Post("/", bookHandler.CreateBook)      // POST   /api/v1/books
		r.Get("/{id}", bookHandler.GetBook)      // GET    /api/v1/books/{id}
		r.Put("/{id}", bookHandler.UpdateBook)   // PUT    /api/v1/books/{id}
		r.Delete("/{id}", bookHandler.DeleteBook) // DELETE /api/v1/books/{id}
	})

	// --- HTTP Server with Graceful Shutdown ---
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start the server in a goroutine so we can listen for shutdown signals.
	go func() {
		log.Info().Str("port", port).Msg("starting server")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("server failed")
		}
	}()

	// Block until we receive SIGINT or SIGTERM.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Info().Msg("shutting down server…")

	// Give outstanding requests 10 seconds to complete.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("forced shutdown")
	}

	log.Info().Msg("server stopped gracefully")
}

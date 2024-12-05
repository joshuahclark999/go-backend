// Package main implements a simple HTTP server with a health check endpoint.
package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joshuahclark999/go-backend/internal/store"
)

// application struct holds the configuration for the application.
type application struct {
	config config
	store  store.Storage
}

// config struct holds the server address configuration.
type config struct {
	addr string
}

// mount sets up the HTTP routes and returns an HTTP request multiplexer (ServeMux).
// It registers the health check endpoint.
func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))
	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheck)
	})

	return r
}

// run starts the HTTP server with the provided ServeMux and configuration settings.
// It sets timeouts for writing, reading, and idle connections.
func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,  // Server address from the configuration.
		Handler:      mux,              // HTTP request multiplexer.
		WriteTimeout: time.Second * 30, // Maximum duration before timing out writes of the response.
		ReadTimeout:  time.Second * 10, // Maximum duration for reading the entire request, including the body.
		IdleTimeout:  time.Minute,      // Maximum amount of time to wait for the next request when keep-alives are enabled.
	}
	log.Printf("server started at:%s", app.config.addr) // Log the server start message with the address.
	return srv.ListenAndServe()                         // Start the HTTP server.
}

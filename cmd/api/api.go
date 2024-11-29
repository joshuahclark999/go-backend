// Package main implements a simple HTTP server with a health check endpoint.
package main

import (
	"log"
	"net/http"
	"time"
)

// application struct holds the configuration for the application.
type application struct {
	config config
}

// config struct holds the server address configuration.
type config struct {
	addr string
}

// mount sets up the HTTP routes and returns an HTTP request multiplexer (ServeMux).
// It registers the health check endpoint.
func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()

	// Register the health check endpoint with the ServeMux.
	mux.HandleFunc("GET /v1/health", app.healthCheck)
	return mux
}

// run starts the HTTP server with the provided ServeMux and configuration settings.
// It sets timeouts for writing, reading, and idle connections.
func (app *application) run(mux *http.ServeMux) error {

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

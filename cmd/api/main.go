package main

import (
	"log"

	"github.com/joshuahclark999/go-backend/internal/store"
)

// main is the entry point for the application. It initializes the configuration,
// creates an application instance, mounts the routes, and starts the server.
//
// The function performs the following steps:
//  1. Initializes a config struct with the server address.
//  2. Creates an application instance with the proviipded configuration.
//  3. Mounts the routes using the application's mount method.
//  4. Starts the server by calling the application's run method with the mux (router).
//     If the server fails to start, it logs the error and exits the program.
func main() {
	cfg := config{
		addr: ":8080",
	}
	store := store.NewStorage(nil)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}

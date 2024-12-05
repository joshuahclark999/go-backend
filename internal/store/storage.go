// Package store provides the storage layer for the application,
// including interfaces for creating posts and users in the database.
package store

import (
	"context"
	"database/sql"
)

// Storage struct holds the interfaces for interacting with the database.
// It includes methods for creating posts and users.
type Storage struct {
	Posts interface {
		// Create inserts a new post into the database.
		// It takes a context for managing request deadlines and cancellation.
		Create(context.Context) error
	}
	Users interface {
		// Create inserts a new user into the database.
		// It takes a context for managing request deadlines and cancellation.
		Create(context.Context) error
	}
}

// NewStorage initializes a new Storage instance with the given database connection.
// It returns a Storage struct with the necessary interfaces for database operations.
func NewStorage(db *sql.DB) Storage {
	return Storage{
		// Initialization of the interfaces should be done here.
		Posts: &PostsStore{db},
		Users: &UsersStore{db},
	}
}

package store

import (
	"context"
	"database/sql"
)

type UsersStore struct {
	db *sql.DB
}

// Create inserts a new post into the store.
// It takes a context.Context as a parameter to allow for request cancellation and timeout handling.
// Returns an error if the operation fails.
func (s *UsersStore) Create(ctx context.Context) error {
	return nil

}

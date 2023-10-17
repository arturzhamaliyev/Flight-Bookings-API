package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
)

// DB represents a type for interfacing with a postgres database.
type DB interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
}

// Store provides functionality for working with a postgres database.
type Store struct {
	db DB
}

// New will instantiate a new instance of Store.
func New(db DB) *Store {
	return &Store{
		db: db,
	}
}

// InsertUser will add a new unique user to the database using the provided data.
func (s *Store) InsertUser(ctx context.Context, user model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	_, err := s.db.
		ExecContext(
			ctx,
			`
			INSERT INTO users(
				first_name,
				last_name,
				password,
				email,
				country,
				created_at,
				updated_at
			)
			VALUES (
				$1, $2, $3, $4, $5, $6, $7
			)
			`,
			user)
	if err != nil {
		return err
	}

	return nil
}

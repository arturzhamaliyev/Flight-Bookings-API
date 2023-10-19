package repository

import "github.com/jmoiron/sqlx"

// Repository provides data access functions for various entities.
type Repository struct {
	Users *UsersRepository
}

// New creates a new Repository with the given database connection.
func New(db *sqlx.DB) Repository {
	return Repository{
		Users: NewUsersRepo(db),
	}
}

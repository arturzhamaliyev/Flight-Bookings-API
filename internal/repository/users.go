package repository

import (
	"context"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/jmoiron/sqlx"
)

const insertUser = `
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
`

// UsersRepository provides functionality for working with a postgres database.
type UsersRepository struct {
	db *sqlx.DB
}

// NewUsersRepo will instantiate a new instance of Repository.
func NewUsersRepo(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{
		db: db,
	}
}

// InsertUser will add a new unique user to the database using the provided data.
func (r *UsersRepository) InsertUser(ctx context.Context, user model.User) error {
	ctx, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	_, err := r.db.
		ExecContext(
			ctx,
			insertUser,
			user.FirstName, user.LastName, user.Password, user.Email, user.Country, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	return nil
}

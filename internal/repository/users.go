package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

const insertUserQuery = `
INSERT INTO users(
	id,
	phone,
	email,
	password,
	created_at,
	updated_at
)
VALUES (
	$1, $2, $3, $4, $5, $6
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
	_, err := r.db.
		ExecContext(
			ctx,
			insertUserQuery,
			user.ID, user.Phone, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			return customErrors.ErrUniqueViolation
		}
		return err
	}
	return nil
}

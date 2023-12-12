package repository

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

const (
	insertUserQuery = `
		INSERT INTO users(
			id,
			role,
			phone,
			email,
			password,
			created_at,
			updated_at
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7
		);
	`

	getUserByEmailQuery = `
		SELECT * 
		FROM users
		WHERE email = $1;
	`

	updateUserQuery = `
		UPDATE users
		SET phone = $1,
			email = $2,
			password = $3,
			updated_at = $4
		WHERE id = $5;
	`

	getUserByIDQuery = `
		SELECT *
		FROM users
		WHERE id = $1;
	`

	deleteUserByIDQuery = `
		DELETE FROM users
		WHERE id = $1;
	`
)

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
			user.ID,
			user.Role,
			user.Phone,
			user.Email,
			user.Password,
			user.CreatedAt,
			user.UpdatedAt,
		)
	if err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) && pgError.Code == pgerrcode.UniqueViolation {
			zap.S().Info(err)
			return customErrors.ErrUniqueViolation
		}
		zap.S().Info(err)
		return err
	}
	return nil
}

// GetUserByEmail
func (r *UsersRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	row := r.db.
		QueryRowContext(
			ctx,
			getUserByEmailQuery,
			email,
		)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)
	if err != nil {
		zap.S().Info(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UsersRepository) UpdateUser(ctx context.Context, user model.User) error {
	_, err := r.db.ExecContext(
		ctx,
		updateUserQuery,
		user.Phone,
		user.Email,
		user.Password,
		user.UpdatedAt,
		user.ID,
	)
	if err != nil {
		zap.S().Info(err)
		return err
	}
	return nil
}

func (r *UsersRepository) GetUserByID(ctx context.Context, ID string) (model.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		getUserByIDQuery,
		ID,
	)

	var user model.User
	err := row.Scan(
		&user.ID,
		&user.Phone,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.Role,
	)
	if err != nil {
		zap.S().Info(err)
		return model.User{}, err
	}

	return user, nil
}

func (r *UsersRepository) DeleteUserByID(ctx context.Context, ID string) error {
	_, err := r.db.ExecContext(
		ctx,
		deleteUserByIDQuery,
		ID,
	)
	if err != nil {
		zap.S().Info(err)
		return err
	}
	return nil
}

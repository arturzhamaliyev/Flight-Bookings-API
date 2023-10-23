package service

import (
	"context"
	"fmt"
	"net/mail"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type (
	// UsersRepository represents a type that provides operations on storing users in database.
	UsersRepository interface {
		InsertUser(ctx context.Context, user model.User) error
	}

	// DBError represents a type that provides errors that can occur in database.
	DBError interface {
		IsUniqueViolation(err error) bool
	}

	// Users represents a type that provides operations on users.
	Users struct {
		repo    UsersRepository
		dbError DBError
	}
)

// NewUsersService will instantiate a new instance of Users.
func NewUsersService(repo UsersRepository, err DBError) *Users {
	return &Users{
		repo:    repo,
		dbError: err,
	}
}

// CreateUser will try to create a user in our database.
func (u *Users) CreateUser(ctx context.Context, user model.User) error {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		return fmt.Errorf("invalid email address: %w", err)
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = u.repo.InsertUser(ctx, user)
	if err != nil {
		if u.dbError.IsUniqueViolation(err) {
			return fmt.Errorf("user already exists: %w", err)
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}
	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

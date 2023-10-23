package service

import (
	"context"
	"errors"
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

	// Users represents a type that provides operations on users.
	Users struct {
		repo UsersRepository
	}
)

// NewUsersService will instantiate a new instance of Users.
func NewUsersService(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
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
		if errors.Is(err, ErrUniqueViolation) {
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

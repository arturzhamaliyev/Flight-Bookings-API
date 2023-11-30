package service

import (
	"context"
	"errors"
	"fmt"
	"net/mail"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

type (
	// UsersRepository represents a type that provides operations on storing users in database.
	UsersRepository interface {
		InsertUser(ctx context.Context, user model.User) error
		GetUserByEmail(ctx context.Context, email string) (model.User, error)
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
		return errors.Join(ErrInvalidEmailAddress, err)
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	err = u.repo.InsertUser(ctx, user)
	if err != nil {
		if errors.Is(err, customErrors.ErrUniqueViolation) {
			return errors.Join(ErrUserExists, err)
		}
		return fmt.Errorf("couldn't create user: %w", err)
	}
	return nil
}

// CheckUserCredentials
func (u *Users) CheckUserCredentials(ctx context.Context, email, password string) error {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil || !comparePassword(user.Password, password) {
		zap.S().Infof("error in check user creds: %v", err)
		return ErrInvalidEmailOrPassword
	}
	return nil
}

// comparePassword compares hashed password of user with given password.
func comparePassword(hashedPassword, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

// hashPassword generates and returns bcrypt hash from the given password.
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

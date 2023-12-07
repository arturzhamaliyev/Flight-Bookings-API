package service

import (
	"context"
	"errors"
	"net/mail"
	"time"

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
		GetUserByID(ctx context.Context, ID string) (model.User, error)
		UpdateUser(ctx context.Context, user model.User) error
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

// GetUserByEmail
func (u *Users) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	user, err := u.repo.GetUserByEmail(ctx, email)
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, customErrors.ErrNoRows) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return user, nil
}

// CreateUser will try to create a user in our database.
func (u *Users) CreateUser(ctx context.Context, user model.User) error {
	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		zap.S().Info(err)
		return ErrInvalidEmailAddress
	}

	user.Password, err = hashPassword(user.Password)
	if err != nil {
		zap.S().Info(err)
		return ErrHashPassword
	}

	err = u.repo.InsertUser(ctx, user)
	if err != nil {
		if errors.Is(err, customErrors.ErrUniqueViolation) {
			zap.S().Info(err)
			return ErrUserExists
		}
		zap.S().Info(err)
		return err
	}
	return nil
}

// ValidateUserPassword
func (u *Users) ValidateUserPassword(hashedPassword, password string) error {
	if comparePassword(hashedPassword, password) {
		return nil
	}
	return ErrInvalidPassword
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

func (u *Users) UpdateUser(ctx context.Context, user model.User) error {
	foundUser, err := u.repo.GetUserByID(ctx, user.ID.String())
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, customErrors.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	foundUser.Password, err = hashPassword(user.Password)
	if err != nil {
		zap.S().Info(err)
		return ErrHashPassword
	}

	foundUser.Email = user.Email
	foundUser.Phone = user.Phone
	foundUser.UpdatedAt = time.Now()

	return u.repo.UpdateUser(ctx, foundUser)
}

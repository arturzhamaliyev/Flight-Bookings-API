package mocks

import (
	"context"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

// UsersServiceMock
type UsersServiceMock struct{}

// NewUsersServiceMock
func NewUsersServiceMock() *UsersServiceMock {
	return &UsersServiceMock{}
}

// GetUserByEmail
func (u *UsersServiceMock) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, nil
}

// CreateUser
func (u *UsersServiceMock) CreateUser(ctx context.Context, user model.User) error {
	return nil
}

// ValidateUserPassword
func (u *UsersServiceMock) ValidateUserPassword(hashedPassword, password string) error {
	return nil
}

package mocks

import (
	"context"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/google/uuid"
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

// UpdateUser
func (u *UsersServiceMock) UpdateUser(ctx context.Context, user model.User) error {
	return nil
}

// DeleteUserByID
func (u *UsersServiceMock) DeleteUserByID(ctx context.Context, ID uuid.UUID) error {
	return nil
}

// GetUserByID
func (u *UsersServiceMock) GetUserByID(ctx context.Context, ID uuid.UUID) (model.User, error) {
	return model.User{}, nil
}

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

// CreateUser
func (u *UsersServiceMock) CreateUser(ctx context.Context, user model.User) error {
	return nil
}
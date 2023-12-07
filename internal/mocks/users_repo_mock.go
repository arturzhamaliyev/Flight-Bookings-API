package mocks

import (
	"context"

	"github.com/google/uuid"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

// UsersRepositoryMock mock struct
type UsersRepositoryMock struct {
	db map[uuid.UUID]model.User
}

// NewUsersRepoMock mock construct
func NewUsersRepoMock() *UsersRepositoryMock {
	return &UsersRepositoryMock{
		db: make(map[uuid.UUID]model.User),
	}
}

// InsertUser mock method
func (r *UsersRepositoryMock) InsertUser(ctx context.Context, user model.User) error {
	for _, u := range r.db {
		if u.Email == user.Email {
			return service.ErrUserExists
		}
	}
	r.db[user.ID] = user
	return nil
}

// GetUserByEmail mock method
func (r *UsersRepositoryMock) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	return model.User{}, nil
}

// UpdateUser
func (u *UsersRepositoryMock) UpdateUser(ctx context.Context, user model.User) error {
	return nil
}

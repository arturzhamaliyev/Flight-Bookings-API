package service

import (
	"context"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

// UsersRepository represents a type that provides operations on storing users in database.
type UsersRepository interface {
	InsertUser(ctx context.Context, user model.User) error
}

// Users represents a type that provides operations on users.
type Users struct {
	repo UsersRepository
}

// NewUsersService will instantiate a new instance of Users.
func NewUsersService(repo UsersRepository) *Users {
	return &Users{
		repo: repo,
	}
}

// CreateUser will try to create a user in our database.
func (u *Users) CreateUser(ctx context.Context, user model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	return u.repo.InsertUser(ctx, user)
}

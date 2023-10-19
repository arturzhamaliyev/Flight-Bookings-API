package users

import (
	"context"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
)

// Store represents a type that provides operations on storing users in database.
type Store interface {
	InsertUser(ctx context.Context, user model.User) error
}

// Users represents a type that provides operations on users.
type Users struct {
	store Store
}

// New will instantiate a new instance of Users.
func New(s Store) *Users {
	return &Users{
		store: s,
	}
}

// CreateUser will try to create a user in our database.
func (u *Users) CreateUser(ctx context.Context, user model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt

	return u.store.InsertUser(ctx, user)
}

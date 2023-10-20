package service

import (
	"context"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/google/uuid"
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
func (u *Users) CreateUser(ctx context.Context, userReq model.CreateUserRequest) error {
	user := model.User{
		ID:        uuid.New(),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Password:  userReq.Password,
		Email:     userReq.Email,
		Country:   userReq.Country,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return u.repo.InsertUser(ctx, user)
}

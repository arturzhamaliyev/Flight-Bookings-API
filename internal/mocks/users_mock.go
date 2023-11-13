package mocks

import (
	"context"

	"github.com/google/uuid"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

// UserMockRepository mock struct
type UserMockRepository struct {
	db map[uuid.UUID]model.User
}

// NewUsersMockRepo mock construct
func NewUsersMockRepo() *UserMockRepository {
	return &UserMockRepository{
		db: make(map[uuid.UUID]model.User),
	}
}

// InsertUser mock method
func (r *UserMockRepository) InsertUser(ctx context.Context, user model.User) error {
	for _, u := range r.db {
		if u.Email == user.Email {
			return service.ErrUserExists
		}
	}
	r.db[user.ID] = user
	return nil
}

package service

import (
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/repository"
)

// Service provides user-related functionality
type Service struct {
	User *Users
}

// New creates a new Service with the given repository.
func New(repo repository.Repository) Service {
	return Service{
		User: NewUsersService(repo.Users),
	}
}

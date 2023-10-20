package response

import (
	"github.com/google/uuid"
)

// CreateUser represents response object of person using this service.
type CreateUser struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
}

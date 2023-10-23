package response

import (
	"github.com/google/uuid"
)

// CreateUser represents response object of person using this service.
type CreateUser struct {
	ID    uuid.UUID `json:"id"`
	Phone *string   `json:"phone,omitempty"`
	Email string    `json:"email"`
}

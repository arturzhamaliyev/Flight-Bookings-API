package response

import (
	"github.com/google/uuid"
)

// SignUp represents response object of person using this service.
type SignUp struct {
	ID    uuid.UUID `json:"id"`
	Phone *string   `json:"phone,omitempty"`
	Email string    `json:"email"`
}

// SignIn
type SignIn struct {
	Token string `json:"token"`
}

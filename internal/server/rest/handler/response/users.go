package response

import (
	"time"

	"github.com/google/uuid"
)

// SignUp represents response object of person using this service.
type SignUp struct {
	ID    uuid.UUID `json:"id"`
	Phone *string   `json:"phone,omitempty"`
	Email string    `json:"email"`
}

type UpdateProfile struct {
	ID        uuid.UUID `json:"id"`
	Phone     *string   `json:"phone,omitempty"`
	Email     string    `json:"email"`
	UpdatedAt time.Time `json:"updatedAt"`
}

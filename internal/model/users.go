package model

import (
	"time"

	"github.com/google/uuid"
)

// User represents a person using this service.
type User struct {
	ID        uuid.UUID
	Role      Role
	Phone     *string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

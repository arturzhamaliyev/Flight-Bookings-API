package model

import "time"

// User represents a person using this service.
type User struct {
	ID        string    `json:"id,omitempty"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Country   string    `json:"country"`
	CreatedAt time.Time `json:"create_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

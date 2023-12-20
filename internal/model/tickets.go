package model

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Flight    Flight
	User      *User
	Rank      Rank
	Price     uint
}

func (t Ticket) IsAvailable() bool {
	return t.User != nil
}

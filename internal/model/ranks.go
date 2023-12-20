package model

import "github.com/google/uuid"

type RankType string

const (
	Economy  RankType = "economy"
	Business RankType = "business"
	Deluxe   RankType = "deluxe"
)

func (rt RankType) String() string {
	return string(rt)
}

type Rank struct {
	ID   uuid.UUID
	Name RankType
}

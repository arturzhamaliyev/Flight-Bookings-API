package model

import (
	"math"

	"github.com/google/uuid"
)

type AirplaneModel string

const (
	AirbusA320 AirplaneModel = "AirbusA320"
	Boeing787  AirplaneModel = "Boeing787"
)

func (am AirplaneModel) String() string {
	return string(am)
}

type Airplane struct {
	ID                 uuid.UUID
	Model              AirplaneModel
	TotalNumberOfSeats uint16
}

func (a Airplane) GetID() uuid.UUID {
	return a.ID
}

func (a Airplane) GetModel() AirplaneModel {
	return a.Model
}

func (a Airplane) GetTotalNumberOfSeats() uint16 {
	return a.TotalNumberOfSeats
}

// Default: 75% Economy, 20% Business, 5% Deluxe.
func (a Airplane) NumberOfSeatsByRank(rank RankType) uint16 {
	totalNumberOfSeats := float64(a.GetTotalNumberOfSeats())
	var res float64

	switch rank {
	case Economy:
		res = totalNumberOfSeats * 0.75
	case Business:
		res = totalNumberOfSeats * 0.2
	case Deluxe:
		res = totalNumberOfSeats * 0.05
	}
	res = math.Floor(res)

	return uint16(res)
}

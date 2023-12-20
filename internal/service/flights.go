package service

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

type (
	FlightsRepository interface {
		InsertFlight(ctx context.Context, flight model.Flight) error
	}

	RanksRepository interface {
		GetAllRanks(ctx context.Context) ([]model.Rank, error)
	}

	AirplanesRepository interface {
		GetAirplaneByID(ctx context.Context, ID string) (model.Airplane, error)
		GetAllAirplanes(ctx context.Context) ([]model.Airplane, error)
	}

	TicketsRepository interface {
		InsertTickets(ctx context.Context, tickets []model.Ticket) error
	}

	FlightsService struct {
		flightsRepo   FlightsRepository
		ranksRepo     RanksRepository
		airplanesRepo AirplanesRepository
		ticketsRepo   TicketsRepository
	}
)

func NewFlightsService(
	flightsRepo FlightsRepository,
	ranksRepo RanksRepository,
	airplanesRepo AirplanesRepository,
	ticketsRepo TicketsRepository,
) *FlightsService {
	return &FlightsService{
		flightsRepo:   flightsRepo,
		ranksRepo:     ranksRepo,
		airplanesRepo: airplanesRepo,
		ticketsRepo:   ticketsRepo,
	}
}

func (f *FlightsService) CreateFlight(ctx context.Context, flight *model.Flight) error {
	airplane, err := f.airplanesRepo.GetAirplaneByID(ctx, flight.Airplane.GetID().String())
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, customErrors.ErrNoRows) {
			return ErrAirplaneNotFound
		}
		return err
	}

	flight.Airplane = airplane
	flight.Tickets = make([]model.Ticket, flight.Airplane.GetTotalNumberOfSeats())

	ranks, err := f.ranksRepo.GetAllRanks(ctx)
	if err != nil {
		zap.S().Info(err)
		return err
	}

	numberOfSeatsByRank := make(map[model.Rank]uint16)
	for _, rank := range ranks {
		numberOfSeatsByRank[rank] = flight.
			Airplane.
			NumberOfSeatsByRank(rank.Name)
	}

	curTime := time.Now()
	ticketsCount := len(flight.Tickets)
	ticketNumber := 0

	for rank, numberOfSeats := range numberOfSeatsByRank {
		for numberOfSeats > 0 && ticketNumber < ticketsCount {
			flight.Tickets[ticketNumber].ID = uuid.New()
			flight.Tickets[ticketNumber].CreatedAt = curTime
			flight.Tickets[ticketNumber].Flight = *flight
			flight.Tickets[ticketNumber].Rank = rank

			ticketNumber++
			numberOfSeats--
		}
	}

	err = f.flightsRepo.InsertFlight(ctx, *flight)
	if err != nil {
		zap.S().Info(err)
		return err
	}

	err = f.ticketsRepo.InsertTickets(ctx, flight.Tickets)
	if err != nil {
		zap.S().Info(err)
		return err
	}

	return nil
}

func (f *FlightsService) GetAllAirplanes(ctx context.Context) ([]model.Airplane, error) {
	airplanes, err := f.airplanesRepo.GetAllAirplanes(ctx)
	if err != nil {
		if errors.Is(err, customErrors.ErrNoRows) {
			return nil, ErrAirplanesNotFound
		}
		return nil, err
	}
	return airplanes, nil
}

package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
)

const (
	insertFlightQuery = `
		INSERT INTO flights(
			id,
			start_date,
			end_date,
			departure_id,
			destination_id,
			created_at,
			airplane_id
		)
		VALUES (
			$1, $2, $3, $4, $5, $6, $7
		);
	`
)

type FlightsRepository struct {
	db *sqlx.DB
}

func NewFlightsRepo(db *sqlx.DB) *FlightsRepository {
	return &FlightsRepository{
		db: db,
	}
}

func (f *FlightsRepository) InsertFlight(ctx context.Context, flight model.Flight) error {
	_, err := f.db.ExecContext(
		ctx,
		insertFlightQuery,
		flight.ID,
		flight.StartDate,
		flight.EndDate,
		flight.Departure.ID,
		flight.Destination.ID,
		flight.CreatedAt,
		flight.Airplane.ID,
	)
	if err != nil {
		zap.S().Info(err)
		return err
	}

	return nil
}

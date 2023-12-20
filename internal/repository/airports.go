package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
)

const (
	getAirportByCoordinatesQuery = `
		SELECT *
		FROM airports
		WHERE (location).latitude = $1
		AND (location).longitude = $2;
	`

	insertAirportQuery = `
		INSERT INTO airports(
			id,
			name,
			city,
			country,
			location	
		)
		VALUES($1,$2,$3,$4,($5,$6));
	`
)

type AirportsRepository struct {
	db *sqlx.DB
}

func NewAirportsRepo(db *sqlx.DB) *AirportsRepository {
	return &AirportsRepository{
		db: db,
	}
}

func (a *AirportsRepository) GetAirportByCoordinates(ctx context.Context, coordinates model.Coordinates) (model.Airport, error) {
	row := a.db.QueryRowContext(
		ctx,
		getAirportByCoordinatesQuery,
		coordinates.Latitude,
		coordinates.Longitude,
	)

	var airport model.Airport
	err := row.Scan(
		&airport.ID,
		&airport.Name,
		&airport.City,
		&airport.Country,
		&airport.Coordinates,
	)
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, sql.ErrNoRows) {
			return model.Airport{}, customErrors.ErrNoRows
		}
		return model.Airport{}, err
	}

	return airport, nil
}

func (a *AirportsRepository) InsertAirport(ctx context.Context, airport model.Airport) error {
	_, err := a.db.ExecContext(
		ctx,
		insertAirportQuery,
		airport.ID,
		airport.Name,
		airport.City,
		airport.Country,
		airport.Coordinates.Latitude,
		airport.Coordinates.Longitude,
	)
	if err != nil {
		return err
	}
	return nil
}

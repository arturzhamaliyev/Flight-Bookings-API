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
	getAllAirplanesQuery = `
		SELECT * FROM airplanes;
	`

	getAirplaneByIDQuery = `
		SELECT * FROM airplanes
		WHERE id = $1;
	`
)

type AirplanesRepository struct {
	db *sqlx.DB
}

func NewAirplanesRepo(db *sqlx.DB) *AirplanesRepository {
	return &AirplanesRepository{db: db}
}

func (a *AirplanesRepository) GetAllAirplanes(ctx context.Context) ([]model.Airplane, error) {
	rows, err := a.db.QueryContext(
		ctx,
		getAllAirplanesQuery,
	)
	if err != nil {
		zap.S().Info(err)
		if errors.Is(err, sql.ErrNoRows) {
			return nil, customErrors.ErrNoRows
		}
		return nil, err
	}
	defer rows.Close()

	var airplanes []model.Airplane

	for rows.Next() {
		var airplane model.Airplane
		err = rows.Scan(
			&airplane.ID,
			&airplane.Model,
			&airplane.TotalNumberOfSeats,
		)
		if err != nil {
			zap.S().Info(err)
			return nil, err
		}

		airplanes = append(airplanes, airplane)
	}

	err = rows.Err()
	if err != nil {
		zap.S().Info(err)
		return nil, err
	}

	return airplanes, nil
}

func (a *AirplanesRepository) GetAirplaneByID(ctx context.Context, ID string) (model.Airplane, error) {
	var airplane model.Airplane
	err := a.db.QueryRowContext(
		ctx,
		getAirplaneByIDQuery,
		ID,
	).Scan(
		&airplane.ID,
		&airplane.Model,
		&airplane.TotalNumberOfSeats,
	)
	if err != nil {
		return model.Airplane{}, err
	}

	return airplane, nil
}

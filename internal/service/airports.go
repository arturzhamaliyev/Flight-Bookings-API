package service

import (
	"context"
	"errors"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/config"
	customErrors "github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/helper"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type (
	AirportsRepository interface {
		GetAirportByCoordinates(ctx context.Context, coordinates model.Coordinates) (model.Airport, error)
		InsertAirport(ctx context.Context, airport model.Airport) error
	}

	AirportsService struct {
		cfg  config.Config
		repo AirportsRepository
	}
)

func NewAirportsService(cfg config.Config, repo AirportsRepository) *AirportsService {
	return &AirportsService{
		cfg:  cfg,
		repo: repo,
	}
}

func (f *AirportsService) FindAirportByCoordinates(ctx context.Context, coordinates model.Coordinates) (model.Airport, error) {
	resp, err := helper.SendRequestToGetExactCoordinatesOfAirport(coordinates, f.cfg.GoogleMaps.APIKey, f.cfg.GoogleMaps.URL)
	if err != nil {
		zap.S().Info(err)
		return model.Airport{}, err
	}

	var city, country string
	for _, component := range resp.Places[0].AddressComponents {
		if component.Types[0] == "administrative_area_level_1" {
			city = component.LongText
		} else if component.Types[0] == "country" {
			country = component.LongText
		}
	}

	exactCoordinates := resp.Places[0].Location
	airport, err := f.repo.GetAirportByCoordinates(ctx, exactCoordinates)
	if err != nil {
		if errors.Is(err, customErrors.ErrNoRows) {
			airport = model.Airport{
				ID:          uuid.New(),
				Name:        resp.Places[0].Name.Text,
				City:        city,
				Country:     country,
				Coordinates: exactCoordinates,
			}

			err = f.repo.InsertAirport(ctx, airport)
			if err != nil {
				zap.S().Info(err)
				return model.Airport{}, err
			}
			return airport, nil
		}
		zap.S().Info(err)
		return model.Airport{}, err
	}

	return airport, nil
}

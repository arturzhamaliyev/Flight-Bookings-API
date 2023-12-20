package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/request"
)

type (
	AirportsService interface {
		FindAirportByCoordinates(ctx context.Context, coordinates model.Coordinates) (model.Airport, error)
	}

	FlightsService interface {
		CreateFlight(ctx context.Context, flight *model.Flight) error
		GetAllAirplanes(ctx context.Context) ([]model.Airplane, error)
	}
)

func (h *Handler) CreateFlight(ctx *gin.Context) {
	var flightReq request.CreateFlight
	err := ctx.ShouldBindJSON(&flightReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	departureAirport, err := h.airportsService.FindAirportByCoordinates(ctx, flightReq.Departure)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	destinationAirport, err := h.airportsService.FindAirportByCoordinates(ctx, flightReq.Destination)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	flight := model.Flight{
		ID:          uuid.New(),
		StartDate:   flightReq.StartDate,
		EndDate:     flightReq.EndDate,
		CreatedAt:   time.Now(),
		Departure:   departureAirport,
		Destination: destinationAirport,
		Airplane: model.Airplane{
			ID: flightReq.AirplaneID,
		},
	}

	err = h.flightsService.CreateFlight(ctx, &flight)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// this is done due to unnesacery show of empty tickets
	flight.Tickets = nil

	ctx.JSON(http.StatusOK, flight)
}

func (h *Handler) GetAllAirplanes(ctx *gin.Context) {
	airplanes, err := h.flightsService.GetAllAirplanes(ctx)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, airplanes)
}

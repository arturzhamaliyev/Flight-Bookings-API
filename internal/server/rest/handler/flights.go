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

type AirportsService interface {
	FindAirportByCoordinates(ctx context.Context, coordinates model.Coordinates) (model.Airport, error)
}

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
		Tickets:     []model.Ticket{},
	}

	ctx.JSON(http.StatusOK, flight)
}

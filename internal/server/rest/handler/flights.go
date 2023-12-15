package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/helper"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler/request"
)

type FlightsService interface{}

func (h *Handler) CreateFlight(ctx *gin.Context) {
	var flightReq request.CreateFlight
	err := ctx.ShouldBindJSON(&flightReq)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	departureAirport, err := helper.FindAirportByCoordinates(flightReq.Departure)
	if err != nil {
		zap.S().Info(err)
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	destinationAirport, err := helper.FindAirportByCoordinates(flightReq.Destination)
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

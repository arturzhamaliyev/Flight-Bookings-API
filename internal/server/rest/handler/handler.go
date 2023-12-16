package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for flight bookings.
type Handler struct {
	usersService    UsersService
	airportsService AirportsService
}

// New will instantiate a new instance of Handler.
func New(usersService UsersService, airportsService AirportsService) Handler {
	return Handler{
		usersService:    usersService,
		airportsService: airportsService,
	}
}

// HealthCheck will response with OK status if server is listening for incoming requests.
// Will replace with Swagger soon.
func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

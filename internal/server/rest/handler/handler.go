package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for flight bookings.
type Handler struct {
	usersService   UsersService
	sessionService SessionService
}

// New will instantiate a new instance of Handler.
func New(usersService UsersService, sessionService SessionService) Handler {
	return Handler{
		usersService:   usersService,
		sessionService: sessionService,
	}
}

// HealthCheck will response with OK status if server is listening for incoming requests.
// Will replace with Swagger soon.
func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

package handler

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for flight bookings.
type Handler struct {
	usersService Users
}

// New will instantiate a new instance of Handler.
func New(usersService *service.Users) Handler {
	return Handler{
		usersService: usersService,
	}
}

// HealthCheck will response with OK status if server is listening for incoming requests.
// Will replace with Swagger soon.
func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

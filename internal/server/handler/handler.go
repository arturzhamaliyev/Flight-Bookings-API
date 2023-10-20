package handler

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for this service.
type Handler struct {
	usersService Users
}

// New will instantiate a new instance of Handler.
func New(usersService *service.Users) Handler {
	return Handler{
		usersService: usersService,
	}
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

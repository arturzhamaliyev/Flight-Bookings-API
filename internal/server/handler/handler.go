package handler

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for this service.
type Handler struct {
	users Users
}

// New will instantiate a new instance of Handler.
func New(services service.Service) Handler {
	return Handler{
		users: services.User,
	}
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

package handler

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// Handler represents a HTTP server handler that can handle requests for this service.
type Handler struct {
	Router *gin.Engine

	users Users
}

// New will instantiate a new instance of Handler.
func New(services service.Service) Handler {
	h := Handler{
		users: services.User,
	}

	r := gin.Default()

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/", h.CreateUser)
	}

	h.Router = r

	return h
}

func (h *Handler) HealthCheck(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusOK)
}

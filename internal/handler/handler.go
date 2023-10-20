package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/model"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

// Handler represents a HTTP server handler that can handle requests for this service.
type Handler struct {
	users Users
}

// New will instantiate a new instance of Server.
func New(cfg config.Config, services service.Service) *http.Server {
	h := &Handler{
		users: services.User,
	}

	r := gin.Default()

	r.GET("/health", h.healthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/", h.createUser)
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
}

func (h *Handler) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "OK")
}

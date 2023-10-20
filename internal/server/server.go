package server

import (
	"fmt"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
	"github.com/gin-gonic/gin"
)

// New will instantiate a new instance of Server.
func New(cfg config.Config, services service.Service) *http.Server {
	h := handler.New(services)

	r := gin.Default()

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/", h.CreateUser)
	}

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
}

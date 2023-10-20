package server

import (
	"fmt"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

// New will instantiate a new instance of Server.
func New(cfg config.Config, services service.Service) *http.Server {
	h := handler.New(services)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: h.Router,
	}
}

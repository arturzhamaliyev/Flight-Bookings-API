package server

import (
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
)

// Router represents a type that provides operations on serving HTTP.
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

// New will instantiate a new instance of Server.
func New(cfg config.Config, r Router) *http.Server {
	return &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}
}

package server

import (
	"fmt"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/gin-gonic/gin"
)

// New will instantiate a new instance of Server.
func New(cfg config.Config, r *gin.Engine) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: r,
	}
}

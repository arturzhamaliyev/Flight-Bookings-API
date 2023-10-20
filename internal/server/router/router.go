package router

import (
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/handler"
	"github.com/gin-gonic/gin"
)

// New will instantiate a new instance of Router.
func New(h handler.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/", h.CreateUser)
	}

	return r
}

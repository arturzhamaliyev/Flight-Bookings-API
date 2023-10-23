package router

import (
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/handler"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// New will instantiate a new instance of Router.
func New(h handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("", h.CreateUser)
	}

	return r
}

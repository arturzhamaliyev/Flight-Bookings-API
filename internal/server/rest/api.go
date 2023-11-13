package rest

import (
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler"
)

// New will instantiate a new instance of Router.
func New(h handler.Handler) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(zap.L(), time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(zap.L(), true))
	r.Use(cors.Default())

	r.GET("/health", h.HealthCheck)

	v1 := r.Group("/api/v1")

	users := v1.Group("/users")
	{
		users.POST("", h.CreateUser)
	}

	return r
}

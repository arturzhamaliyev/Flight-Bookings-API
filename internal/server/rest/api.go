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

	v1 := r.Group("/api/v1")

	v1.GET("/health", h.HealthCheck)
	v1.POST("/add-admin", cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:8080"},
		AllowMethods: []string{"POST"},
	}), h.AddAdmin)

	admin := v1.Group("admin", h.JWTAuthAdmin)
	{
		admin.PUT("/:userID/update-profile", h.UpdateProfileByID)
	}

	users := v1.Group("/users")
	{
		users.POST("/sign-up", h.SignUp)
		users.POST("/sign-in", h.SignIn)
		users.POST("/sign-out", h.SignOut)

		users.PUT("/update-profile", h.JWTAuthCustomer, h.UpdateProfile)
	}

	return r
}

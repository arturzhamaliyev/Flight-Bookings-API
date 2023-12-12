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
		AllowOrigins: []string{"http://localhost"},
		AllowMethods: []string{"POST"},
	}), h.AddAdmin)

	admin := v1.Group("/admin", h.JWTAuthAdmin)
	{
		users := admin.Group("/users")
		{
			profile := users.Group("/profile/:id")
			{
				profile.GET("/", h.GetProfileByID)
				profile.PUT("/update", h.UpdateProfileByID)
				profile.DELETE("/delete", h.DeleteProfileByID)
			}
		}
	}

	users := v1.Group("/users")
	{
		users.POST("/sign-up", h.SignUp)
		users.POST("/sign-in", h.SignIn)
		users.POST("/sign-out", h.SignOut)

		profile := users.Group("/profile", h.JWTAuthCustomer)
		{
			profile.GET("/", h.GetProfile)
			profile.PUT("/update", h.UpdateProfile)
			profile.DELETE("/delete", h.DeleteProfile)
		}
	}

	return r
}

package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

// DB represents a type that can be used to interact with the database.
type DB interface {
	PingContext(ctx context.Context) error
}

// Server represents a HTTP server that can handle requests for this service.
type Server struct {
	logger *zap.SugaredLogger
	users  Users
	db     DB
}

// New will instantiate a new instance of Server.
func New(cfg *config.Config, logger *zap.SugaredLogger, users Users, db DB) *http.Server {
	s := &Server{
		logger: logger,
		users:  users,
		db:     db,
	}

	r := gin.Default()

	r.GET("/health", s.healthCheck)

	v1 := r.Group("/v1")

	user := v1.Group("/users")
	{
		user.POST("/", s.createUser)
	}

	return &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           r,
		ReadHeaderTimeout: 60 * time.Second,
	}
}

func (s *Server) healthCheck(ctx *gin.Context) {
	if err := s.db.PingContext(ctx); err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}

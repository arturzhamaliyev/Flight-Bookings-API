package http

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/model"
	"github.com/gin-gonic/gin"
)

// Users represents a type that provides operations on users.
type Users interface {
	CreateUser(ctx context.Context, user model.User) error
}

// DB represents a type that can be used to interact with the database.
type DB interface {
	PingContext(context.Context) error
}

// Server represents a HTTP server that can handle requests for this service.
type Server struct {
	users Users
	db    DB
}

// New will instantiate a new instance of Server.
func New(users Users, db DB) *Server {
	return &Server{
		users: users,
		db:    db,
	}
}

// AddRoutes will add the routes this server supports to the router.
func (s *Server) AddRoutes(r *gin.Engine) {
	r.GET("/health", s.healthCheck)

	v1 := r.Group("/v1")

	users := v1.Group("/users")
	{
		users.POST("/", s.createUser)
	}
}

func (s *Server) healthCheck(ctx *gin.Context) {
	if err := s.db.PingContext(ctx); err != nil {
		handleError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}

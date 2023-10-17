package http

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	"github.com/gin-gonic/gin"
)

// ErrServer is the error returned when the server stops due to an error.
const ErrServer = errors.Error("listen stopped with error")

const (
	readHeaderTimeout = 60 * time.Second
)

// Config represents the configuration of the http listener.
type Config struct {
	Port string `yaml:"port"`
}

type Service interface {
	AddRoutes(r *gin.Engine)
}

type Server struct {
	server *http.Server
	port   string
}

// New instantiates a new instance of Server.
func New(s Service, cfg Config) *Server {
	r := gin.Default()

	s.AddRoutes(r)

	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf(":%s", cfg.Port),
			BaseContext: func(net.Listener) context.Context {
				baseContext := context.Background()
				return logging.With(baseContext, logging.From(baseContext))
			},
			Handler:           r,
			ReadHeaderTimeout: readHeaderTimeout,
		},
		port: cfg.Port,
	}
}

// Listen starts the server and listens on the configured port.
func (s *Server) Listen(ctx context.Context) error {
	logging.From(ctx).Info(fmt.Sprintf("http server starting on port: %s", s.port))

	err := s.server.ListenAndServe()
	if err != nil {
		return ErrServer.Wrap(err)
	}

	logging.From(ctx).Info("http server stopped")
	return nil
}

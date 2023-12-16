package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/db/driver"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/repository"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/rest/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/swagger"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"
)

func main() {
	// Create new Logger instance with default production logging configuration.
	loggerDefault, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	zap.ReplaceGlobals(loggerDefault)
	defer func() {
		err := loggerDefault.Sync()
		if err != nil {
			log.Println(err)
		}
	}()
	logger := loggerDefault.Sugar()

	logger.Info("app starting...")

	// Load configuration from config/config.yaml which contains details such as DB connection params
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to the postgres DB
	db, err := driver.ConnectPostgres(ctx, logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			logger.Info(err)
		}
	}()

	// Instantiate and connect all our classes
	usersRepo := repository.NewUsersRepo(db)
	flightsRepo := repository.NewAirportsRepo(db)
	usersService := service.NewUsersService(usersRepo)
	flightsService := service.NewFlightsService(cfg, flightsRepo)
	handler := handler.New(usersService, flightsService)
	router := rest.New(handler)
	s := server.New(cfg, router)
	go func() {
		swagger.New(cfg)
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		logger.Info("shutdown start")

		ctx, cancel = context.WithTimeout(ctx, time.Minute)
		defer cancel()
		err := s.Shutdown(ctx)
		if err != nil {
			logger.Infof("failed to shutdown: %v", err)
		}

		logger.Info("shutdown end")
	}()

	// Start listening for HTTP requests
	logger.Infof("server listening on port: %v", cfg.Server)
	err = s.ListenAndServe()
	if err != nil {
		logger.Infof("failed to serve on port: %v due to: %v", cfg.Server.Port, err)
	}
}

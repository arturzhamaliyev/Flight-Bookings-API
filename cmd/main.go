package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/router"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/server/handler"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/service"

	"github.com/jmoiron/sqlx"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/repository"
	"go.uber.org/zap"

	_ "github.com/jackc/pgx/v5/stdlib"
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
	db, err := initDatabase(ctx, logger, cfg)
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
	usersService := service.NewUsersService(usersRepo)
	handler := handler.New(usersService)
	router := router.New(handler)
	s := server.New(cfg, router)

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
	err = s.ListenAndServe()
	if err != nil {
		logger.Infof("failed to serve on port: %v due to: %v", cfg.Port, err)
	}
}

func initDatabase(ctx context.Context, logger *zap.SugaredLogger, cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.DBAddr)
	if err != nil {
		logger.Infof("failed connect to db: %v", err)
		return nil, err
	}
	return db, nil
}

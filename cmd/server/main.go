package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"

	"github.com/jmoiron/sqlx"

	httptransport "github.com/arturzhamaliyev/Flight-Bookings-API/internal/transport/http"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/store"
	"go.uber.org/zap"

	_ "github.com/lib/pq"
)

func main() {
	// Create new Logger instance with default production logging configuration.
	logg, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}
	logger := logg.Sugar()
	defer logger.Sync()

	logger.Info("app starting...")

	// Load configuration from config/config.yaml which contains details such as DB connection params
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-c
		cancel()
	}()

	// Connect to the postgres DB
	db, err := initDatabase(ctx, logger, cfg)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()

	// Instantiate and connect all our classes
	store := store.New(db)
	users := users.New(store)
	server := httptransport.New(cfg, logger, users, db)

	go func() {
		// Start listening for HTTP requests
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatalf("failed to serve on port: %v", cfg.Port)
		}
	}()

	<-ctx.Done()
	logger.Info("gracefull shutdown started")

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)
	if err != nil {
		logger.Fatalf("failed to shutdown", err)
	}
}

func initDatabase(ctx context.Context, logger *zap.SugaredLogger, cfg *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DBAddr)
	if err != nil {
		logger.Infof("failed connect to db: %v", err)
		return nil, err
	}
	return db, nil
}

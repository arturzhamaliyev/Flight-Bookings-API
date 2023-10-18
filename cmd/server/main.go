package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/app"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/drivers/psql"
	"github.com/gin-gonic/gin"

	// "github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/listeners/http"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	httptransport "github.com/arturzhamaliyev/Flight-Bookings-API/internal/transport/http"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/store"
	"go.uber.org/zap"
)

func main() {
	app.Start(func(ctx context.Context, a *app.App) (func(), error) {
		// Load configuration from config/config.yaml which contains details such as DB connection params
		cfg, err := config.Load(ctx)
		if err != nil {
			return nil, err
		}

		// Connect to the postgres DB
		db, err := initDatabase(ctx, cfg, a)
		if err != nil {
			return nil, err
		}

		// Run our migrations which will update the DB or create it if it doesn't exist
		err = db.MigratePostgres(ctx, "file://migrations")
		if err != nil {
			return nil, err
		}
		a.OnShutdown(func() {
			// Temp for development so database is cleared on shutdown
			if err := db.RevertMigrations(ctx, "file://migrations"); err != nil {
				logging.From(ctx).Error("failed to revert migrations", zap.Error(err))
			}
		})

		// Instantiate and connect all our classes
		store := store.New(db.GetDB())
		users := users.New(store)
		s := httptransport.New(users, db.GetDB())

		r := gin.Default()
		s.AddRoutes(r)

		server := &http.Server{
			Addr:              ":" + cfg.Port,
			Handler:           r,
			ReadHeaderTimeout: 60 * time.Second,
		}

		// Start listening for HTTP requests
		return func() {
			err := server.ListenAndServe()
			if err != nil {
				logging.From(ctx).Error(fmt.Sprintf("failed to serve on port: %v", cfg.Port), zap.Error(err))
			}
		}, nil
	})
}

func initDatabase(ctx context.Context, cfg *config.Config, app *app.App) (*psql.Driver, error) {
	db := psql.New(cfg)

	err := db.Connect(ctx)
	if err != nil {
		logging.From(ctx).Error("failed connect to db", zap.Error(err))
		return nil, err
	}

	app.OnShutdown(func() {
		// Shutdown connection when server terminated
		logging.From(ctx).Info("shutting down db connection")
		if err := db.Close(ctx); err != nil {
			logging.From(ctx).Error("failed to close db connection", zap.Error(err))
		}
	})

	return db, nil
}

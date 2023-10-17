package main

import (
	"context"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/app"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/drivers/psql"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/listeners/http"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	httptransport "github.com/arturzhamaliyev/Flight-Bookings-API/internal/transport/http"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/users/store"
	"github.com/cenkalti/backoff/v4"
	"go.uber.org/zap"
)

func main() {
	app.Start(func(ctx context.Context, a *app.App) ([]app.Listener, error) {
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

		// Create a HTTP server
		h := http.New(s, cfg.HTTP)

		// Start listening for HTTP requests
		return []app.Listener{
			h,
		}, nil
	})
}

func initDatabase(ctx context.Context, cfg *config.Config, app *app.App) (*psql.Driver, error) {
	db := psql.New(cfg.PSQL)

	err := backoff.Retry(func() error {
		return db.Connect(ctx)
	}, backoff.NewExponentialBackOff())
	if err != nil {
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

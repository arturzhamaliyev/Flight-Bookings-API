package driver

import (
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/platform/config"
)

// ConnectPostgres connects to postgres database with pgx driver.
func ConnectPostgres(ctx context.Context, logger *zap.SugaredLogger, cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", cfg.Postgres.DatabaseURL)
	if err != nil {
		logger.Infof("failed connect to db: %v", err)
		return nil, err
	}
	return db, nil
}

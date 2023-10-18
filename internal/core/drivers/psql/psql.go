package psql

import (
	"context"
	"fmt"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/jmoiron/sqlx"
)

const (
	// ErrConnect is returned when we cannot connect to the database.
	ErrConnect = errors.Error("failed to connect to postgres db")
	// ErrClose is returned when we cannot close the database.
	ErrClose = errors.Error("failed to close postgres db connection")
)

type Driver struct {
	cfg *config.Config
	db  *sqlx.DB
}

// New instantiates an instance of the Driver.
func New(cfg *config.Config) *Driver {
	return &Driver{
		cfg: cfg,
	}
}

// Connect connects to the database.
func (d *Driver) Connect(ctx context.Context) error {
	fmt.Println(d.cfg)
	db, err := sqlx.Connect("postgres", d.cfg.DB)
	if err != nil {
		return ErrConnect.Wrap(err)
	}

	d.db = db

	return nil
}

// Close closes the database connection.
func (d *Driver) Close(ctx context.Context) error {
	if err := d.db.Close(); err != nil {
		return ErrClose.Wrap(err)
	}
	return nil
}

// GetDB returns the underlying database connection.
func (d *Driver) GetDB() *sqlx.DB {
	return d.db
}

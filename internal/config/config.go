package config

import (
	"errors"

	e "github.com/arturzhamaliyev/Flight-Bookings-API/internal/errors"
	"github.com/kelseyhightower/envconfig"
)

// ErrProcess is returned when the environment variables can not be processed.
var ErrProcess = errors.New("failed to process variables")

// Config represents the configuration of our application.
type Config struct {
	Port   int    `default:"8080"`
	DBAddr string `envconfig:"DB_ADDR"`
}

// Load loads the configuration from the file.
func Load() (*Config, error) {
	var cfg Config

	err := envconfig.Process("cfg", &cfg)
	if err != nil {
		return nil, e.Wrap(ErrProcess, err)
	}

	return &cfg, nil
}

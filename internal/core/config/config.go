package config

import (
	"os"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/drivers/psql"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/listeners/http"
	"gopkg.in/yaml.v3"
)

const (
	// ErrRead is returned when we cannot read the config file.
	ErrRead = errors.Error("failed to read file")
	// ErrUnmarshal is returned when we cannot unmarshal the config file.
	ErrUnmarshal = errors.Error("failed to unmarshal file")
)

// AppConfig represents the configuration of our application.
type AppConfig struct {
	HTTP http.Config `yaml:"http"`
	PSQL psql.Config `yaml:"psql"`
}

// Load loads the configuration from a yaml file on disk.
func Load(cfg any) error {
	data, err := os.ReadFile("config/config.yaml")
	if err != nil {
		return ErrRead.Wrap(err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return ErrUnmarshal.Wrap(err)
	}

	return nil
}

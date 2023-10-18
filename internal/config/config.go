package config

import (
	"context"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/spf13/viper"
)

const (
	// ErrRead is returned when the configuration file cannot be read.
	ErrRead = errors.Error("failed to read file")
	// ErrUnmarshal is returned when the configuration file cannot be unmarshalled.
	ErrUnmarshal = errors.Error("failed to unmarshal file")
)

const (
	configName = "config"
	configType = "json"
	configPath = "./config"
)

// Config represents the configuration of our application.
type Config struct {
	Port string `json:"port"`
	DB   string `json:"db"`
}

// Load loads the configuration from the file.
func Load(ctx context.Context) (*Config, error) {
	cfg := new(Config)

	viper.SetConfigName(configName)
	viper.SetConfigType(configType)
	viper.AddConfigPath(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, ErrRead.Wrap(err)
	}

	err = viper.Unmarshal(cfg)
	if err != nil {
		return nil, ErrUnmarshal.Wrap(err)
	}

	return cfg, nil
}

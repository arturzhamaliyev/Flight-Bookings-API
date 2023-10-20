package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the configuration of our application.
type Config struct {
	Port   int    `default:"8080"`
	DBAddr string `split_words:"true"`
}

// Load loads the configuration from the file.
func Load() (Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to process variables: %w", err)
	}

	return cfg, nil
}

package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Config represents the configuration of our application.
type (
	Config struct {
		Server   Server
		Postgres Postgres
		Swagger  Swagger
	}

	Swagger struct {
		Port string `envconfig:"SWAGGER_PORT" default:"8081"`
	}

	Postgres struct {
		DatabaseURL string `envconfig:"DATABASE_URL"`
	}

	Server struct {
		Port string `envconfig:"SERVER_PORT" default:"8080"`
	}
)

// Load loads the configuration from the file.
func Load() (Config, error) {
	var cfg Config

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("failed to process variables: %w", err)
	}

	return cfg, nil
}

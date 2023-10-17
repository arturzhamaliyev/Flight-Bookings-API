package config

import (
	"context"
	"fmt"
	"os"

	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/config"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/errors"
	"github.com/arturzhamaliyev/Flight-Bookings-API/internal/core/logging"
	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)

const (
	// ErrInvalidEnvironment is returned when the SPEAKEASY_ENVIRONMENT environment variable is not set.
	ErrInvalidEnvironment = errors.Error("SPEAKEASY_ENVIRONMENT is not set")
	// ErrRead is returned when the configuration file cannot be read.
	ErrRead = errors.Error("failed to read file")
	// ErrUnmarshal is returned when the configuration file cannot be unmarshalled.
	ErrUnmarshal = errors.Error("failed to unmarshal file")
	// ErrEnvVars is returned when the environment variables are invalid.
	ErrEnvVars = errors.Error("failed parsing env vars")
	// ErrValidation is returned when the configuration is invalid.
	ErrValidation = errors.Error("invalid configuration")
)

var (
	baseConfigPath = "config/config.yaml"
	envConfigPath  = "config/config-%s.yaml"
)

// Config represents the configuration of our application.
type Config struct {
	config.AppConfig `yaml:",inline"`
}

// Load loads the configuration from the config/config.yaml file.
func Load(ctx context.Context) (*Config, error) {
	cfg := new(Config)

	err := loadFromFiles(ctx, cfg)
	if err != nil {
		return nil, err
	}

	err = env.Parse(cfg)
	if err != nil {
		return nil, ErrEnvVars.Wrap(err)
	}

	validate := validator.New()
	err = validate.Struct(cfg)
	if err != nil {
		return nil, ErrValidation.Wrap(err)
	}

	return cfg, nil
}

func loadFromFiles(ctx context.Context, cfg any) error {
	environ := os.Getenv("SPEAKEASY_ENVIRONMENT")
	if environ == "" {
		return ErrInvalidEnvironment
	}

	err := loadYaml(ctx, baseConfigPath, cfg)
	if err != nil {
		return err
	}

	p := fmt.Sprintf(envConfigPath, environ)

	_, err = os.Stat(p)
	if !errors.Is(err, os.ErrNotExist) {
		err = loadYaml(ctx, p, cfg)
		if err != nil {
			return err
		}
	}

	return nil
}

func loadYaml(ctx context.Context, filename string, cfg any) error {
	logging.From(ctx).Info("Loading configuration", zap.String("path", filename))

	data, err := os.ReadFile(filename)
	if err != nil {
		return ErrRead.Wrap(err)
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return ErrUnmarshal.Wrap(err)
	}

	return nil
}

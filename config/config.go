package config

import (
	"fmt"
	"os"

	"github.com/nishujangra/balancerx/models"
	"gopkg.in/yaml.v2"
)

func LoadConfig(path string) (*models.Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("cannot read config: %w", err)
	}

	var cfg models.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("cannot parse config: %w", err)
	}

	// Default port if not set
	if cfg.Port == "" {
		cfg.Port = "8080"
	}

	// Default strategy if not set
	if cfg.Strategy == "" {
		cfg.Strategy = "round-robin"
	}

	// Default protocol if not set
	if cfg.Protocol == "" {
		cfg.Protocol = "http"
	}

	// Default health check path if not set
	if cfg.HealthCheck.Path == "" {
		cfg.HealthCheck.Path = "/health"
	}

	return &cfg, nil
}

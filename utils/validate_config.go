package utils

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/nishujangra/balancerx/models"
)

func ValidateConfig(cfg *models.Config) (bool, error) {
	// Validate port
	if cfg.Port == "" {
		return false, errors.New("port is required")
	}
	if _, err := strconv.Atoi(cfg.Port); err != nil {
		return false, errors.New("port must be a number")
	}

	// Validate strategy
	validStrategies := map[string]bool{"round-robin": true, "random": true}
	if !validStrategies[cfg.Strategy] {
		return false, errors.New("invalid strategy: must be 'round-robin' or 'random'")
	}

	// Validate backends
	if len(cfg.Backends) == 0 {
		return false, errors.New("at least one backend is required")
	}
	for _, backend := range cfg.Backends {
		if !isValidBackend(backend) {
			return false, errors.New("invalid backend: " + backend)
		}
	}

	// Validate protocol
	if cfg.Protocol != "http" && cfg.Protocol != "https" {
		return false, errors.New("protocol must be 'http' or 'https'")
	}

	// Validate health check interval
	if cfg.HealthCheck.Interval != "" {
		if _, err := time.ParseDuration(cfg.HealthCheck.Interval); err != nil {
			return false, errors.New("invalid health_check.interval: must be a valid duration string (e.g., '10s')")
		}
	}

	// Validate health check path
	if cfg.HealthCheck.Path == "" || !strings.HasPrefix(cfg.HealthCheck.Path, "/") {
		return false, errors.New("invalid health_check.path: must start with '/'")
	}

	return true, nil
}

// isValidBackend checks if the backend is a valid host (IP or domain) with optional port
func isValidBackend(backend string) bool {
	if strings.HasPrefix(backend, "http://") || strings.HasPrefix(backend, "https://") {
		u, err := url.Parse(backend)
		return err == nil && u.Host != ""
	}

	// Try IP:Port or hostname:Port
	host, _, err := net.SplitHostPort(backend)
	if err != nil {
		return false
	}
	return net.ParseIP(host) != nil || isDomainName(host)
}

// isDomainName checks if a string could be a valid domain name
func isDomainName(s string) bool {
	return strings.Contains(s, ".")
}

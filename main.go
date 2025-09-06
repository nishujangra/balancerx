package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"github.com/nishujangra/balancerx/balancer"
	"github.com/nishujangra/balancerx/config"
	"github.com/nishujangra/balancerx/models"
	"github.com/nishujangra/balancerx/proxies"
	"github.com/nishujangra/balancerx/utils"
)

func main() {
	// CLI
	configPath := flag.String("config", "/etc/balancerx/config.yaml", "Path to configuration file")
	flag.Parse()

	// Try to find config file with fallback logic
	var cfg *models.Config
	var err error

	// If user specified a config path, use it directly
	if *configPath != "/etc/balancerx/config.yaml" {
		cfg, err = config.LoadConfig(*configPath)
		if err != nil {
			log.Fatalf("Failed to load config from %s: %v", *configPath, err)
		}
	} else {
		// Try multiple config locations in order of preference
		configLocations := []string{
			"/etc/balancerx/config.yaml", // System-wide config
			"./config.yaml",              // Local config (for development)
			"config.yaml",                // Current directory
		}

		for _, location := range configLocations {
			if _, err := os.Stat(location); err == nil {
				cfg, err = config.LoadConfig(location)
				if err == nil {
					log.Printf("Loaded config from: %s", location)
					break
				}
			}
		}

		if cfg == nil {
			log.Fatalf("Failed to load config from any of the following locations: %v", configLocations)
		}
	}

	if ok, err := utils.ValidateConfig(cfg); !ok {
		log.Fatalf("Invalid config: %v", err)
	}

	// Set up logging to file
	// Try multiple log locations in order of preference
	logLocations := []string{
		"/var/log/balancerx/balancerx.log", // System-wide log
		"log/balancerx.log",                // Local log (for development)
		"balancerx.log",                    // Current directory
	}

	var logFile *os.File
	for _, location := range logLocations {
		// Create directory if it doesn't exist
		dir := filepath.Dir(location)
		if err := os.MkdirAll(dir, 0755); err == nil {
			logFile, err = os.OpenFile(location, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err == nil {
				log.Printf("Logging to: %s", location)
				break
			}
		}
	}

	if logFile == nil {
		log.Fatalf("Failed to open log file in any of the following locations: %v", logLocations)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Load strategy
	lb, err := balancer.New(cfg.Strategy, cfg.Backends, cfg.HealthCheck)
	if err != nil {
		log.Fatalf("Load balancer strategy error: %v", err)
	}

	log.Printf("Running BalancerX on port %s using '%s' strategy", cfg.Port, cfg.Strategy)

	var mu sync.RWMutex

	switch cfg.Protocol {
	case "http":
		handler := proxies.NewHTTPProxy(&mu, cfg, lb)
		log.Printf("Starting HTTP proxy on :%s", cfg.Port)
		log.Fatal(http.ListenAndServe(":"+cfg.Port, handler))

	case "tcp":
		tcpProxy := proxies.NewTCPProxy(&mu, cfg, lb)
		log.Printf("Starting TCP proxy on :%s", cfg.Port)
		log.Fatal(tcpProxy.Start())

	default:
		log.Fatalf("Unsupported protocol: %s", cfg.Protocol)
	}
}

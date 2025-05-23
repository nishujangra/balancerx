package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/nishujangra/balancerx/balancer"
	"github.com/nishujangra/balancerx/config"
	"github.com/nishujangra/balancerx/proxies"
)

func main() {
	// CLI
	configPath := flag.String("config", "config.yaml", "Path to configuration file")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Set up logging to file
	logFile, err := os.OpenFile("log/balancerx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)

	// Load strategy
	lb, err := balancer.New(cfg.Strategy, cfg.Backends)
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

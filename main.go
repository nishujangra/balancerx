package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/nishujangra/balancerx/balancer"
	"github.com/nishujangra/balancerx/config"
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

	// Proxy handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		target := lb.Next()

		if target == "" {
			log.Println("❌ No healthy backend available")
			http.Error(w, "No healthy backends available", http.StatusServiceUnavailable)
			return
		}

		remote, err := url.Parse(target)
		if err != nil {
			log.Printf("[ERROR] [%s] Invalid backend URL: %s", start.Format(time.RFC3339), target)
			http.Error(w, "Invalid backend", http.StatusInternalServerError)
			return
		}

		log.Printf("[FORWARD] [%s] %s %s -> %s", start.Format(time.RFC3339), r.Method, r.URL.Path, target)

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.ServeHTTP(w, r)

		duration := time.Since(start)
		log.Printf("Forwarded to %s in %s", target, duration)
	})

	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Fatal(http.ListenAndServe(addr, nil))
}

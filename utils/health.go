package utils

import (
	"net"
	"net/http"
	"strings"
	"time"
)

func IsBackendAlive(url string, health_check_path string) bool {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return checkHTTP(url, health_check_path)
	}
	return checkTCP(url)
}

func checkHTTP(url string, health_check_path string) bool {
	client := &http.Client{Timeout: 1 * time.Second}

	resp, err := client.Get(url + health_check_path)
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == 200
}

func checkTCP(address string) bool {
	conn, err := net.DialTimeout("tcp", address, 1*time.Second)

	if err != nil {
		return false
	}
	conn.Close()

	return true
}

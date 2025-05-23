package utils

import (
	"net/http"
	"time"
)

func IsBackendAlive(url string) bool {
	client := http.Client{
		Timeout: 500 * time.Millisecond,
	}

	resp, err := client.Get(url + "/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK
}

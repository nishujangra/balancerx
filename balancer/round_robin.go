package balancer

import (
	"net/http"
	"sync/atomic"
	"time"
)

type RoundRobin struct {
	backends []string
	current  uint64
}

func NewRoundRobin(backends []string) *RoundRobin {
	return &RoundRobin{backends: backends}
}

func isBackendAlive(url string) bool {
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

func (rr *RoundRobin) Next() string {
	total := len(rr.backends)

	for i := 0; i < total; i++ {
		index := atomic.AddUint64(&rr.current, 1)
		backend := rr.backends[index%uint64(total)]

		if isBackendAlive(backend) {
			return backend
		}
	}

	// Fallback: no healthy backends found
	return ""
}

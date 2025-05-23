package balancer

import (
	"sync/atomic"

	"github.com/nishujangra/balancerx/utils"
)

type RoundRobin struct {
	backends []string
	current  uint64
}

func NewRoundRobin(backends []string) *RoundRobin {
	return &RoundRobin{backends: backends}
}

func (rr *RoundRobin) Next() string {
	total := len(rr.backends)

	for i := 0; i < total; i++ {
		index := atomic.AddUint64(&rr.current, 1)
		backend := rr.backends[index%uint64(total)]

		if utils.IsBackendAlive(backend) {
			return backend
		}
	}

	// Fallback: no healthy backends found
	return ""
}

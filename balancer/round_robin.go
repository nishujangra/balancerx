package balancer

import (
	"sync/atomic"

	"github.com/nishujangra/balancerx/utils"
)

type RoundRobin struct {
	Backends []string
	Current  uint64
}

func NewRoundRobin(backends []string) *RoundRobin {
	return &RoundRobin{Backends: backends}
}

func (rr *RoundRobin) Next() string {
	total := len(rr.Backends)

	for i := 0; i < total; i++ {
		index := atomic.AddUint64(&rr.Current, 1)
		backend := rr.Backends[index%uint64(total)]

		if utils.IsBackendAlive(backend) {
			return backend
		}
		return backend
	}

	// Fallback: no healthy backends found
	return ""
}

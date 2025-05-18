package balancer

import "sync/atomic"

type RoundRobin struct {
	backends []string
	current  uint64
}

func NewRoundRobin(backends []string) *RoundRobin {
	return &RoundRobin{backends: backends}
}

func (rr *RoundRobin) Next() string {
	i := atomic.AddUint64(&rr.current, 1)
	return rr.backends[i%uint64(len(rr.backends))]
}

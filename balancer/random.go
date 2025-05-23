package balancer

import (
	"math/rand"
	"time"

	"github.com/nishujangra/balancerx/utils"
)

type Random struct {
	Backends []string
}

func NewRandom(backends []string) *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{Backends: backends}
}

func (r *Random) Next() string {
	perm := rand.Perm(len(r.Backends)) // shuffle order
	for _, i := range perm {
		backend := r.Backends[i]
		if utils.IsBackendAlive(backend) {
			return backend
		}
	}
	// fallback: no healthy backend
	return ""
}

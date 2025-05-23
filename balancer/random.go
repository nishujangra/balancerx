package balancer

import (
	"math/rand"
	"time"

	"github.com/nishujangra/balancerx/utils"
)

type Random struct {
	backends []string
}

func NewRandom(backends []string) *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{backends: backends}
}

func (r *Random) Next() string {
	perm := rand.Perm(len(r.backends)) // shuffle order
	for _, i := range perm {
		backend := r.backends[i]
		if utils.IsBackendAlive(backend) {
			return backend
		}
	}
	// fallback: no healthy backend
	return ""
}

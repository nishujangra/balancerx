package balancer

import (
	"math/rand"
	"time"

	"github.com/nishujangra/balancerx/models"
	"github.com/nishujangra/balancerx/utils"
)

type Random struct {
	Backends    []string
	HealthCheck models.HealthCheck
}

func NewRandom(backends []string, hc models.HealthCheck) *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{
		Backends:    backends,
		HealthCheck: hc,
	}
}

func (r *Random) Next() string {
	perm := rand.Perm(len(r.Backends)) // shuffle order
	for _, i := range perm {
		backend := r.Backends[i]
		if utils.IsBackendAlive(backend, r.HealthCheck.Path) {
			return backend
		}
	}
	// fallback: no healthy backend
	return ""
}

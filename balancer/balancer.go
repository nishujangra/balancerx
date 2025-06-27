package balancer

import (
	"errors"

	"github.com/nishujangra/balancerx/models"
)

type Strategy models.LoadBalancingStrategy

// Factory
func New(strategy string, backends []string, hc models.HealthCheck) (Strategy, error) {
	switch strategy {
	case "round-robin":
		return NewRoundRobin(backends, hc), nil
	case "random":
		return NewRandom(backends, hc), nil
	default:
		return nil, errors.New("unsupported strategy: " + strategy)
	}
}

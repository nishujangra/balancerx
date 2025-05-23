package balancer

import (
	"errors"

	"github.com/nishujangra/balancerx/models"
)

type Strategy models.LoadBalancingStrategy

// Factory
func New(strategy string, backends []string) (Strategy, error) {
	switch strategy {
	case "round-robin":
		return NewRoundRobin(backends), nil
	case "random":
		return NewRandom(backends), nil
	default:
		return nil, errors.New("unsupported strategy: " + strategy)
	}
}

package balancer

import (
	"errors"
)

// Interface for strategies
type Strategy interface {
	Next() string
}

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

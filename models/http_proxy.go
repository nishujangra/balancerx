package models

import "sync"

type HTTPProxy struct {
	Mu  *sync.RWMutex
	Cfg *Config
	LB  LoadBalancingStrategy
}

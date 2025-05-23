package models

import "sync"

type TCPProxy struct {
	Mu  *sync.RWMutex
	Cfg *Config
	LB  LoadBalancingStrategy
}

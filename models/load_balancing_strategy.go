package models

type LoadBalancingStrategy interface {
	Next() string
}

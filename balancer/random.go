package balancer

import (
	"math/rand"
	"time"
)

type Random struct {
	backends []string
}

func NewRandom(backends []string) *Random {
	rand.Seed(time.Now().UnixNano())
	return &Random{backends: backends}
}

func (r *Random) Next() string {
	i := rand.Intn(len(r.backends))
	return r.backends[i]
}

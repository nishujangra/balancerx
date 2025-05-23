package proxies

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"time"

	"github.com/nishujangra/balancerx/models"
)

type HTTPProxy models.HTTPProxy

func NewHTTPProxy(mu *sync.RWMutex, cfg *models.Config, lb models.LoadBalancingStrategy) *HTTPProxy {
	return &HTTPProxy{
		Mu:  mu,
		Cfg: cfg,
		LB:  lb,
	}
}

func (p *HTTPProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	p.Mu.RLock()
	target := p.LB.Next()
	p.Mu.RUnlock()

	remote, err := url.Parse(target)
	if err != nil {
		log.Printf("[ERROR] Invalid backend: %s", target)
		http.Error(w, "Invalid backend", http.StatusInternalServerError)
		return
	}

	r.Header.Set("X-Forwarded-For", r.RemoteAddr)

	proxy := httputil.NewSingleHostReverseProxy(remote)

	proxy.ModifyResponse = func(resp *http.Response) error {
		latency := time.Since(start)
		log.Printf("[RESPONSE] %s -> %d (%s)", target, resp.StatusCode, latency)
		return nil
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, req *http.Request, e error) {
		log.Printf("[FAILED] %s -> %v", target, e)
		http.Error(w, "Backend error", http.StatusBadGateway)
	}

	proxy.ServeHTTP(w, r)
}

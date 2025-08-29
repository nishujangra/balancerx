package proxies

import (
	"io"
	"log"
	"net"
	"sync"

	"github.com/nishujangra/balancerx/models"
)

type TCPProxy models.TCPProxy

func NewTCPProxy(mu *sync.RWMutex, cfg *models.Config, lb models.LoadBalancingStrategy) *TCPProxy {
	return &TCPProxy{
		Mu:  mu,
		Cfg: cfg,
		LB:  lb,
	}
}

func (p *TCPProxy) Start() error {
	listener, err := net.Listen("tcp", ":"+p.Cfg.Port)
	if err != nil {
		return err
	}

	defer listener.Close()
	log.Printf("[TCP] Listening on :%s", p.Cfg.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("[TCP] Accept error: %v", err)
			continue
		}

		go p.handle(conn)
	}
}

func (p *TCPProxy) handle(clientConn net.Conn) {
	defer clientConn.Close()

	p.Mu.RLock()
	target := p.LB.Next()
	p.Mu.RUnlock()

	backendConn, err := net.Dial("tcp", target)
	if err != nil {
		log.Printf("[TCP] Connection failed to %s: %v", target, err)
		return
	}
	defer backendConn.Close()

	log.Printf("[TCP] Forwarding to %s", target)

	go io.Copy(backendConn, clientConn)
	io.Copy(clientConn, backendConn)
}

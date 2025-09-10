# BalancerX

**BalancerX** is a lightwright load balancer written in Go. It supports proxies in both HTTP and TCP, with multiple load balancing strategies, health checks -- all configurations controlled by `config.yaml`.

`/etc/balancerx/config.yaml` when installed with `.deb` package.

---

## ✨ Features

* 🔁 **Round-robin and Random** load balancing strategies
* 📂 **YAML-based configuration** for HTTP or TCP protocols
* 🩺 **Active health checks**: HTTP endpoint checks or TCP connection probes
* ⚡ **HTTP reverse proxy** using `net/http/httputil`
* 📜 **Logging**: at `/var/log/balancerx/balancerx.log`
* 🔧 **Easily extendable** with new strategies and protocol support
* 🚀 **Health checker service** (available but not yet integrated)

---

## 🚀 Quick Start

### Installation

#### Debian Package (Recommended)
```bash
# Download and install
wget https://github.com/nishujangra/balancerx/releases/latest/download/balancerx_1.0.0.deb
sudo dpkg -i balancerx_1.0.0.deb

# Start the service
sudo systemctl start balancerx
sudo systemctl enable balancerx
```

#### From Source
```bash
git clone https://github.com/nishujangra/balancerx.git
cd balancerx
go build -o build/balancerx main.go
./build/balancerx -config=config.yaml
```

### Basic Configuration

Create a `config.yaml` file:

```yaml
port: 8080
protocol: http
strategy: round-robin # or random
backends:
  - http://localhost:9001
  - http://localhost:9002
  - http://localhost:9003
health_check:
  path: /health
```

---

## 📊 Performance

BalancerX demonstrates excellent performance characteristics:

| Load Balancer | Binary Size | Memory Usage (RSS) | Package Files |
|---------------|-------------|-------------------|---------------|
| **BalancerX** | 9.4M        | 2,468 KB          | 18 files      |
| nginx         | 1.3M        | Not running       | 12 files      |

### Key Benefits

* **Low Memory Footprint**: Only ~2.5MB RAM usage
* **Fewer Dependencies** → Leaner binary & fewer attack surfaces
* **Native Go HTTP** → Predictable, no framework bloat
* **Smart Goroutine Management** → Scalable concurrency without memory blowups

---

## 🔌 Supported Protocols

| Protocol | Description |
|----------|-------------|
| `http`   | Reverse proxy for HTTP with health checks and header handling |
| `tcp`    | Transparent forwarding of raw TCP connections with basic health checks |

---

## ⚙️ Load Balancing Strategies

| Name          | Description |
|---------------|-------------|
| `round-robin` | Cycles through backends in fixed order |
| `random`      | Randomly selects a backend per request |
| `least-conn`  | *(Planned)* Chooses backend with fewest connections |
| `ip-hash`     | *(Planned)* Sticky routing by client IP hash |

---

## 🧪 Testing

### HTTP Backends

Sample backend server using golang is provided in `dummy-server/dummy-golang.go`. Copy that sample backend server if you want to try `BalancerX` in your local-system

or You can use your own server but that server must have health check route you provided and if not provided in `config.yaml` then `/health` is the default endpoint for health check

```sh
go run dummy-golang.go 9001
go run dummy-golang.go 9002
```

or 

```sh
# For running dummy serves from port 9000 to 9100
for i in $(seq 9000 9100); do
    go run dummy-golang.py $i &
done

# Stop running servers from port 9000 to 9100
for port in $(seq 9000 9100); do
    pid=$(lsof -t -i:$port)
    if [ -n "$pid" ]; then
        kill -9 $pid
    fi
done

```

### TCP Backends
```bash
# Start TCP servers
nc -lk 6001 &
nc -lk 6002 &

# Connect via BalancerX
telnet localhost 9090
```

---

## 📚 Documentation

- [Installation Guide](getting-started/installation.md) - Detailed installation instructions
- [Configuration](getting-started/configuration.md) - Complete configuration reference
- [Load Balancing Strategies](user-guide/strategies.md) - Available strategies and usage
- [Health Checks](user-guide/health-checks.md) - Health monitoring configuration
- [Performance Benchmarks](performance/benchmarks.md) - Detailed performance analysis

## 🤝 Contributing

Contributions and suggestions are welcome! See our [Contributing Guide](development/contributing.md) for details.

## 📜 License

AGPLv3 © 2025 Nishant

---

**BalancerX** is developed and maintained by [Nishant](https://github.com/nishujangra).

# BalancerX

**BalancerX** is a lightweight, high-performance load balancer written in Go. It supports both HTTP and TCP proxying, multiple load balancing strategies, active health checks, and flexible configuration — all controlled via a simple `config.yaml` file.

---

## ✨ Features

* 🔁 Round-robin and 🎲 Random load balancing strategies
* 📂 YAML-based configuration for HTTP or TCP protocols
* 🩺 Active health checks: HTTP endpoint checks or TCP connection probes
* 🪵 Request and connection logging to files (with easy console extension)
* ⚡ HTTP reverse proxy using `net/http/httputil`
* 🔧 Easily extendable with new strategies and protocol support

---

## 🚀 Getting Started

### 1. Clone the Project

```bash
git clone https://github.com/nishujangra/balancerx.git
cd balancerx
```

### 2. Create Your `config.yaml`

Example for HTTP mode:

```yaml
port: 8080
protocol: http          # "http" or "tcp"
strategy: round-robin   # Load balancing strategy: "round-robin" or "random"
backends:
  - http://localhost:9001
  - http://localhost:9002
  - http://localhost:9003
health_check:
  interval: 10s         # Health check interval (HTTP only)
  path: /health         # Recommended for backend reliability (HTTP only)
```

Example for TCP mode:

```yaml
port: 9090
protocol: tcp
strategy: random
backends:
  - localhost:6001
  - localhost:6002
```

➡️ For full configuration details and advanced options, see [docs/config.md](docs/config.md).

---

## 🏃 Run BalancerX

```bash
go run main.go -config=config.yaml
```

If `-config` is omitted, it defaults to `config.yaml`.

---

## 🧪 Testing the Load Balancer

### HTTP Backends Example (Using Python)

```bash
# Terminal 1
python3 -m http.server 9001

# Terminal 2
python3 -m http.server 9002

# Terminal 3
python3 -m http.server 9003
```

Test load balancing:

```bash
curl http://localhost:8080
```

### TCP Backends Example

Run TCP echo servers or other services on the ports you configured:

```bash
# Example with netcat
nc -lk 6001
nc -lk 6002
```

Connect via BalancerX's listening port (e.g., `telnet localhost 9090`).

---

Got it 🚀 — you already have a strong README, so let’s extend it with a **Benchmarks section** that shows real test results and highlights BalancerX’s performance. I’ll keep it professional, resume-friendly, and easy for others to replicate.

Here’s the updated README snippet with **Benchmarks** added 👇

---

## 📊 Benchmarks

For benchmarks go to [benchmarks/](./benchmarks/README.md)

---

This way:  
- Shows **measurable proof** of BalancerX’s performance.  
- Demonstrates **stability under stress** (important for recruiters/engineers reading).  
- Leaves room for **future improvements** (shows you think about scaling).

---

## ⚙️ Supported Load Balancing Strategies

| Name          | Description                                   |
| ------------- | --------------------------------------------- |
| `round-robin` | Cycles through backends in fixed order        |
| `random`      | Randomly selects a backend per request        |
| `least-conn`  | *(Planned)* Chooses backend with fewest conns |
| `ip-hash`     | *(Planned)* Sticky routing by client IP hash  |

More strategies are planned and easy to integrate.

---

## 🔌 Supported Protocols

| Protocol | Description                                                            |
| -------- | ---------------------------------------------------------------------- |
| `http`   | Reverse proxy for HTTP with health checks and header handling          |
| `tcp`    | Transparent forwarding of raw TCP connections with basic health checks |

---

## 📄 Logging

BalancerX logs connection and forwarding details to `log/balancerx.log`:

```
2025/05/24 02:30:45 [FORWARD] GET / -> http://localhost:9003
2025/05/24 02:46:33 [TCP] Forwarding to localhost:6001
2025/05/24 02:46:39 [TCP] Connection failed to localhost:6002: connect: connection refused
```

🔧 Logs can be easily extended to output to both file and console.

---

## 🗂 Folder Structure

```
balancerx/
├── main.go
├── config/
│   ├── config.go           # Config loader
├── balancer/
│   ├── balancer.go         # Base interface
│   ├── round_robin.go
│   └── random.go
├── proxies/
│   ├── http_proxy.go       # HTTP proxy logic
│   └── tcp_proxy.go        # TCP proxy logic
├── log/
│   └── balancerx.log
├── docs/
│   ├── config.md           # Full configuration guide
│   └── protocols.md        # Protocol handling details
├── config.yaml
└── README.md
```

---

## 📌 Project Roadmap

* [x] HTTP & TCP proxy support
* [x] Round-robin & random strategies
* [x] Active health checks (HTTP & TCP)
* [ ] Admin API to expose backend status
* [ ] Least-connections & IP-hash strategies
* [ ] Dockerfile for container deployment

---

## 📜 License

MIT License © 2025 Nishant

---

## 🤝 Contributing

Contributions and suggestions are welcome!

```bash
git checkout -b feature/your-feature
git commit -m "Add your feature"
git push origin feature/your-feature
```

Open a PR or raise an issue on [GitHub Issues](https://github.com/nishujangra/balancerx/issues).

---

## 👨‍💻 Author

**BalancerX** is developed and maintained by [Nishant](https://github.com/nishujangra).
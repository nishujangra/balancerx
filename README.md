# BalancerX

**BalancerX** is a lightweight, high-performance load balancer written in Go. It supports HTTP and TCP proxying, multiple load balancing strategies, and health checks for backend services. All you need to do is configure the `config.yaml` file.

---

## ✨ Features

* 🔁 Round-robin and 🎲 Random strategies
* 📂 YAML-based configuration with support for `http` and `tcp` protocols
* 🩺 Active health checks for backend availability (HTTP path or TCP dial)
* 🪵 Request and connection logging with optional file output
* ⚡ HTTP reverse proxy support using `net/http/httputil`
* 🔧 Easy to extend with new strategies and protocols

---

## 🚀 Getting Started

### 1. Clone the project

```bash
git clone https://github.com/nishujangra/balancerx.git
cd balancerx
```

### 2. Prepare a `config.yaml`

Create a file named `config.yaml` in the project root:

```yaml
port: 8080
protocol: http          # "http" or "tcp"
strategy: round-robin   # or "random"
backends:
  - http://localhost:9001
  - http://localhost:9002
  - http://localhost:9003
health_check:
  interval: 10s
  path: /health         # Used only in HTTP mode
```

> ✅ Run actual backend servers at the listed URLs or host\:ports (for TCP).

---

## 🏃 Run It

```bash
go run main.go -config=config.yaml
```

If `-config` is omitted, it defaults to `config.yaml`.

---

## 🧪 Try It

For HTTP backend testing, start dummy servers (e.g., Python):

```bash
# Terminal 1
python3 -m http.server 9001

# Terminal 2
python3 -m http.server 9002

# Terminal 3
python3 -m http.server 9003
```

Then:

```bash
curl http://localhost:8080
```

BalancerX will forward requests to backends based on the selected strategy.

For TCP, run services on configured ports and connect through BalancerX’s listening port.

---

## ⚙️ Supported Strategies

| Name          | Description                                    |
| ------------- | ---------------------------------------------- |
| `round-robin` | Cycles through backends in order               |
| `random`      | Chooses a backend at random for each call      |
| least-conn    | (Planned) Pick backend with fewest connections |
| ip-hash       | (Planned) Route clients by IP hash             |

> Additional strategies like `least-connections`, `ip-hash`, and others can be added easily.

---

## 🔌 Supported Protocols

| Protocol | Description                                                                                |
| -------- | ------------------------------------------------------------------------------------------ |
| `http`   | Acts as an HTTP reverse proxy with HTTP health checks and request/response handling        |
| `tcp`    | Forwards raw TCP connections; uses TCP dial health checks; no HTTP inspection or rewriting |

---

## 📄 Logs

BalancerX writes logs to a file named `balancerx.log` inside the `/log` directory:

```
2025/05/24 02:30:45 [FORWARD] [2025-05-24T02:30:45+05:30] GET / -> http://localhost:9003
2025/05/24 02:30:45 Forwarded to http://localhost:9003 in 1.709607ms

2025/05/24 02:46:29 [TCP] Connection failed to localhost:9003: dial tcp 127.0.0.1:9003: connect: connection refused
2025/05/24 02:46:33 [TCP] Forwarding to localhost:9001
2025/05/24 02:46:39 [TCP] Forwarding to localhost:9002
```

> Modify logging in `main.go` to log to both file and console if desired.

---

## 🛠 Folder Structure

```
balancerx/
├── main.go
├── config/
│   └── config.go
├── balancer/
│   ├── balancer.go
│   ├── round_robin.go
│   └── random.go
├── proxies/
│   ├── http_proxy.go        # HTTP proxy implementation
│   └── tcp_proxy.go         # TCP proxy implementation
├── log/
│   └── balancerx.log
├── docs/                    # Documentations
├── config.yaml
└── README.md
```

---

## 📌 TODO

* [x] Add health check system
* [x] Support TCP and HTTP proxy protocols
* [ ] Add admin API to show backend status
* [ ] Add support for IP-hash and least-connections strategies
* [ ] Dockerfile for containerized deployment

---

## 📜 License

MIT © 2025 Nishant

---

## 🤝 Contributing

Contributions, issues, and feature requests are welcome!

```bash
git checkout -b feature/your-feature
git commit -m "Add your feature"
git push origin feature/your-feature
```

Please open a pull request or discussion in the [GitHub Issues](https://github.com/nishujangra/balancerx/issues) page.

---

## 👨‍💻 Author

**BalancerX** is created and maintained by Nishant.

Follow updates and new features via [GitHub](https://github.com/nishujangra/balancerx)
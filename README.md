# BalancerX

**BalancerX** is a lightweight, high-performance load balancer written in Go. It supports HTTP-based reverse proxying, simple load balancing strategies, and health checks for backend services. All the user has to change is the `config.yaml` file.

---

## ✨ Features

- 🔁 Round-robin and 🎲 Random strategies
- 📂 YAML-based configuration
- 🪵 Request logging with optional file output
- 🔧 Easy to extend with new strategies
- ⚡ Reverse proxy support using `net/http/httputil`

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
strategy: round-robin  # or "random"
backends:
  - http://localhost:9001
  - http://localhost:9002
  - http://localhost:9003
health_check:
  interval: 10s
  path: /health
```

> ✅ You must run actual backend servers at the listed URLs (e.g., with Python or Go).

---

## 🏃 Run It

```bash
go run main.go -config=config.yaml
```

If `-config` is omitted, it defaults to `config.yaml`.

---

## 🧪 Try It

Start some dummy backend servers (e.g., using Python):

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

BalancerX will forward the request to one of the backends, based on the configured strategy.

---

## ⚙️ Supported Strategies

| Name          | Description                               |
| ------------- | ----------------------------------------- |
| `round-robin` | Cycles through backends in order          |
| `random`      | Chooses a backend at random for each call |
| least-conn  | (Planned) Pick backend with fewest connections |
| ip-hash     | (Planned) Route clients by IP hash             |

> More strategies like `least-connections`, `ip-hash`, etc., can be added easily.

---

## 📄 Logs

BalancerX writes logs to a file named `balancerx.log` in the `/log` directory:

```
[FORWARD] [2025-05-19T10:12:32Z] GET / -> http://localhost:9001
[RESPONSE] [2025-05-19T10:12:32Z] http://localhost:9001 -> 200
[FAILED] [2025-05-19T10:12:33Z] http://localhost:9002 -> dial tcp: connection refused
```

> You can modify `main.go` to log to both file and console if desired.

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
├── config.yaml
├── balancerx.log
└── README.md
```

---

## 📌 TODO

* [ ] Add health check system
* [ ] Add admin API to show backend status
* [ ] Add support for IP-hash and least-connections
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

**BalancerX** is created and maintained by [Nishant].

Follow updates and new features via [GitHub](https://github.com/nishujangra/balancerx)
# Configuration Guide

This document explains the `config.yaml` used by **BalancerX** to control behavior, supported protocols, load balancing strategies, and health check configuration.

---

## 🗂️ Overview

The configuration file defines:

✅ The port for incoming traffic
✅ Protocol type: HTTP or TCP
✅ Load balancing strategy
✅ Backend servers
✅ Health check settings (HTTP only)

---

## 📄 Example `config.yaml` for HTTP

```yaml
port: 8080
protocol: http            # "http" or "tcp"
strategy: round-robin      # Load balancing strategy: round-robin, random, etc.
backends:
  - http://localhost:9001
  - http://localhost:9002
health_check:
  interval: 10s            # Interval between health checks (optional, HTTP only)
  path: /health            # Recommended health check path (optional, default "/health")
```

---

## 📄 Example `config.yaml` for TCP

```yaml
port: 9090
protocol: tcp
strategy: round-robin
backends:
  - localhost:6001
  - localhost:6002
```

⚠️ **Note:** TCP backends do not support HTTP health checks. BalancerX performs basic TCP connection checks.

---

## 🔧 Configuration Fields

| Field          | Type       | Required | Description                                                          |
| -------------- | ---------- | -------- | -------------------------------------------------------------------- |
| `port`         | int/string | ✅ Yes    | Port number BalancerX listens on (e.g., `8080`)                      |
| `protocol`     | string     | ✅ Yes    | Proxy mode: `"http"` (HTTP reverse proxy) or `"tcp"` (raw TCP proxy) |
| `strategy`     | string     | ✅ Yes    | Load balancing strategy: `round-robin`, `random`, (more planned)     |
| `backends`     | list       | ✅ Yes    | List of backend servers. Format depends on `protocol`.               |
| `health_check` | map        | Optional | HTTP-only. Periodic health checks. Contains `interval` and `path`.   |

---

## 🎛️ Field Details

### `port`

* TCP port to listen for incoming connections.
* Accepts numeric values like `8080`, `9090`.

### `protocol`

* `"http"` for HTTP reverse proxy mode (supports health checks).
* `"tcp"` for raw TCP proxy mode (health checks are TCP-level only).

### `strategy`

* Controls how backends are selected:

| Strategy      | Description                        | Additional Requirements |
| ------------- | ---------------------------------- | ----------------------- |
| `round-robin` | Cycle through backends in sequence | None                    |
| `random`      | Random backend selection           | None                    |

*Planned:*

* `least-conn`: Fewest active connections
* `ip-hash`: Sticky sessions per client IP

---

### `backends`

| Protocol | Backend Format Example                |
| -------- | ------------------------------------- |
| `http`   | `http://localhost:9001` (scheme req.) |
| `tcp`    | `host:port` (e.g., `localhost:6001`)  |

* For **HTTP**, include full URL with `http://` scheme.
* For **TCP**, only `host:port` format is valid.

---

### `health_check` (HTTP-only)

* Optional block for active backend health monitoring.
* Ignored in TCP mode (TCP uses basic connection checks).

| Field      | Type   | Recommended     | Description                                     |
| ---------- | ------ | ----------- | ----------------------------------------------- |
| `interval` | string | `"10s"`     | Time between health checks (e.g., `5s`, `30s`)  |
| `path`     | string | `"/health"` | HTTP path to check; recommended for reliability |

Example:

```yaml
health_check:
  interval: 15s
  path: /health
```

**Notes:**

* Using `/health` or a simple 200-OK endpoint is recommended but not required.
* Unhealthy backends are skipped until they pass health checks again.

---

## ⚠️ Validation Rules

| Field          | Required | Notes                                                        |
| -------------- | -------- | ------------------------------------------------------------ |
| `port`         | Yes      | Must be available and not used by another service            |
| `protocol`     | Yes      | `"http"` or `"tcp"` only                                     |
| `strategy`     | Yes      | Must be one of: `round-robin`, `random`                      |
| `backends`     | Yes      | At least one backend, formatted appropriately per `protocol` |
| `health_check` | Optional | Allowed only with `protocol: http`; ignored for TCP          |

---

## 🛠 Tips for Reliability

✅ Prefer using `/health` endpoints for HTTP backends
✅ For TCP, ensure backend ports accept connections before adding to the list
✅ Do not mix HTTP and TCP backends under one config — define per protocol

---

## 📢 Questions or Contributions

* Open an issue on [GitHub](https://github.com/nishujangra/balancerx/issues) for questions or feature suggestions.
* Pull requests for new strategies or protocol improvements are welcome.
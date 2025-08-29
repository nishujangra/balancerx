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
  interval: 10s            # Health check interval (currently unused)
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
| `health_check` | map        | Optional | HTTP-only. Health check configuration with `interval` and `path`.    |

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

* **Current Implementation**: Health checks are performed on every request when selecting a backend.
* **Future Enhancement**: A background health checker service is implemented but not yet integrated.
* Ignored in TCP mode (TCP uses basic connection checks).

| Field      | Type   | Current Status | Description                                     |
| ---------- | ------ | -------------- | ----------------------------------------------- |
| `interval` | string | ⚠️ **Unused**  | Time between health checks (e.g., `5s`, `30s`)  |
| `path`     | string | ✅ **Active**   | HTTP path to check; recommended for reliability |

Example:

```yaml
health_check:
  interval: 15s
  path: /health
```

**Current Behavior:**

* **Per-request Health Checks**: Health status is checked every time a backend is selected
* **Real-time Validation**: Always gets current health status
* **Performance Impact**: Each request includes a health check
* **Reliability**: Ensures only healthy backends receive traffic

**Health Check Logs:**

```
[FORWARD] GET / -> http://localhost:9003
[TCP] Forwarding to localhost:6001
[TCP] Connection failed to localhost:6002: connect: connection refused
```

**Notes:**

* Using `/health` or a simple 200-OK endpoint is recommended but not required.
* Unhealthy backends are automatically skipped until they pass health checks again.
* The `interval` field is currently unused but reserved for future background health checking.

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

✅ **Prefer using `/health` endpoints** for HTTP backends
✅ **For TCP**, ensure backend ports accept connections before adding to the list
✅ **Do not mix HTTP and TCP backends** under one config — define per protocol
✅ **Monitor logs** to identify backend issues early

---

## 🔮 Future Health Checking

A background health checker service has been implemented and will provide:

* **Periodic Health Checks**: Run health checks on configurable intervals
* **Background Processing**: No impact on request performance
* **Status Caching**: Maintain health status between requests
* **Status Logging**: Clear logs when backends become healthy/unhealthy

**Planned Integration:**

* Connect health checker to load balancing strategies
* Use cached health status instead of per-request checks
* Implement graceful shutdown with health checker cleanup

---

## 📢 Questions or Contributions

* Open an issue on [GitHub](https://github.com/nishujangra/balancerx/issues) for questions or feature suggestions.
* Pull requests for new strategies or protocol improvements are welcome.
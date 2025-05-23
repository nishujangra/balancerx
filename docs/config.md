# Configuration Guide

This document explains the configuration file `config.yaml` used by **BalancerX** to control the load balancer's behavior.

---

## Overview

The configuration file is written in YAML format and lets you define:

* The port your load balancer listens on
* Protocol type (HTTP or TCP)
* Load balancing strategy
* Backend servers to forward traffic to
* Health check settings (for HTTP backends)

---

## Example `config.yaml`

```yaml
port: 8080
protocol: http          # "http" or "tcp"
strategy: round-robin   # Load balancing strategy: round-robin, random, etc.
backends:
  - http://localhost:9001
  - http://localhost:9002
health_check:
  interval: 10s        # Interval between health checks (HTTP only)
  path: /health        # HTTP path used for health checks
```

---

## Configuration Fields

| Field          | Type       | Required | Description                                                                                                                                  |
| -------------- | ---------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `port`         | string/int | Yes      | The TCP port your BalancerX proxy listens on (e.g., `8080` or `"8080"`)                                                                      |
| `protocol`     | string     | Yes      | Protocol used by the backends. `"http"` for HTTP reverse proxy, `"tcp"` for raw TCP proxy                                                    |
| `strategy`     | string     | Yes      | Load balancing strategy. Supported: `round-robin`, `random` (more planned)                                                                   |
| `backends`     | list       | Yes      | List of backend servers. Must be URLs with protocol for HTTP (e.g., `http://localhost:9001`), or host\:port for TCP (e.g., `localhost:6001`) |
| `health_check` | map        | No       | Settings for active health checks. Only supported for HTTP backends                                                                          |

---

## Details per Field

### `port`

* The port on which BalancerX listens for incoming requests.
* Example: `8080`

### `protocol`

* Determines the proxy type.
* `"http"` means BalancerX acts as a reverse proxy for HTTP services.
* `"tcp"` means BalancerX forwards raw TCP connections.

### `strategy`

* Defines how BalancerX selects a backend server.
* Supported strategies:

  * `round-robin`: cycles through backends evenly
  * `random`: picks a backend randomly

### `backends`

* List of backend server addresses.
* For HTTP backends, **include the scheme**, e.g., `http://localhost:9001`
* For TCP backends, use only `host:port`, e.g., `localhost:6001`

### `health_check`

* Optional; only for HTTP backends.
* Controls periodic health checks to backend servers.
* Fields:

  * `interval`: Duration between health checks (e.g., `10s`, `30s`)
  * `path`: HTTP path to check (usually `/health`)

---

## Example for TCP protocol

```yaml
port: 9090
protocol: tcp
strategy: round-robin
backends:
  - localhost:6001
  - localhost:6002
```

**Note:** TCP backends do not support HTTP health checks; BalancerX will perform simple TCP connection tests instead.

---

## Tips

* Ensure backend servers are reachable and responding on the specified addresses.
* Use `/health` endpoints on HTTP backends to improve reliability.
* Adjust `health_check.interval` based on backend responsiveness and network conditions.

---

If you have questions or want to contribute improvements to the config system, please open an issue or pull request on the [GitHub repository](https://github.com/nishujangra/balancerx).

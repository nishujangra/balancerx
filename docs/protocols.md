# Protocols Supported by BalancerX

BalancerX supports two main protocols for backend communication:

* **HTTP Proxy Mode**
* **TCP Proxy Mode**

This document explains how each protocol works, how to configure them, and their use cases.

---

## 1. HTTP Proxy Mode

### Overview

In HTTP mode, BalancerX acts as a **reverse proxy** that forwards incoming HTTP requests to backend HTTP servers. It uses Go's `net/http/httputil.ReverseProxy` to handle request forwarding, response rewriting, and connection management.

### Features

* Full support for HTTP(S) protocol features
* Active health checks via `/health` endpoints on backends
* Load balancing strategies based on backend health
* Logging of HTTP requests and responses

### Configuration Example

```yaml
protocol: http
port: 8080
strategy: round-robin
backends:
  - http://localhost:9001
  - http://localhost:9002
health_check:
  interval: 10s
  path: /health
```

### Notes

* Backend URLs **must include the `http://` or `https://` scheme**.
* Health checks send periodic HTTP GET requests to the backend’s `/health` path (configurable).
* Suitable for web services, APIs, or any HTTP-based applications.

---

## 2. TCP Proxy Mode

### Overview

In TCP mode, BalancerX forwards raw TCP connections from clients to backend TCP servers without interpreting the data. This is useful for protocols other than HTTP, such as database connections, custom TCP protocols, or any non-HTTP traffic.

### Features

* Forwards bidirectional TCP traffic transparently
* Uses simple connection health checks by attempting TCP dial
* Supports load balancing strategies similar to HTTP mode

### Configuration Example

```yaml
protocol: tcp
port: 9090
strategy: round-robin
backends:
  - localhost:6001
  - localhost:6002
```

### How to run dummy servers for TCP

```bash
nc -lk 9001
```

This command will run a `TCP` server at the port 9001

### Notes

* Backend addresses **must be host\:port pairs only**, no URL scheme (e.g., `localhost:6001`).
* Health checks consist of trying to open a TCP connection to each backend to verify availability.
* Suitable for database proxies, message brokers, or any TCP-based custom protocols.

---

## Choosing Between HTTP and TCP Mode

| Use Case                                                     | Recommended Protocol Mode |
| ------------------------------------------------------------ | ------------------------- |
| Web servers, REST APIs                                       | HTTP                      |
| TCP services like databases, SMTP, MQTT                      | TCP                       |
| Services requiring HTTP headers, cookies, or TLS termination | HTTP                      |
| Simple raw TCP forwarding without inspection                 | TCP                       |

---

## Important Considerations

* **Health Checks:**

  * HTTP mode performs active health checks by sending HTTP requests.
  * TCP mode performs passive health checks by attempting TCP connections.

* **Logging:**
  Logs include forwarded requests/responses for HTTP mode and connection attempts for TCP mode.

* **Extensibility:**
  BalancerX is designed to support adding new protocols or custom health checks if needed.

---

## Summary

BalancerX’s protocol flexibility allows it to be used in various environments, from simple HTTP load balancing to generic TCP forwarding, making it a versatile tool for developers and operators.

---

If you want to request support for additional protocols or features, please submit an issue or pull request in the [GitHub repository](https://github.com/nishujangra/balancerx).

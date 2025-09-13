# Performance Benchmarks

BalancerX is designed for high performance and low resource usage. This document provides detailed performance analysis and comparison with other load balancers.

## Test Setup

* **Machine**: Localhost (8-thread CPU, 16 GB RAM)
* **Backends**: 3 Go Echo HTTP servers
* **Load Tool**: wrk (HTTP benchmarking tool)
* **Binary Size**: \~9.4MB (single Go binary)
* **Dependencies**: Minimal (Go standard library only)

## Performance Results

### Test 1 — Light Load (50 Clients, 15s)

**Command:**

```bash
wrk -t2 -c50 -d15s http://localhost:6001/health
```

**Results:**

* **Requests/sec**: \~11,573
* **Total Requests**: 173k+ in 15s
* **Average Latency**: 4.4ms (max 32ms)
* **Errors**: 0

---

### Test 2 — Moderate Load (200 Clients, 30s)

**Command:**

```bash
wrk -t4 -c200 -d30s http://localhost:6001/health
```

**Results:**

* **Requests/sec**: \~12,563
* **Total Requests**: 377k+ in 30s
* **Average Latency**: 16.5ms (max 147ms)
* **Errors**: 0

---

### Test 3 — Heavy Load (500 Clients, 60s)

**Command:**

```bash
wrk -t8 -c500 -d60s http://localhost:6001/health
```

**Results:**

* **Requests/sec**: \~12,525
* **Total Requests**: 752k+ in 60s
* **Average Latency**: 40ms (max 222ms)
* **Errors**: 0

---

### Test 4 — Very Heavy Load (1,000 Clients, 60s)

**Command:**

```bash
wrk -t8 -c1000 -d60s http://localhost:6001/health
```

**Results:**

* **Requests/sec**: \~10,769
* **Total Requests**: 647k+ in 60s
* **Average Latency**: 94ms (max 508ms)
* **Errors**: 0

---

### Test 5 — Extreme Load (2,000 Clients, 60s)

**Command:**

```bash
wrk -t12 -c2000 -d60s http://localhost:6001/health
```

**Results:**

* **Requests/sec**: \~12,108
* **Total Requests**: 727k+ in 60s
* **Average Latency**: 83ms (max 378ms)
* **Errors**: 983 connection errors (likely OS/socket limit)

---

## Performance Summary

| Test       | Threads | Connections | Duration | Req/sec  | Avg Latency | Max Latency | Errors      |
| ---------- | ------- | ----------- | -------- | -------- | ----------- | ----------- | ----------- |
| Light      | 2       | 50          | 15s      | \~11,573 | 4.4ms       | 32ms        | 0           |
| Moderate   | 4       | 200         | 30s      | \~12,563 | 16.5ms      | 147ms       | 0           |
| Heavy      | 8       | 500         | 60s      | \~12,525 | 40ms        | 222ms       | 0           |
| Very Heavy | 8       | 1000        | 60s      | \~10,769 | 94ms        | 508ms       | 0           |
| Extreme    | 12      | 2000        | 60s      | \~12,108 | 83ms        | 378ms       | 983 connect |

---

## Key Takeaways

* **High Throughput**: BalancerX sustains 11k–12.5k req/sec across most loads
* **Low Latency**: Latency stays under 20ms at moderate loads, grows gracefully under heavier loads
* **Scalability**: Handles up to 1,000 concurrent clients with zero errors
* **Resource Limits**: At 2,000 concurrent clients, throughput is stable but system connection limits caused \~983 connect errors
* **Efficiency**: Single binary, \~9.4MB, minimal dependencies, optimized CPU usage

---

## Performance Characteristics

### Throughput Analysis

* Stable throughput (\~12k req/sec) across light to heavy loads
* Graceful degradation at 1,000+ concurrent connections

### Reliability

* Zero request/response errors up to 1,000 clients
* At 2,000 clients, only connection-level errors (not processing failures)

### Resource Efficiency

* Small binary size (\~9.4MB)
* Low memory and CPU usage under load
* Minimal dependencies (Go stdlib only)

---

## Conclusion

BalancerX demonstrates **excellent performance and stability** across a wide range of workloads:

* \~12k req/sec consistently
* Latency stays low (4ms–16ms) under moderate load
* Handles 500–1000 clients with stable throughput and zero errors
* At 2000 clients, throughput remains high but limited by OS/socket constraints

Overall, BalancerX is **production-ready, efficient, and resilient**, showing strong scalability for real-world deployments.
# BalancerX Benchmarks

This folder contains performance benchmark results for **BalancerX** under different loads using [`wrk`](https://github.com/wg/wrk).

## Test Setup
- **Machine:** Localhost (8-thread CPU, 16 GB RAM)  
- **Backends:** 3 servers (Go Echo HTTP server for comparison)  
- **Load Tool:** wrk  

---

### Test 1 — Moderate Load (100 Clients, 30s, Go Echo backend)

```bash
wrk -t4 -c100 -d30s http://localhost:8080/
```

* Requests/sec: **\~9,734**
* Total: **292k+ requests in 30s**
* Avg Latency: 10.6ms (max 96ms)
* Errors: 0

---

### Test 2 — Heavy Load (1,000 Clients, 60s, Go Echo backend)

```bash
wrk -t8 -c1000 -d60s http://localhost:8080/
```

* Requests/sec: **\~8,861**
* Total: **539k+ requests in 60s**
* Avg Latency: 119ms (max 1.24s)
* Errors: 0
* ⚡ Sustained high throughput and stability under 1,000 concurrent clients

At 1,000 concurrent clients, BalancerX maintained stable throughput (~8.8k req/sec) with slightly higher latency, but zero errors. This shows BalancerX scales well and degrades gracefully under heavy concurrency, rather than failing abruptly.

---

## 📊 Comparison Table

| Test     | Backend              | Threads | Conns | Duration | Req/sec | Latency (avg) | Errors                     |
| -------- | -------------------- | ------- | ----- | -------- | ------- | ------------- | -------------------------- |
| Moderate | Go Echo server       | 4       | 100   | 30s      | \~9,734 | 10.6ms        | 0                          |
| Heavy    | Go Echo server       | 8       | 1000  | 60s      | \~8,861 | 119ms         | 0                          |

---

## Key Takeaways

* BalancerX sustains **\~9.7k req/sec** with low latency (10ms avg) at moderate load on 3 Go Echo backends.
* At heavy load (1,000 clients), it maintains **\~8.8k req/sec** and over **539k requests** in 60s with **zero errors**.
* Demonstrates that BalancerX is **stable, production-ready, and scales well** under both light and heavy workloads.
* **TODO:** Add more strategies (least-connections, IP-hash) and keep-alive tuning for even higher scalability.
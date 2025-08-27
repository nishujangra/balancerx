# BalancerX Benchmarks

This folder contains performance benchmark results for **BalancerX** under different loads using [`wrk`](https://github.com/wg/wrk).

## Test Setup
- **Machine:** Localhost (8-thread CPU, 16 GB RAM)  
- **Backends:** 3 servers (`http.server` / lightweight HTTP echo server)  
- **Load Tool:** wrk  

---

### Test 1 — Moderate Load (100 Clients, 30s)

```bash
wrk -t4 -c100 -d30s http://localhost:8080/
```

* Requests/sec: **\~2,350**
* Total: **70k+ requests in 30s**
* Avg Latency: 173ms
* Errors: 0 socket errors
* ✅ Stable load distribution across 3 backends

---

### Test 2 — Heavy Load (1,000 Clients, 60s)

```bash
wrk -t8 -c1000 -d60s http://localhost:8080/
```

* Requests/sec: **\~656**
* Total: **39k+ requests in 60s**
* Avg Latency: 480ms (spikes up to 2s)
* Errors: \~15k timeouts, \~21k non-2xx (due to backend bottlenecks)
* ⚠️ Graceful degradation under overload — no BalancerX crash.

---

## Key Takeaways

* Handles **2k+ req/sec** reliably with moderate concurrency.
* Maintains stability and connection integrity under stress.
* **TODO:** Future optimizations (keep-alives, async I/O, least-connections) will further increase scalability.
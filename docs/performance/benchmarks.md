# Performance Benchmarks

BalancerX is designed for high performance and low resource usage. This document provides detailed performance analysis and comparison with other load balancers.

## Test Setup

- **Machine**: Localhost (8-thread CPU, 16 GB RAM)
- **Backends**: 3 Go Echo HTTP servers
- **Load Tool**: wrk (HTTP benchmarking tool)
- **Binary Size**: ~9.4MB (single Go binary)
- **Dependencies**: Minimal (Go standard library only)

## Performance Results

### Test 1 — Moderate Load (100 Clients, 30s)

**Command:**
```bash
wrk -t4 -c100 -d30s http://localhost:8080/
```

**Results:**
- **Requests/sec**: ~9,734
- **Total Requests**: 292k+ requests in 30s
- **Average Latency**: 10.6ms (max 96ms)
- **Errors**: 0

### Test 2 — Heavy Load (1,000 Clients, 60s)

**Command:**
```bash
wrk -t8 -c1000 -d60s http://localhost:8080/
```

**Results:**
- **Requests/sec**: ~8,861
- **Total Requests**: 539k+ requests in 60s
- **Average Latency**: 119ms (max 1.24s)
- **Errors**: 0
- **Stability**: Sustained high throughput and stability under 1,000 concurrent clients

## Performance Summary

| Test | Threads | Connections | Duration | Req/sec | Avg Latency | Max Latency | Errors |
|------|---------|-------------|----------|---------|-------------|-------------|--------|
| Moderate | 4 | 100 | 30s | ~9,734 | 10.6ms | 96ms | 0 |
| Heavy | 8 | 1,000 | 60s | ~8,861 | 119ms | 1.24s | 0 |

## Key Takeaways

- **High Throughput**: BalancerX sustains ~9.7k req/sec with low latency (10ms avg) at moderate load
- **Heavy Load Performance**: At 1,000 clients, maintains ~8.8k req/sec and over 539k requests in 60s with zero errors
- **Stability**: Demonstrates that BalancerX is stable, production-ready, and scales well under both light and heavy workloads
- **Zero Errors**: Both test scenarios achieved 0 errors, showing robust error handling
- **Graceful Degradation**: Under heavy load, BalancerX maintains stable throughput with slightly higher latency rather than failing abruptly

## Performance Characteristics

### Throughput Analysis
- **Moderate Load**: ~9,734 requests/second with 10.6ms average latency
- **Heavy Load**: ~8,861 requests/second with 119ms average latency
- **Scalability**: Handles 1,000 concurrent connections without errors
- **Reliability**: Zero errors across all test scenarios

### Resource Efficiency
- **Binary Size**: ~9.4MB single executable
- **Dependencies**: Minimal (Go standard library only)
- **Memory Usage**: Efficient memory management under load
- **CPU Usage**: Optimized for high-throughput scenarios

## Benchmarking Commands

### Running Performance Tests

#### Install wrk (HTTP benchmarking tool)

```bash
# Ubuntu/Debian
sudo apt install wrk

# macOS
brew install wrk

# Or build from source
git clone https://github.com/wg/wrk.git
cd wrk
make
```

#### Moderate Load Test (100 clients, 30 seconds)

```bash
wrk -t4 -c100 -d30s http://localhost:8080/
```

#### Heavy Load Test (1,000 clients, 60 seconds)

```bash
wrk -t8 -c1000 -d60s http://localhost:8080/
```

### Monitoring During Tests

```bash
# Monitor memory usage
ps -C balancerx -o pid,comm,rss

# Monitor CPU usage
top -p $(pgrep balancerx)

# Monitor network connections
ss -tuln | grep :8080

# Check binary size
ls -lh /usr/bin/balancerx
```


### Log Analysis

```bash
# Analyze request patterns
sudo journalctl -u balancerx | grep "FORWARD" | awk '{print $NF}' | sort | uniq -c

# Monitor response times
sudo journalctl -u balancerx | grep "RESPONSE" | tail -20
```

## Conclusion

BalancerX demonstrates excellent performance characteristics based on real-world testing:

- **High Throughput**: Sustains ~9.7k req/sec at moderate load, ~8.8k req/sec under heavy load
- **Low Latency**: 10.6ms average latency at moderate load
- **Zero Errors**: Robust error handling with 0 errors across all test scenarios
- **Scalability**: Handles 1,000 concurrent connections without failures
- **Stability**: Graceful degradation under heavy load rather than abrupt failures
- **Resource Efficiency**: ~9.4MB single binary with minimal dependencies

The benchmark results show that BalancerX is production-ready and scales well under both light and heavy workloads, making it an excellent choice for load balancing scenarios.

## Future Improvements

- Add more load balancing strategies (least-connections, IP-hash)
- Implement keep-alive tuning for even higher scalability
- Add more comprehensive performance testing scenarios
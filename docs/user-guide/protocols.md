# Protocols

BalancerX supports multiple protocols for load balancing. This document explains the supported protocols and their characteristics.

## Supported Protocols

### HTTP Protocol

The HTTP protocol provides reverse proxy functionality for HTTP/HTTPS traffic.

#### Features

- **Reverse Proxy**: Forwards HTTP requests to backend servers
- **Header Preservation**: Maintains original request headers
- **Health Checks**: HTTP endpoint-based health monitoring
- **Load Balancing**: Distributes requests across multiple backends
- **Error Handling**: Graceful handling of backend failures

#### Configuration

```yaml
port: 8080
protocol: http
strategy: round-robin
backends:
  - http://api1.example.com:8080
  - http://api2.example.com:8080
  - http://api3.example.com:8080
health_check:
  path: /health
```

#### HTTP Request Flow

```
1. Client sends HTTP request to BalancerX
2. BalancerX receives request on configured port
3. Load balancer selects backend using strategy
4. X-Forwarded-For header added with client IP
5. Request forwarded to selected backend
6. Response returned to client
7. Request/response logged
```

#### HTTP Headers

BalancerX handles headers as follows:

- **Original headers**: All client headers are automatically forwarded by the reverse proxy
- **X-Forwarded-For**: Automatically added with the client's IP address
- **Host header**: Automatically handled by the reverse proxy
- **Connection headers**: Automatically managed by the reverse proxy

**Note**: X-Real-IP header is not currently implemented in v1.0.0.

#### HTTP Health Checks

```yaml
health_check:
  path: /health         # Health check endpoint
```

**Health Check Process:**
1. Makes HTTP GET request to `backend + path`
2. Considers backend healthy if:
   - Connection succeeds
   - HTTP status code is 200
   - Response received within timeout

### TCP Protocol

The TCP protocol provides transparent TCP connection forwarding.

#### Features

- **Connection Forwarding**: Transparent TCP connection proxying
- **Bidirectional Data**: Forwards data in both directions
- **Health Checks**: TCP connection-based health monitoring
- **Load Balancing**: Distributes connections across backends
- **Connection Management**: Handles connection lifecycle

#### Configuration

```yaml
port: 9090
protocol: tcp
strategy: round-robin
backends:
  - db1.example.com:5432
  - db2.example.com:5432
  - db3.example.com:5432
```

#### TCP Connection Flow

```
1. Client connects to BalancerX
2. BalancerX accepts connection
3. Load balancer selects backend using strategy
4. Connection established to selected backend
5. Data proxied bidirectionally
6. Connection closed when done
```


## Use Cases

### HTTP Protocol Use Cases

#### Web APIs

```yaml
port: 8080
protocol: http
strategy: round-robin
backends:
  - http://api1.internal:8080
  - http://api2.internal:8080
  - http://api3.internal:8080
health_check:
  path: /api/health
```

#### Web Applications

```yaml
port: 80
protocol: http
strategy: round-robin
backends:
  - http://web1.internal:8080
  - http://web2.internal:8080
  - http://web3.internal:8080
health_check:
  path: /health
```

#### Microservices

```yaml
port: 8080
protocol: http
strategy: random
backends:
  - http://service1.internal:8080
  - http://service2.internal:8080
  - http://service3.internal:8080
health_check:
  path: /health
```

### TCP Protocol Use Cases

#### Database Load Balancing

```yaml
port: 5432
protocol: tcp
strategy: round-robin
backends:
  - db1.internal:5432
  - db2.internal:5432
  - db3.internal:5432
```

#### Redis Cluster

```yaml
port: 6379
protocol: tcp
strategy: round-robin
backends:
  - redis1.internal:6379
  - redis2.internal:6379
  - redis3.internal:6379
```

#### Custom TCP Services

```yaml
port: 9999
protocol: tcp
strategy: random
backends:
  - service1.internal:9999
  - service2.internal:9999
  - service3.internal:9999
```

## Protocol-Specific Configuration

#### Header Configuration

```yaml
protocol: http
backends:
  - http://api1.example.com:8080
  - http://api2.example.com:8080
# Headers are automatically handled
# X-Forwarded-For is automatically added
```

### TCP Configuration Options

#### Connection Settings

```yaml
protocol: tcp
backends:
  - db1.example.com:5432
  - db2.example.com:5432
```

**Note**: TCP protocol does not support health checks in v1.0.0. Connections are forwarded directly to backends.

## Troubleshooting

### HTTP Protocol Issues

#### Backend Not Responding

```bash
# Test backend directly
curl http://api1.example.com:8080/health

# Check health check endpoint
curl -I http://api1.example.com:8080/health

# Monitor BalancerX logs
sudo journalctl -u balancerx | grep "FORWARD"
```

#### Header Issues

```bash
# Check forwarded headers
curl -H "X-Test: value" http://localhost:8080

# Monitor header forwarding
sudo journalctl -u balancerx | grep "X-Forwarded-For"
```

### TCP Protocol Issues

#### Connection Failures

```bash
# Test backend connectivity
telnet db1.example.com 5432

# Check connection logs
sudo journalctl -u balancerx | grep "TCP"

# Monitor connection patterns
sudo journalctl -u balancerx | grep "Connection"
```

#### Data Forwarding Issues

```bash
# Test data forwarding
echo "test" | nc localhost 9090

# Monitor data flow
sudo journalctl -u balancerx | grep "Forwarding"
```
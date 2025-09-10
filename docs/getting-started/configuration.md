# Configuration

BalancerX uses a simple YAML configuration file to define all aspects of load balancing behavior.

## Configuration File Location

- **System-wide**: `/etc/balancerx/config.yaml` (default)
- **Custom path**: Specify with `-config` flag: `balancerx -config=/path/to/config.yaml`

## Basic Configuration Structure

```yaml
# Required fields
port: 8080                    # Port to listen on
protocol: http               # Protocol: "http" or "tcp"
strategy: round-robin        # Load balancing strategy
backends:                    # List of backend servers
  - http://localhost:9001
  - http://localhost:9002

# Optional fields
health_check:                # Health check configuration
  path: /health             # Health check endpoint (HTTP only)
```

## HTTP Configuration

### Basic HTTP Load Balancer

```yaml
port: 8080
protocol: http
strategy: round-robin
backends:
  - http://localhost:9001
  - http://localhost:9002
  - http://localhost:9003
health_check:
  path: /health
```

### HTTP with Custom Health Check

```yaml
port: 8080
protocol: http
strategy: random
backends:
  - http://api1.example.com:8080
  - http://api2.example.com:8080
  - http://api3.example.com:8080
health_check:
  path: /api/health
```

## TCP Configuration

### Basic TCP Load Balancer

```yaml
port: 9090
protocol: tcp
strategy: round-robin
backends:
  - localhost:6001
  - localhost:6002
  - localhost:6003
```

**Note**: TCP load balancing does not support health checks. Connections are forwarded directly to backends without health monitoring.


## Load Balancing Strategies

### Round-Robin

```yaml
strategy: round-robin
```

Distributes requests evenly across backends in a fixed order.

### Random

```yaml
strategy: random
```

Randomly selects a backend for each request.

## Health Check Configuration

### HTTP Health Checks

```yaml
health_check:
  path: /health          # Endpoint to check (required for HTTP)
```

**Note**: Health checks are only available for HTTP protocol. TCP load balancing forwards connections directly to backends without health monitoring.


## Complete Example

Here's a production-ready configuration:

```yaml
# Main configuration
port: 8080
protocol: http
strategy: round-robin

# Backend servers
backends:
  - http://api1.internal:8080
  - http://api2.internal:8080
  - http://api3.internal:8080

# Health monitoring
health_check:
  path: /api/health
```

## Configuration Validation

### Validate Configuration

```bash
# Check if config is valid
balancerx -config=config.yaml -validate

# Or test with dry run
balancerx -config=config.yaml -dry-run
```

### Common Validation Errors

**Invalid port:**
```yaml
port: 99999  # Error: Port must be between 1-65535
```

**Invalid protocol:**
```yaml
protocol: udp  # Error: Protocol must be "http" or "tcp"
```

**Invalid strategy:**
```yaml
strategy: least-conn  # Error: Strategy not implemented yet
```

**Invalid backend URL:**
```yaml
backends:
  - invalid-url  # Error: Invalid backend URL format
```

## Configuration Best Practices

### 1. Use Health Checks

Always configure health checks for HTTP load balancing in production:

```yaml
health_check:
  path: /health
```

**Note**: Health checks are only available for HTTP protocol. TCP load balancing forwards connections directly to backends without health monitoring.

### 2. Use Descriptive Backend URLs

```yaml
backends:
  - http://api-primary.internal:8080
  - http://api-secondary.internal:8080
  - http://api-backup.internal:8080
```

### 4. Test Configuration

Always test your configuration:

```bash
# Validate syntax
balancerx -config=config.yaml -validate

# Test with a small load
for i in {1..10}; do curl http://localhost:8080; done
```

## Troubleshooting Configuration

### Configuration Not Loading

```bash
# Check file permissions
ls -la /etc/balancerx/config.yaml

# Check file syntax
cat /etc/balancerx/config.yaml | yq eval .
```

### Backends Not Responding

```bash
# Test backend connectivity
curl http://localhost:9001/health
telnet localhost 6001

# Check DNS resolution
nslookup api.example.com
```

### Health Checks Failing

```bash
# Test health check endpoint
curl http://localhost:9001/health

# Check health check logs
sudo journalctl -u balancerx | grep "health"
```
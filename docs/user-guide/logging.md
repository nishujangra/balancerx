# Logging

BalancerX provides comprehensive logging for monitoring, debugging, and auditing. This document explains the logging system and how to access logs.

## Logging Overview

### Log Types

BalancerX generates several types of logs:

- **Request logs**: HTTP request forwarding details
- **Connection logs**: TCP connection handling
- **Health check logs**: Backend health monitoring
- **Error logs**: System errors and failures
- **Configuration logs**: Configuration changes and validation

### Log Format

All logs follow a consistent format:

```
2025/08/27 15:43:30 [LEVEL] Message details
```

**Example logs:**
```
2025/08/27 15:43:30 [RESPONSE] http://localhost:9003 -> 200 (18.682896ms)
2025/05/19 03:18:26 [FORWARD] [2025-05-19T03:18:26+05:30] GET / -> http://localhost:9002
2025/05/24 02:46:22 [TCP] Forwarding to localhost:9002
2025/08/30 02:36:09 [HEALTH] Starting health checker with interval: 10s
2025/05/19 03:18:26 http: proxy error: dial tcp 127.0.0.1:9002: connect: connection refused
```

## Log Access

### Systemd Journal (journalctl)

BalancerX logs are automatically captured by systemd and can be accessed via journalctl:

```bash
# View all logs
sudo journalctl -u balancerx

# Follow logs in real-time
sudo journalctl -u balancerx -f

# View recent logs
sudo journalctl -u balancerx -n 50

# View logs since a specific time
sudo journalctl -u balancerx --since "1 hour ago"
```

### Log File

BalancerX also writes logs to a file:

```bash
# View log file
sudo tail -f /var/log/balancerx/balancerx.log

# View recent entries
sudo tail -n 100 /var/log/balancerx/balancerx.log
```

**Note**: Logging configuration is not available in v1.0.0. All logs are written at the default level and managed by systemd.

## Log Categories

### Request Logs

#### HTTP Request Logs

```
2025/05/19 03:18:26 [FORWARD] [2025-05-19T03:18:26+05:30] GET / -> http://localhost:9002
2025/05/19 03:18:31 [FORWARD] [2025-05-19T03:18:31+05:30] GET / -> http://localhost:9003
2025/05/19 03:18:34 [FORWARD] [2025-05-19T03:18:34+05:30] GET / -> http://localhost:9001
```

**Log Details:**
- **Method**: HTTP method (GET, POST, PUT, etc.)
- **Path**: Request path
- **Backend**: Selected backend URL
- **Timestamp**: Request timestamp with timezone

#### HTTP Response Logs

```
2025/08/27 15:43:30 [RESPONSE] http://localhost:9003 -> 200 (18.682896ms)
2025/08/27 15:43:30 [RESPONSE] http://localhost:9001 -> 200 (16.457063ms)
2025/08/27 15:43:30 [RESPONSE] http://localhost:9002 -> 200 (53.655024ms)
```

**Log Details:**
- **Backend**: Backend URL that responded
- **Status Code**: HTTP response status code
- **Response Time**: Time taken to get response from backend

#### TCP Connection Logs

```
2025/05/24 02:46:22 [TCP] Forwarding to localhost:9002
2025/05/24 02:46:33 [TCP] Forwarding to localhost:9001
2025/05/24 02:46:39 [TCP] Forwarding to localhost:9002
2025/05/24 02:46:29 [TCP] Connection failed to localhost:9003: dial tcp 127.0.0.1:9003: connect: connection refused
```

**Log Details:**
- **Action**: Connection action (Forwarding, Connection failed, Listening)
- **Backend**: Backend address and port
- **Error**: Connection error details (if applicable)
- **Timestamp**: Connection timestamp

### Error Logs

#### Backend Errors

```
2025/05/19 03:18:26 http: proxy error: dial tcp 127.0.0.1:9002: connect: connection refused
2025/05/24 02:46:29 [TCP] Connection failed to localhost:9003: dial tcp 127.0.0.1:9003: connect: connection refused
```

#### System Errors

```
2025/05/20 03:30:43 listen tcp :8080: bind: address already in use
2025/05/20 03:33:04 listen tcp: address 8080: missing port in address
2025/05/24 02:45:24 [TCP] Connection failed to http://localhost:9002: dial tcp: address http://localhost:9002: too many colons in address
```

### Configuration Logs

#### Service Events

```
2025/09/07 03:26:41 Running BalancerX on port 8080 using 'round-robin' strategy
2025/09/07 03:26:41 Starting HTTP proxy on :8080
2025/08/30 02:36:15 Shutdown signal received, shutting down gracefully...
2025/08/30 02:36:15 HTTP server stopped
2025/08/30 02:36:15 BalancerX shutdown complete
```

**Note**: Configuration logs show service startup, shutdown, and basic configuration information.

## Log Analysis

### Request Analysis

#### Backend Distribution

```bash
# Count requests per backend
sudo journalctl -u balancerx | grep "FORWARD" | awk '{print $NF}' | sort | uniq -c

# Output example:
#   150 http://localhost:9001
#   148 http://localhost:9002
#   152 http://localhost:9003
```

#### Request Patterns

```bash
# Analyze request methods
sudo journalctl -u balancerx | grep "FORWARD" | awk '{print $3}' | sort | uniq -c

# Output example:
#   300 GET
#   100 POST
#   50 PUT
```

#### Response Times

```bash
# Monitor request timing (if available)
sudo journalctl -u balancerx | grep "FORWARD" | awk '{print $1, $2, $NF}'
```

### Health Check Analysis

#### Health Check Status

```bash
# Monitor health check startup
sudo journalctl -u balancerx | grep "HEALTH" | tail -20

# Check health checker configuration
sudo journalctl -u balancerx | grep "Starting health checker"
```

**Note**: Individual backend health status monitoring is not currently implemented in v1.0.0. Health check logs only show when the health checker starts.

### Error Analysis

#### Error Frequency

```bash
# Count error types
sudo journalctl -u balancerx | grep "ERROR" | awk '{print $3}' | sort | uniq -c

# Monitor error trends
sudo journalctl -u balancerx | grep "ERROR" | tail -20
```

#### Backend Availability

```bash
# Monitor backend connection failures
sudo journalctl -u balancerx | grep "Connection failed" | wc -l

# Check proxy errors
sudo journalctl -u balancerx | grep "proxy error" | tail -10

# Monitor port binding issues
sudo journalctl -u balancerx | grep "address already in use"
```
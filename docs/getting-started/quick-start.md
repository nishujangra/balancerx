# Quick Start

Get BalancerX up and running in minutes with this quick start guide.

## Prerequisites

- Linux system (Ubuntu/Debian recommended)
- Basic understanding of load balancing concepts

## Step 1: Install BalancerX

```bash
# Download and install
wget https://github.com/nishujangra/balancerx/releases/latest/download/balancerx_1.0.0.deb
sudo dpkg -i balancerx_1.0.0.deb
```

## Step 2: Create Test Backends

First, compile the Go dummy server:

```bash
# Navigate to the dummy-server directory
cd /path/to/balancerx/dummy-server

# Compile the dummy server
go build -o dummy-server dummy-golang.go
```

Set up some HTTP servers to test load balancing:

```bash
# Terminal 1 - Backend 1
./dummy-server 9001

# Terminal 2 - Backend 2  
./dummy-server 9002

# Terminal 3 - Backend 3
./dummy-server 9003
```

## Step 3: Configure BalancerX

Create a configuration file:

```bash
sudo nano /etc/balancerx/config.yaml
```

Add this configuration:

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

## Step 4: Start BalancerX

```bash
# Start the service
sudo systemctl start balancerx

# Check status
sudo systemctl status balancerx

# View logs
sudo journalctl -u balancerx -f
```

## Step 5: Test Load Balancing

```bash
# Make multiple requests to see load balancing in action
curl http://localhost:8080
curl http://localhost:8080
curl http://localhost:8080
```

You should see responses like "Hello from Echo server on port 9001!" being distributed across your backends. Check the logs to see the distribution:

```bash
sudo journalctl -u balancerx | tail -10
```

## Step 6: Test Health Checks

Test the health check endpoint directly:

```bash
# Test health endpoints on each backend
curl http://localhost:9001/health
curl http://localhost:9002/health
curl http://localhost:9003/health
```

You should see responses like "Ok at 9001!" from each backend.

Simulate a backend failure:

```bash
# Stop one backend (Ctrl+C in its terminal or kill the process)
pkill -f "dummy-server 9002"
```

Make more requests:

```bash
curl http://localhost:8080
curl http://localhost:8080
```

Notice that BalancerX will skip the failed backend and only route to healthy ones.

## Step 7: Try Different Strategies

Edit the configuration to try the random strategy:

```bash
sudo nano /etc/balancerx/config.yaml
```

Change the strategy:

```yaml
strategy: random  # Changed from round-robin
```

Restart BalancerX:

```bash
sudo systemctl restart balancerx
```

Test again:

```bash
curl http://localhost:8080
curl http://localhost:8080
curl http://localhost:8080
```

## TCP Load Balancing Example

For TCP load balancing, create a new config:

```yaml
port: 9090
protocol: tcp
strategy: round-robin
backends:
  - localhost:6001
  - localhost:6002
```

Set up TCP servers:

```bash
# Terminal 1
nc -lk 6001

# Terminal 2  
nc -lk 6002
```

Test TCP load balancing:

```bash
# Connect via BalancerX
telnet localhost 9090
```

## What's Next?

Now that you have BalancerX running:

1. **Explore Configuration**: See the [Configuration Guide](configuration.md) for all available options
2. **Learn Strategies**: Understand [Load Balancing Strategies](user-guide/strategies.md) in detail
3. **Set up Health Checks**: Configure [Health Monitoring](user-guide/health-checks.md) for production
4. **Monitor Performance**: Check [Performance Benchmarks](performance/benchmarks.md) for optimization

## Troubleshooting

### Common Issues

**BalancerX won't start:**
```bash
# Check logs
sudo journalctl -u balancerx -n 50
```

**No load balancing happening:**
```bash
# Check if backends are reachable
curl http://localhost:9001
curl http://localhost:9002
curl http://localhost:9003

# Check health endpoints
curl http://localhost:9001/health
curl http://localhost:9002/health
curl http://localhost:9003/health

# Check BalancerX logs
sudo journalctl -u balancerx -f
```

**Port already in use:**
```bash
# Find what's using the port
sudo netstat -tlnp | grep :8080

# Change port in config or kill the process
kill -9 ${PID}
```

## Cleanup

To stop everything:

```bash
# Stop BalancerX
sudo systemctl stop balancerx

# Stop test backends
pkill -f "dummy-server"
pkill -f "nc -lk"
```

Congratulations! You've successfully set up and tested BalancerX. 🎉

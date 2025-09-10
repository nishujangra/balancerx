# Installation

BalancerX can be installed in several ways. Choose the method that best fits your environment.

## Method 1: Debian Package (Recommended)

The easiest way to install BalancerX is using the pre-built Debian package.

### Download and Install

```bash
# Download the latest release
wget https://github.com/nishujangra/balancerx/releases/latest/download/balancerx_1.0.0.deb

# Install the package
sudo dpkg -i balancerx_1.0.0.deb

# Start the service
sudo systemctl start balancerx

# Enable auto-start on boot
sudo systemctl enable balancerx
```

### What Gets Installed

The Debian package automatically:

- Installs BalancerX to `/usr/bin/balancerx`
- Creates system user `balancerx`
- Sets up systemd service
- Creates default configuration at `/etc/balancerx/config.yaml`
- Sets up logging at `/var/log/balancerx/balancerx.log`

### Service Management

```bash
# Check status
sudo systemctl status balancerx

# View logs
sudo journalctl -u balancerx -f

# Edit configuration
sudo nano /etc/balancerx/config.yaml

# Restart after config changes
sudo systemctl restart balancerx

# Stop the service
sudo systemctl stop balancerx
```

## Method 2: Build from Source

### Prerequisites

- Go 1.19 or later
- Git

### Build Steps

```bash
# Clone the repository
git clone https://github.com/nishujangra/balancerx.git
cd balancerx

# Build the binary
go build -o build/balancerx main.go

# Run directly
./build/balancerx -config=config.yaml
```

## Method 3: Docker (Coming Soon)

Docker support is planned for future releases. This will provide:

- Containerized deployment
- Docker Compose examples
- Kubernetes manifests

## Verification

After installation, verify BalancerX is working:

```bash
# Check if binary exists and is executable
ls -la /usr/bin/balancerx

# Man Page 
man balancerx

# Test with a simple config
echo "port: 8080
protocol: http
strategy: round-robin
backends:
  - http://httpbin.org/get" > test-config.yaml

# Run with test config
balancerx -config=test-config.yaml
```

## Troubleshooting

### Common Issues

#### Permission Denied
```bash
# Make sure the binary is executable
chmod +x /usr/bin/balancerx

# Or if building from source
chmod +x build/balancerx
```

#### Port Already in Use
```bash
# Check what's using the port
lsof -i:8080

# Kill the process or change port in config
sudo kill -9 <PID>
```

#### Configuration File Not Found
```bash
# Check if config file exists
ls -la /etc/balancerx/config.yaml

# Or specify custom config path
balancerx -config=/path/to/your/config.yaml
```

### Getting Help

- Check the [Configuration Guide](configuration.md) for config issues
- Review [Health Checks](user-guide/health-checks.md) for backend connectivity
- Open an issue on [GitHub](https://github.com/nishujangra/balancerx/issues)

## Next Steps

After installation:

1. [Configure BalancerX](configuration.md) for your environment
2. [Set up backends](user-guide/strategies.md) and test load balancing
3. [Configure health checks](user-guide/health-checks.md) for reliability
4. [Monitor performance](performance/benchmarks.md) and optimize

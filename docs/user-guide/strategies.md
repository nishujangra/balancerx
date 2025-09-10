# Load Balancing Strategies

BalancerX supports multiple load balancing strategies to distribute traffic across backend servers. Each strategy has different characteristics and use cases.

## Available Strategies

### Round-Robin

**Strategy name**: `round-robin`

Distributes requests evenly across backends in a fixed, circular order.

#### How it Works

1. Maintains a counter for the current backend
2. Routes each request to the next backend in sequence
3. Wraps around to the first backend after the last one

#### Example

```yaml
strategy: round-robin
backends:
  - http://backend1:8080
  - http://backend2:8080
  - http://backend3:8080
```

**Request Distribution:**
- Request 1 → backend1
- Request 2 → backend2  
- Request 3 → backend3
- Request 4 → backend1
- Request 5 → backend2
- And so on...

#### Use Cases

- **Equal capacity backends**: When all backends have similar performance
- **Predictable load**: When you want consistent, even distribution
- **Simple scenarios**: When you don't need complex routing logic

#### Advantages

- ✅ Predictable and even distribution
- ✅ Simple to understand and debug
- ✅ Low computational overhead
- ✅ Works well with health checks

#### Disadvantages

- ❌ Doesn't consider backend load or capacity
- ❌ May not be optimal for backends with different performance
- ❌ Can create hotspots if backends have different response times

### Random

**Strategy name**: `random`

Randomly selects a backend for each request.

#### How it Works

1. Generates a random number for each request
2. Uses modulo operation to select a backend
3. No state maintained between requests

#### Example

```yaml
strategy: random
backends:
  - http://backend1:8080
  - http://backend2:8080
  - http://backend3:8080
```

**Request Distribution:**
- Request 1 → backend2 (random)
- Request 2 → backend1 (random)
- Request 3 → backend3 (random)
- Request 4 → backend1 (random)
- And so on...

#### Use Cases

- **Stateless applications**: When requests are independent
- **High traffic**: When you have many concurrent requests
- **Load testing**: When you want to simulate random load patterns

#### Advantages

- ✅ No state to maintain
- ✅ Good for high-concurrency scenarios
- ✅ Simple implementation
- ✅ Naturally distributes load over time

#### Disadvantages

- ❌ Less predictable than round-robin
- ❌ May create temporary imbalances
- ❌ Doesn't consider backend performance

## Planned Strategies

### Least-Connections

**Strategy name**: `least-conn` *(Coming Soon)*

Routes requests to the backend with the fewest active connections.

#### How it Will Work

1. Track active connections per backend
2. Route new requests to backend with minimum connections
3. Update connection counts on connect/disconnect

#### Use Cases

- **Long-running connections**: When connections have varying durations
- **Different backend capacities**: When backends have different processing power
- **Real-time applications**: WebSocket, streaming, etc.

### IP-Hash

**Strategy name**: `ip-hash` *(Coming Soon)*

Routes requests based on client IP address hash for sticky sessions.

#### How it Will Work

1. Hash the client IP address
2. Use hash to consistently route to the same backend
3. Maintain session affinity

#### Use Cases

- **Session-based applications**: When you need sticky sessions
- **Caching scenarios**: When backend caches are not shared
- **Stateful applications**: When backend state matters

### Weighted Round-Robin

**Strategy name**: `weighted-rr` *(Coming Soon)*

Round-robin with configurable weights for each backend.

#### How it Will Work

1. Assign weights to each backend
2. Route more requests to higher-weighted backends
3. Maintain weighted distribution

#### Use Cases

- **Different backend capacities**: When backends have different performance
- **Gradual scaling**: When adding/removing backends gradually
- **Cost optimization**: When some backends are more expensive

## Strategy Comparison

| Strategy | Predictability | State | Use Case | Complexity |
|----------|---------------|-------|----------|------------|
| `round-robin` | High | Low | Equal backends | Low |
| `random` | Low | None | High concurrency | Low |
| `least-conn` | Medium | High | Varying loads | Medium |
| `ip-hash` | High | Medium | Sticky sessions | Medium |
| `weighted-rr` | High | Low | Different capacities | Medium |


## Performance Considerations

### Round-Robin Performance

- **CPU**: Very low overhead
- **Memory**: Minimal state (just a counter)
- **Latency**: No additional latency

### Random Performance

- **CPU**: Low overhead (random number generation)
- **Memory**: No state maintained
- **Latency**: No additional latency

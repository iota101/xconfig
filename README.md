# xconfig

[![CI](https://github.com/iota101/xconfig/actions/workflows/ci.yml/badge.svg)](https://github.com/iota101/xconfig/actions/workflows/ci.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/iota101/xconfig.svg)](https://pkg.go.dev/github.com/iota101/xconfig)
[![Go Report Card](https://goreportcard.com/badge/github.com/iota101/xconfig)](https://goreportcard.com/report/github.com/iota101/xconfig)
[![Latest Tag](https://img.shields.io/github/v/tag/iota101/xconfig?label=latest%20tag&sort=semver)](https://github.com/iota101/xconfig/tags)

Typed configuration from YAML files and environment variables for Go.

## Features

- Load configuration from YAML files
- Load secrets from environment variables
- Type-safe value access with `String()`, `Int()`, `Bool()`, etc.
- Default values with `StringOr()`, `IntOr()`, `BoolOr()`, etc.
- Dot notation for nested YAML keys
- Test helpers for mocking config and secrets

## Installation

```bash
go get github.com/iota101/xconfig
```

## Quick Start

### Define Keys

```go
package main

import "github.com/iota101/xconfig"

// Config keys (dot notation for nesting)
const (
    ServerHost xconfig.K = "server.host"
    ServerPort xconfig.K = "server.port"
    DBHost     xconfig.K = "database.host"
)

// Environment keys
const (
    DBPassword xconfig.E = "DATABASE_PASSWORD"
    APIKey     xconfig.E = "API_KEY"
)
```

### Load Configuration

```go
// Load from YAML file
cfg, err := xconfig.FromYAML("config.yaml")
if err != nil {
    log.Fatal(err)
}

// Load from environment
secret := xconfig.FromEnv()

// Use values
host := cfg.Get(ServerHost).StringOr("localhost")
port := cfg.Get(ServerPort).IntOr(8080)
dbPass := secret.Get(DBPassword).String() // panics if not set
```

### YAML File Example

```yaml
# config.yaml
server:
  host: "0.0.0.0"
  port: 3000

database:
  host: "postgres.example.com"
```

## API Reference

### Types

```go
type K string  // Config key (dot notation: "database.host")
type E string  // Environment variable key
```

### Interfaces

```go
type Config interface {
    Get(key K) Value
    Has(key K) bool
}

type Secret interface {
    Get(key E) Value
    Has(key E) bool
}

type Value interface {
    String() string          // panics if missing
    Int() int                // panics if missing
    Int64() int64            // panics if missing
    Float64() float64        // panics if missing
    Bool() bool              // panics if missing

    StringOr(def string) string    // returns default if missing
    IntOr(def int) int
    Int64Or(def int64) int64
    Float64Or(def float64) float64
    BoolOr(def bool) bool

    IsEmpty() bool
}
```

### Constructors

| Function | Returns | Description |
|----------|---------|-------------|
| `FromYAML(path string)` | `Config, error` | Load config from YAML file |
| `FromEnv()` | `Secret` | Load secrets from environment |
| `FromMap(map[K]any)` | `Config` | Create mock config for tests |
| `FromEnvMap(map[E]string)` | `Secret` | Create mock secrets for tests |

## Testing

Use mock constructors to avoid real files and environment variables in tests:

```go
func TestMyService(t *testing.T) {
    cfg := xconfig.FromMap(map[xconfig.K]any{
        "server.port":   8080,
        "server.host":   "localhost",
        "feature.debug": true,
    })

    secret := xconfig.FromEnvMap(map[xconfig.E]string{
        "API_KEY":     "test-key",
        "DB_PASSWORD": "test-pass",
    })

    svc := NewMyService(cfg, secret)
    // ... test your service
}
```

## Development

### Using Task

```bash
# Install task: https://taskfile.dev
brew install go-task

# Run tests
task test

# Run tests with coverage
task test:cover

# Run all checks
task check

# Create a version tag
task tag -- v1.0.0

# Push a version tag
task tag:push -- v1.0.0
```

### Manual Commands

```bash
go test ./...              # Run tests
go test -race -cover ./... # Tests with race detector
go fmt ./...               # Format code
go vet ./...               # Check for issues
```

## License

MIT

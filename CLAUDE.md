# xconfig

Typed configuration from YAML files and environment variables for Go.

## Quick Start

```bash
go test ./...                    # Run tests
go test -race -cover ./...       # Tests with race detector and coverage
```

## Structure

```
xconfig/
├── config.go       # Public types: K, E, Value, Config, Secret interfaces
├── value.go        # Value implementation with type conversions
├── yaml.go         # FromYAML() — load Config from YAML file
├── env.go          # FromEnv() — load Secret from environment
├── mock.go         # FromMap(), FromEnvMap() — test helpers
└── xconfig_test.go # Table-driven tests
```

## API

### Types

```go
type K string  // Config key: "database.host" (dot notation for nesting)
type E string  // Env key: "DATABASE_PASSWORD"
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
    String() string          // Panics if missing
    Int() int
    Bool() bool
    StringOr(def string) string  // Returns default if missing
    IntOr(def int) int
    IsEmpty() bool
}
```

### Constructors

| Function | Returns | Use Case |
|----------|---------|----------|
| `FromYAML(path)` | `Config, error` | Load from YAML file |
| `FromEnv()` | `Secret` | Read real env vars |
| `FromMap(map[K]any)` | `Config` | Mock in tests |
| `FromEnvMap(map[E]string)` | `Secret` | Mock secrets in tests |

## Usage Pattern

```go
// Define keys
const (
    DBHost xconfig.K = "database.host"
    DBPass xconfig.E = "DATABASE_PASSWORD"
)

// Load
cfg, _ := xconfig.FromYAML("config.yaml")
secret := xconfig.FromEnv()

// Use
host := cfg.Get(DBHost).StringOr("localhost")
pass := secret.Get(DBPass).String()  // Panics if not set
```

## Testing

```go
cfg := xconfig.FromMap(map[xconfig.K]any{
    "server.port": 8080,
})
secret := xconfig.FromEnvMap(map[xconfig.E]string{
    "API_KEY": "test-key",
})
```

## Notes

- YAML keys use dot notation: `"server.host"` → `server.host` in YAML
- `String()`, `Int()`, `Bool()` panic if key missing — use `*Or()` variants for optional values
- `FromMap` uses flat keys (no nesting): `"server.port"` not `server: {port: ...}`
- Library pattern: flat structure, no internal/pkg needed

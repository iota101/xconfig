# xconfig - Project Overview

## Purpose
xconfig is a Go library for typed configuration management. It loads configuration from YAML files and secrets from environment variables with a clean, type-safe API.

## Tech Stack
- **Language**: Go
- **Module**: `github.com/iota101/xconfig`
- **Dependencies**: `gopkg.in/yaml.v3` for YAML parsing

## Project Structure
```
xconfig/
├── config.go       # Public types: K, E, Value, Config, Secret interfaces
├── value.go        # Value implementation with type conversions
├── yaml.go         # FromYAML() — load Config from YAML file
├── env.go          # FromEnv() — load Secret from environment
├── mock.go         # FromMap(), FromEnvMap() — test helpers
└── xconfig_test.go # Table-driven tests
```

## Key Types
- `K` - Config key type (dot notation: "database.host")
- `E` - Environment variable key type
- `Config` - Interface for YAML-based configuration
- `Secret` - Interface for environment-based secrets
- `Value` - Interface for typed value access (String, Int, Bool, etc.)

## Key Functions
- `FromYAML(path)` - Load Config from YAML file
- `FromEnv()` - Load Secret from real environment variables
- `FromMap(map[K]any)` - Mock Config for tests
- `FromEnvMap(map[E]string)` - Mock Secret for tests

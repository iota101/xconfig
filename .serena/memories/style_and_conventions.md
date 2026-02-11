# xconfig - Style and Conventions

## Code Style
- **Flat structure**: No internal/pkg directories, library-style layout
- **Naming**: Standard Go conventions (camelCase for private, PascalCase for public)
- **Types**: Custom string types for keys (`K` for config, `E` for env)

## API Design
- Interfaces (`Config`, `Secret`, `Value`) define contracts
- Panicking methods (`String()`, `Int()`) for required values
- `*Or()` variants for optional values with defaults
- Dot notation for nested YAML keys: `"server.host"`

## Testing
- Table-driven tests preferred
- Mock functions for testing: `FromMap()`, `FromEnvMap()`
- Tests in `xconfig_test.go`

## Patterns
- No over-engineering, keep it simple
- No unnecessary abstractions
- Flat keys in mock functions (no nesting)

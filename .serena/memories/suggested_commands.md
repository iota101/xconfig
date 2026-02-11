# xconfig - Development Commands

## Testing
```bash
# Run all tests
go test ./...

# Tests with race detector and coverage
go test -race -cover ./...

# Verbose test output
go test -v ./...

# Run specific test
go test -run TestName ./...
```

## Building & Checking
```bash
# Check compilation
go build ./...

# Format code
go fmt ./...

# Vet for issues
go vet ./...

# Tidy dependencies
go mod tidy
```

## Utility Commands (macOS/Darwin)
```bash
# List files
ls -la

# Find files
find . -name "*.go"

# Search in files
grep -r "pattern" .
```

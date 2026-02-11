# xconfig - Task Completion Checklist

When completing a task, run:

1. **Format code**
   ```bash
   go fmt ./...
   ```

2. **Vet for issues**
   ```bash
   go vet ./...
   ```

3. **Run tests**
   ```bash
   go test -race -cover ./...
   ```

4. **Tidy dependencies** (if go.mod changed)
   ```bash
   go mod tidy
   ```

All checks should pass before considering a task complete.

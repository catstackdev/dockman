.PHONY: build test clean install lint fmt completions

# Build binary
build:
	go build -o dockman .

# Run tests
test:
	go test -v -race ./...

# Run tests with coverage
test-coverage:
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

# Install locally
install:
	go install

# Clean build artifacts
clean:
	rm -f dockman
	rm -f coverage.out
	rm -rf dist/
	rm -rf completions/

# Lint code
lint:
	golangci-lint run

# Format code
fmt:
	gofmt -s -w .
	goimports -w .

# Generate shell completions
completions:
	./scripts/completions.sh

# Run all checks before commit
check: fmt lint test
	@echo "✓ All checks passed!"

# Release (for maintainers)
release:
	@echo "Creating release..."
	@read -p "Enter version (e.g., 0.5.0): " version; \
	git tag -a v$$version -m "Release v$$version"; \
	git push origin v$$version
	@echo "✓ Release v$$version created!"

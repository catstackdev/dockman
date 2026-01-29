# Contributing to Dockman

Thanks for your interest in contributing! 🎉

## Development Setup

```bash
# Clone the repository
git clone https://github.com/catstackdev/dockman.git
cd dockman

# Install dependencies
go mod download

# Build
go build -o dockman .

# Run tests
go test ./...

# Install locally
go install
```

## Running Tests

```bash
# All tests
go test ./...

# With coverage
go test ./... -cover

# Specific package
go test ./internal/compose -v
```

## Code Style

- Follow standard Go conventions
- Run `gofmt` before committing
- Add tests for new features
- Update README for user-facing changes

## Pull Request Process

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes
4. Run tests (`go test ./...`)
5. Commit your changes (`git commit -m 'Add amazing feature'`)
6. Push to the branch (`git push origin feature/amazing-feature`)
7. Open a Pull Request

## Release Process

Releases are automated via GitHub Actions when a tag is pushed:

```bash
git tag -a v0.5.0 -m "Release v0.5.0"
git push origin v0.5.0
```

## Questions?

Open an issue or discussion on GitHub!

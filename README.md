# 🐳 Dockman

Docker Compose manager with presets and shortcuts.

## Status: 🚧 Work in Progress

Currently learning Go by building a practical CLI tool. Phase 1 complete!

## Installation
```bash
go install github.com/catstackdev/dockman@latest
```

## Usage
```bash
# Start services
dockman up                    # All services
dockman up api postgres       # Specific services

# View logs
dockman logs -f               # Follow all logs
dockman logs api -f           # Follow specific service

# Stop services
dockman down                  # Stop everything
```

## Roadmap

- [x] Phase 1: Basic commands (up, down, logs)
- [ ] Phase 2: Presets system
- [ ] Phase 3: Service management (restart, ps, clean)
- [ ] Phase 4: Health checks

## Requirements

- Docker & Docker Compose installed
- Go 1.21+ (for building from source)

## Development
```bash
git clone https://github.com/catstackdev/dockman.git
cd dockman
go build -o dockman .
./dockman --help
```

## License

MIT

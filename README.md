# 🐳 Dockman

[![CI](https://github.com/catstackdev/dockman/workflows/CI/badge.svg)](https://github.com/catstackdev/dockman/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/catstackdev/dockman)](https://goreportcard.com/report/github.com/catstackdev/dockman)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/catstackdev/dockman)](https://github.com/catstackdev/dockman/releases)
Docker Compose manager with presets, shortcuts, and quality-of-life
improvements.

## Features

✨ **Auto-detect** docker-compose.yml from any subdirectory  
🎯 **Presets** - Start groups of services with one command  
🔧 **Custom Aliases** - Define project-specific shortcuts  
📦 **Smart cleanup** - Remove stopped containers, volumes, images  
🚀 **Quick access** - Shell into containers instantly  
🎨 **Pretty output** - Colored, easy-to-read terminal output

## Installation

```bash
go install github.com/catstackdev/dockman@latest
```

## Quick Start

```bash
# Start services
dockman up                    # All services
dockman up dev                # Using preset
dockman up api postgres       # Specific services

# View logs
dockman logs -f               # Follow all logs
dockman logs api -f           # Follow specific service

# Manage containers
dockman ps                    # List containers
dockman restart api           # Restart service
dockman exec api              # Shell into container
dockman down                  # Stop everything

# Build & pull
dockman build                 # Build images
dockman build --no-cache      # Fresh build
dockman pull                  # Update images

# Utilities
dockman port api 3000         # Show port mapping
dockman stats                 # Resource usage
dockman events                # Watch events
dockman list                  # List all services
```

## All Commands

### Core Operations

| Command              | Aliases      | Description       |
| -------------------- | ------------ | ----------------- |
| `up [services]`      | `u`, `start` | Start services    |
| `down`               | `d`, `stop`  | Stop all services |
| `restart [services]` | `r`          | Restart services  |
| `logs [services] -f` | `l`          | View logs         |
| `exec <service>`     | `e`, `sh`    | Shell access      |

### Service Management

| Command              | Description          |
| -------------------- | -------------------- |
| `build [services]`   | Build/rebuild images |
| `pull [services]`    | Pull latest images   |
| `pause [services]`   | Pause services       |
| `unpause [services]` | Unpause services     |
| `kill [services]`    | Force stop services  |

### Information & Monitoring

| Command                 | Aliases          | Description            |
| ----------------------- | ---------------- | ---------------------- |
| `ps`                    |                  | List containers        |
| `stats [services]`      | `top`            | Resource usage         |
| `port <service> [port]` |                  | Show port mappings     |
| `list`                  | `ls`, `services` | List all services      |
| `events`                |                  | Watch container events |

### Maintenance

| Command             | Description               |
| ------------------- | ------------------------- |
| `clean [-v\|--all]` | Clean up resources        |
| `validate`          | Check compose file syntax |

### Configuration

| Command           | Description         |
| ----------------- | ------------------- |
| `init`            | Create .dockman.yml |
| `config [--edit]` | View/edit config    |
| `preset list`     | List presets        |
| `aliases`         | List custom aliases |
| `info`            | Show project info   |
| `version`         | Show version        |

## Custom Aliases

Create `.dockman.yml`:

```yaml
default_preset: dev
auto_pull: false
aliases:
  db: 'up postgres redis'
  api: 'up api postgres'
  build-api: 'build --no-cache api'
```

Use them:

```bash
dockman db                    # → dockman up postgres redis
dockman api                   # → dockman up api postgres
dockman build-api             # → dockman build --no-cache api
```

## Presets

Configure in `~/.dockman/presets.yaml`:

```yaml
presets:
  dev:
    services: [postgres, redis, api, frontend]
    description: Full development stack

  api-only:
    services: [postgres, api]
    description: Backend development
```

## Examples

```bash
# Daily workflow
dockman up dev                # Start development environment
dockman logs api -f           # Watch API logs
dockman exec api npm test     # Run tests
dockman restart api           # Restart after changes

# Building
dockman build api             # Build API image
dockman build --no-cache      # Clean build
dockman pull && dockman up    # Update and start

# Debugging
dockman port api 3000         # Check port mapping
dockman stats                 # Monitor resources
dockman events                # Watch container events
dockman exec api /bin/sh      # Debug inside container

# Cleanup
dockman down                  # Stop all
dockman clean -v              # Remove volumes too
dockman clean --all           # Nuclear option
```

## Shell Completion

```bash
# Zsh
dockman completion zsh > ~/.zsh/completions/_dockman

# Bash
dockman completion bash > /usr/local/etc/bash_completion.d/dockman

# Fish
dockman completion fish > ~/.config/fish/completions/dockman.fish
```

## License

MIT

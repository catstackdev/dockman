# 🐳 Dockman

Docker Compose manager with presets and shortcuts.

## Features

✨ **Auto-detect** docker-compose.yml from any subdirectory  
🎯 **Presets** - Start groups of services with one command  
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

# Maintenance
dockman pull                  # Update images
dockman pull -f               # Skip confirmation prompt
dockman clean                 # Remove stopped containers
dockman clean -v              # Also remove volumes
dockman clean --all           # Nuclear option

# Info
dockman info                  # Show detected project
dockman version               # Show version
```

## Presets

Create `~/.dockman/presets.yaml`:

```yaml
presets:
  dev:
    services: [postgres, redis, api, frontend]
    description: Full development stack

  api-only:
    services: [postgres, api]
    description: Backend development

  db:
    services: [postgres, redis]
    description: Databases only
```

Then use:

```bash
dockman up dev
dockman up api-only
```

## Commands

| Command                 | Description                  |
| ----------------------- | ---------------------------- |
| `up [preset\|services]` | Start services or preset     |
| `down`                  | Stop all services            |
| `logs [services] -f`    | View logs (follow with -f)   |
| `ps [-q]`               | List containers              |
| `restart [services]`    | Restart services             |
| `exec <service> [cmd]`  | Execute command in container |
| `pull [services]`       | Pull latest images           |
| `clean [-v\|--all]`     | Clean up resources           |
| `preset list`           | List available presets       |
| `info`                  | Show project info            |
| `version`               | Show version                 |

## Shell Completion

```bash
# Zsh
dockman completion zsh > ~/.zsh/completions/_dockman

# Bash
dockman completion bash > /usr/local/etc/bash_completion.d/dockman

# Fish
dockman completion fish > ~/.config/fish/completions/dockman.fish
```

## Development

```bash
git clone https://github.com/catstackdev/dockman.git
cd dockman
go build -o dockman .
./dockman --help
```

## Roadmap

- [x] Phase 1: Basic commands (up, down, logs)
- [x] Phase 2: Presets system
- [x] Phase 3: Advanced features (clean, exec, pull)
- [ ] Phase 4: Health checks & monitoring
- [ ] Phase 5: Multi-project management

## License

MIT

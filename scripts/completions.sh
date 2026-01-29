#!/bin/bash
set -e

# Generate shell completions for release

mkdir -p completions

# Build binary first
go build -o dockman .

# Generate completions
./dockman completion bash > completions/dockman.bash
./dockman completion zsh > completions/_dockman
./dockman completion fish > completions/dockman.fish

echo "✓ Shell completions generated in ./completions/"

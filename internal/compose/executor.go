// internal/compose/executor.go
package compose

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/catstackdev/dockman/pkg/output"
)

// Executor handles docker-compose commands
type Executor struct {
	ComposeFile string
}

// NewExecutor creates a new executor
func NewExecutor() *Executor {
	return &Executor{
		ComposeFile: "docker-compose.yml", // default file
	}
}

// Up starts services
func (e *Executor) Up(services []string) error {
	args := []string{"up", "-d"}
	args = append(args, services...)

	output.Info(fmt.Sprintf("Starting services: %s", strings.Join(services, ", ")))
	return e.runCommand(args...)
}

// Down stops all services
func (e *Executor) Down() error {
	output.Info("Stopping all services...")
	return e.runCommand("down")
}

// Logs shows logs for services
func (e *Executor) Logs(services []string, follow bool) error {
	args := []string{"logs"}
	if follow {
		args = append(args, "-f")
	}
	args = append(args, services...)

	output.Info(fmt.Sprintf("Showing logs for: %s", strings.Join(services, ", ")))
	return e.runCommand(args...)
}

// runCommand executes docker-compose with given arguments
func (e *Executor) runCommand(args ...string) error {
	// Build command: docker-compose -f docker-compose.yml <args>
	cmdArgs := []string{"-f", e.ComposeFile}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("docker-compose", cmdArgs...)

	// Connect output to terminal
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	// Run command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("docker-compose failed: %w", err)
	}

	return nil
}

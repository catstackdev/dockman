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
	ProjectDir  string
}

// NewExecutor creates a new executor with auto-detected compose file
func NewExecutor() (*Executor, error) {
	composePath, err := FindComposeFile()
	if err != nil {
		return nil, err
	}

	return &Executor{
		ComposeFile: composePath,
		ProjectDir:  GetProjectDir(composePath),
	}, nil
}

// Up starts services
// func (e *Executor) Up(services []string) error {
// 	args := []string{"up", "-d"}
// 	args = append(args, services...)
//
// 	if len(services) == 0 {
// 		output.Info("Starting all services...")
// 	} else {
// 		output.Info(fmt.Sprintf("Starting services: %s", strings.Join(services, ", ")))
// 	}
//
// 	return e.runCommand(args...)
// }

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

	if len(services) == 0 {
		output.Info("Showing logs for all services...")
	} else {
		output.Info(fmt.Sprintf("Showing logs for: %s", strings.Join(services, ", ")))
	}

	return e.runCommand(args...)
}

// Ps shows container status
func (e *Executor) Ps() error {
	return e.runCommand("ps")
}

// Restart restarts services
func (e *Executor) Restart(services []string) error {
	args := []string{"restart"}
	args = append(args, services...)
	return e.runCommand(args...)
}

// runCommand executes docker-compose with given arguments
func (e *Executor) runCommand(args ...string) error {
	// Build command: docker-compose -f /path/to/docker-compose.yml <args>
	cmdArgs := []string{"-f", e.ComposeFile}
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("docker-compose", cmdArgs...)

	// Change to project directory (important for relative paths in compose file)
	cmd.Dir = e.ProjectDir

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

// GetInfo returns executor info for debugging
func (e *Executor) GetInfo() string {
	return fmt.Sprintf("Project: %s\nCompose file: %s", e.ProjectDir, e.ComposeFile)
}

// NewExecutorWithFile creates executor with specific compose file
func NewExecutorWithFile(composePath string) (*Executor, error) {
	if !fileExists(composePath) {
		return nil, fmt.Errorf("compose file not found: %s", composePath)
	}

	return &Executor{
		ComposeFile: composePath,
		ProjectDir:  GetProjectDir(composePath),
	}, nil
}

func (e *Executor) Clean(removeVolumes, removeAll bool) error {
	args := []string{"down"}

	if removeVolumes {
		args = append(args, "-v")
	}

	if removeAll {
		args = append(args, "--rmi", "all", "-v")
	}

	// Remove orphan containers
	args = append(args, "--remove-orphans")

	return e.runCommand(args...)
}

// Exec executes a command in a service container
func (e *Executor) Exec(service string, command []string) error {
	args := []string{"exec", service}
	args = append(args, command...)

	return e.runCommand(args...)
}

// PsQuiet shows only container IDs
func (e *Executor) PsQuiet() error {
	return e.runCommand("ps", "-q")
}

// Pull pulls images (updated with better output)
func (e *Executor) Pull(services []string) error {
	args := []string{"pull"}

	// Add --ignore-pull-failures to continue on errors
	args = append(args, "--ignore-pull-failures")

	args = append(args, services...)

	return e.runCommand(args...)
}

// ShowImages displays configured images from docker-compose.yml
func (e *Executor) ShowImages(services []string) error {
	args := []string{"config", "--services"}

	if len(services) > 0 {
		// Show specific services only
		for _, service := range services {
			args = append(args, service)
		}
	}

	return e.runCommand(args...)
}

// Stats shows container resource usage
func (e *Executor) Stats(services []string, noStream bool) error {
	args := []string{"stats"}

	if noStream {
		args = append(args, "--no-stream")
	}

	args = append(args, services...)

	return e.runCommand(args...)
}

// ListServices lists all services in compose file
func (e *Executor) ListServices() error {
	return e.runCommand("config", "--services")
}

// Validate checks compose file syntax
func (e *Executor) Validate() error {
	return e.runCommand("config", "--quiet")
}

// Port shows port mappings for a service
func (e *Executor) Port(service, privatePort string) error {
	args := []string{"port", service}
	if privatePort != "" {
		args = append(args, privatePort)
	}
	return e.runCommand(args...)
}

// Build builds or rebuilds services
func (e *Executor) Build(services []string, noCache, parallel, pull bool) error {
	args := []string{"build"}

	if noCache {
		args = append(args, "--no-cache")
	}

	if !parallel {
		args = append(args, "--no-parallel")
	}

	if pull {
		args = append(args, "--pull")
	}

	args = append(args, services...)

	return e.runCommand(args...)
}

// Events streams container events
func (e *Executor) Events() error {
	return e.runCommand("events")
}

// Pause pauses running services
func (e *Executor) Pause(services []string) error {
	args := []string{"pause"}
	args = append(args, services...)
	return e.runCommand(args...)
}

// Unpause unpauses paused services
func (e *Executor) Unpause(services []string) error {
	args := []string{"unpause"}
	args = append(args, services...)
	return e.runCommand(args...)
}

// Kill force stops services
func (e *Executor) Kill(services []string, signal string) error {
	args := []string{"kill"}

	if signal != "" && signal != "SIGKILL" {
		args = append(args, "-s", signal)
	}

	args = append(args, services...)
	return e.runCommand(args...)
}

// UpWithEnv starts services with custom env file
func (e *Executor) UpWithEnv(services []string, envFile string) error {
	args := []string{"up", "-d"}

	if envFile != "" {
		args = append([]string{"--env-file", envFile}, args...)
	}

	args = append(args, services...)

	if len(services) == 0 {
		output.Info("Starting all services...")
	} else {
		output.Info(fmt.Sprintf("Starting services: %s", strings.Join(services, ", ")))
	}

	return e.runCommand(args...)
}

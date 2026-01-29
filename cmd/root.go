// cmd/root.go
package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/internal/config"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var composeFile string

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "dockman",
	Short: "Docker Compose manager with presets",
	Long: `Dockman is a CLI tool that makes docker-compose easier to use.
It provides shortcuts, presets, and better log viewing.`,
	Example: `  dockman up dev           # Start dev preset
  dockman logs api -f      # Follow API logs
  dockman d                # Alias for 'down'
  dockman db               # Custom alias from .dockman.yml`,
	// Handle unknown commands as potential aliases
	SilenceErrors: true,
	SilenceUsage:  true,
}

// Execute runs the CLI with alias resolution
func Execute() {
	// Check for aliases before executing
	if len(os.Args) > 1 {
		if resolvedArgs := tryResolveAlias(os.Args[1:]); resolvedArgs != nil {
			// Replace args with resolved alias
			os.Args = append([]string{os.Args[0]}, resolvedArgs...)
		}
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flag
	rootCmd.PersistentFlags().StringVar(&composeFile, "file", "", "Specify docker-compose file path")

	// Register all commands
	rootCmd.AddCommand(upCmd)
	rootCmd.AddCommand(downCmd)
	rootCmd.AddCommand(logsCmd)
	rootCmd.AddCommand(presetCmd)
	rootCmd.AddCommand(psCmd)
	rootCmd.AddCommand(restartCmd)
	rootCmd.AddCommand(infoCmd)
	rootCmd.AddCommand(cleanCmd)
	rootCmd.AddCommand(execCmd)
	rootCmd.AddCommand(pullCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(statsCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(configCmd)
	rootCmd.AddCommand(validateCmd)
	rootCmd.AddCommand(aliasesCmd)
}

// tryResolveAlias attempts to resolve command as an alias
func tryResolveAlias(args []string) []string {
	if len(args) == 0 {
		return nil
	}

	// Try to get project directory
	executor, err := compose.NewExecutor()
	if err != nil {
		// No compose file, can't load aliases
		return nil
	}

	// Load project config
	cfg, err := config.LoadProjectConfig(executor.ProjectDir)
	if err != nil {
		return nil
	}

	// Check if first arg is an alias
	commandName := args[0]
	aliasCommand, exists := cfg.Aliases[commandName]

	if !exists {
		return nil
	}

	// Show what alias is doing
	output.Info(fmt.Sprintf("Alias '%s' → dockman %s",
		output.Cyan(commandName),
		output.Gray(aliasCommand)))

	// Parse alias command
	aliasArgs := strings.Fields(aliasCommand)

	// Append remaining args from original command
	if len(args) > 1 {
		aliasArgs = append(aliasArgs, args[1:]...)
	}

	return aliasArgs
}

// getExecutor returns executor with optional file override
func getExecutor() (*compose.Executor, error) {
	if composeFile != "" {
		return compose.NewExecutorWithFile(composeFile)
	}
	return compose.NewExecutor()
}

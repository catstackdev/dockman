package cmd

import (
	"fmt"
	"os"

	"github.com/catstackdev/dockman/internal/compose"
	"github.com/spf13/cobra"
)

var composeFile string

// rootCmd is the base command
var rootCmd = &cobra.Command{
	Use:   "dockman",
	Short: "Docker Compose manager with presets",
	Long: `Dockman is a CLI tool that makes docker-compose easier to use.
It provides shortcuts, presets, and better log viewing.`,
}

// Execute runs the CLI
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Global flag - NO SHORTHAND to avoid conflict with logs -f
	rootCmd.PersistentFlags().StringVar(&composeFile, "file", "", "Specify docker-compose file path")

	// Register subcommands here
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
}

// getExecutor returns executor with optional file override
func getExecutor() (*compose.Executor, error) {
	if composeFile != "" {
		// Use specified file
		return compose.NewExecutorWithFile(composeFile)
	}
	// Auto-detect
	return compose.NewExecutor()
}

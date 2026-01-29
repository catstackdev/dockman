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
	// Add examples to root command
	Example: `  dockman up dev           # Start dev preset
  dockman logs api -f      # Follow API logs
  dockman d                # Alias for 'down'
  dockman r api            # Alias for 'restart api'`,
}

// Execute runs the CLI
func Execute() {
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
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(shCmd)
}

// getExecutor returns executor with optional file override
func getExecutor() (*compose.Executor, error) {
	if composeFile != "" {
		return compose.NewExecutorWithFile(composeFile)
	}
	return compose.NewExecutor()
}

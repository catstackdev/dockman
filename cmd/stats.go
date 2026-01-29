package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var statsNoStream bool

var statsCmd = &cobra.Command{
	Use:     "stats [services...]",
	Aliases: []string{"top"},
	Short:   "Show container resource usage statistics",
	Long:    `Display live resource usage statistics for running containers (CPU, memory, network, disk I/O)`,
	Example: `  dockman stats              # Show stats for all containers
  dockman stats api          # Show stats for api container
  dockman stats --no-stream  # Show stats once, don't stream`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(args) == 0 {
			output.Info("Showing resource usage for all containers (Ctrl+C to exit)...")
		} else {
			output.Info(fmt.Sprintf("Showing resource usage for: %v", args))
		}

		if err := executor.Stats(args, statsNoStream); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

func init() {
	statsCmd.Flags().BoolVar(&statsNoStream, "no-stream", false, "Disable streaming stats and only pull the first result")
}

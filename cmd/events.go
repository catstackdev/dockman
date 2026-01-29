package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Receive real-time events from containers",
	Long:  `Stream container events (start, stop, die, etc.) in real-time`,
	Example: `  dockman events            # Watch all events
  # Press Ctrl+C to stop`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		output.Info("Watching container events (Ctrl+C to stop)...")

		if err := executor.Events(); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

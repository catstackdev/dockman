// cmd/logs.go
package cmd

import (
	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var follow bool // flag for -f

var logsCmd = &cobra.Command{
	Use:     "logs [services...]",
	Aliases: []string{"l"},
	Short:   "View service logs",
	Long:    `View logs from one or more services`,
	Example: `  dockman logs           # View all logs
  dockman logs api -f    # Follow API logs
  dockman logs api postgres  # View specific services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := compose.NewExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if err := executor.Logs(args, follow); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

func init() {
	// Add -f flag to logs command
	logsCmd.Flags().BoolVarP(&follow, "follow", "f", false, "Follow log output")
}

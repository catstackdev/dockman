package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var killSignal string

var killCmd = &cobra.Command{
	Use:   "kill [services...]",
	Short: "Force stop services",
	Long:  `Kill running services by sending SIGKILL (or custom signal)`,
	Example: `  dockman kill api          # Kill api service
  dockman kill -s SIGTERM   # Send SIGTERM instead`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(args) == 0 {
			output.Warning("Killing all services...")
		} else {
			output.Warning(fmt.Sprintf("Killing services: %v", args))
		}

		if err := executor.Kill(args, killSignal); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services killed!")
		return nil
	},
}

func init() {
	killCmd.Flags().StringVarP(&killSignal, "signal", "s", "SIGKILL", "Signal to send")
}

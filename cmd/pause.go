package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause [services...]",
	Short: "Pause services",
	Long:  `Pause running services (freeze processes)`,
	Example: `  dockman pause api         # Pause api service
  dockman pause              # Pause all services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(args) == 0 {
			output.Info("Pausing all services...")
		} else {
			output.Info(fmt.Sprintf("Pausing services: %v", args))
		}

		if err := executor.Pause(args); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services paused!")
		return nil
	},
}

var unpauseCmd = &cobra.Command{
	Use:   "unpause [services...]",
	Short: "Unpause services",
	Long:  `Unpause paused services`,
	Example: `  dockman unpause api       # Unpause api service
  dockman unpause            # Unpause all services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(args) == 0 {
			output.Info("Unpausing all services...")
		} else {
			output.Info(fmt.Sprintf("Unpausing services: %v", args))
		}

		if err := executor.Unpause(args); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services unpaused!")
		return nil
	},
}

package cmd

import (
	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var restartCmd = &cobra.Command{
	Use:   "restart [services...]",
	Short: "Restart services",
	Long:  `Restart one or more services`,
	Example: `  dockman restart api        # Restart API
  dockman restart             # Restart all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := compose.NewExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		output.Info("Restarting services...")
		if err := executor.Restart(args); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services restarted!")
		return nil
	},
}

// cmd/up.go
package cmd

import (
	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [services...]",
	Short: "Start services",
	Long:  `Start one or more services defined in docker-compose.yml`,
	Example: `  dockman up           # Start all services
  dockman up api postgres  # Start specific services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor := compose.NewExecutor()

		if err := executor.Up(args); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services started successfully!")
		return nil
	},
}

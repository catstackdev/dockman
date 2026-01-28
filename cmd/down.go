// cmd/down.go
package cmd

import (
	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "Stop all services",
	Long:  `Stop and remove all containers defined in docker-compose.yml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor := compose.NewExecutor()

		if err := executor.Down(); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("All services stopped!")
		return nil
	},
}

package cmd

import (
	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List running containers",
	Long:  `Show status of all containers defined in docker-compose.yml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := compose.NewExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if err := executor.Ps(); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

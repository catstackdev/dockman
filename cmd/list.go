package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "services"},
	Short:   "List all services in docker-compose.yml",
	Long:    `Display all services defined in the docker-compose.yml file`,
	Example: `  dockman list              # List all services
  dockman ls                # Alias for list`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		output.Info("Services defined in docker-compose.yml:")
		if err := executor.ListServices(); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

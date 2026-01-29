package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:     "validate",
	Aliases: []string{"check"},
	Short:   "Validate docker-compose.yml syntax",
	Long:    `Check if docker-compose.yml has valid syntax and configuration`,
	Example: `  dockman validate          # Check compose file
  dockman check             # Alias for validate`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		output.Info("Validating docker-compose.yml...")

		if err := executor.Validate(); err != nil {
			output.Error("Validation failed!")
			return err
		}

		output.Success("docker-compose.yml is valid!")
		return nil
	},
}

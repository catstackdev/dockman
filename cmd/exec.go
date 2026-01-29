package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec <service> [command]",
	Aliases: []string{"e", "sh"},
	Short:   "Execute command in a service container",
	Long:    `Execute a command in a running service container. Defaults to /bin/sh if no command specified.`,
	Example: `  dockman exec api              # Open shell in api container
  dockman exec api npm test     # Run npm test in api
  dockman exec postgres psql    # Open psql in postgres`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		service := args[0]
		command := []string{"/bin/sh"}

		if len(args) > 1 {
			command = args[1:]
		}

		output.Info(fmt.Sprintf("Executing in %s: %v", service, command))
		if err := executor.Exec(service, command); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

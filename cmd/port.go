package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var portCmd = &cobra.Command{
	Use:   "port [service] [private_port]",
	Short: "Show public port for a port binding",
	Long:  `Print the public port for a port binding`,
	Example: `  dockman port api 3000     # Show public port for api:3000
  dockman port postgres     # Show all ports for postgres`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		service := args[0]
		privatePort := ""

		if len(args) > 1 {
			privatePort = args[1]
		}

		if err := executor.Port(service, privatePort); err != nil {
			output.Error(err.Error())
			return err
		}

		return nil
	},
}

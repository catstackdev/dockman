package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var psQuiet bool

var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Long:  `Show status of all containers defined in docker-compose.yml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if psQuiet {
			if err := executor.PsQuiet(); err != nil {
				output.Error(err.Error())
				return err
			}
		} else {
			if err := executor.Ps(); err != nil {
				output.Error(err.Error())
				return err
			}
		}

		return nil
	},
}

func init() {
	psCmd.Flags().BoolVarP(&psQuiet, "quiet", "q", false, "Only show container IDs")
}

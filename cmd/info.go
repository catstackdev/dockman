package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/internal/compose"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show project information",
	Long:  `Display information about detected docker-compose project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := compose.NewExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		fmt.Println("\n📦 Project Information:\n")
		fmt.Println(executor.GetInfo())
		fmt.Println()

		return nil
	},
}

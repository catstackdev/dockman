package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var (
	cleanVolumes bool
	cleanAll     bool
)

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "Clean up Docker resources",
	Long:  `Remove stopped containers, unused networks, and optionally volumes`,
	Example: `  dockman clean              # Remove stopped containers
  dockman clean -v           # Also remove volumes
  dockman clean --all        # Remove everything (containers, volumes, images)`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if cleanAll {
			output.Warning("Removing all containers, volumes, and images...")
			if err := executor.Clean(true, true); err != nil {
				output.Error(err.Error())
				return err
			}
		} else if cleanVolumes {
			output.Info("Removing stopped containers and volumes...")
			if err := executor.Clean(true, false); err != nil {
				output.Error(err.Error())
				return err
			}
		} else {
			output.Info("Removing stopped containers...")
			if err := executor.Clean(false, false); err != nil {
				output.Error(err.Error())
				return err
			}
		}

		output.Success("Cleanup complete!")
		return nil
	},
}

func init() {
	cleanCmd.Flags().BoolVarP(&cleanVolumes, "volumes", "v", false, "Also remove volumes")
	cleanCmd.Flags().BoolVar(&cleanAll, "all", false, "Remove everything (containers, volumes, images)")
}

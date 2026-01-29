package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var (
	buildNoCache  bool
	buildParallel bool
	buildPull     bool
)

var buildCmd = &cobra.Command{
	Use:   "build [services...]",
	Short: "Build or rebuild services",
	Long:  `Build or rebuild service images defined in docker-compose.yml`,
	Example: `  dockman build                # Build all services
  dockman build api            # Build api service
  dockman build --no-cache     # Build without cache
  dockman build --pull         # Always pull newer base images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(args) == 0 {
			output.Info("Building all services...")
		} else {
			output.Info(fmt.Sprintf("Building services: %v", args))
		}

		if err := executor.Build(args, buildNoCache, buildParallel, buildPull); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Build complete!")
		return nil
	},
}

func init() {
	buildCmd.Flags().BoolVar(&buildNoCache, "no-cache", false, "Build without using cache")
	buildCmd.Flags().BoolVar(&buildParallel, "parallel", true, "Build images in parallel")
	buildCmd.Flags().BoolVar(&buildPull, "pull", false, "Always attempt to pull newer base images")
}

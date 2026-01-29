package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	pullForce bool
	pullCheck bool
)

var pullCmd = &cobra.Command{
	Use:   "pull [services...]",
	Short: "Pull service images",
	Long: `Pull service images from registry with confirmation.
	
Docker Compose checks image digests and only downloads updates.
Use -f to skip confirmation prompt.`,
	Example: `  dockman pull              # Pull with confirmation
  dockman pull -f           # Pull without confirmation
  dockman pull api -f       # Pull specific service
  dockman pull --check      # Show configured images`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		// Check mode
		if pullCheck {
			output.Info("Configured images in docker-compose.yml:")
			if err := executor.ShowImages(args); err != nil {
				output.Error(err.Error())
				return err
			}
			return nil
		}

		// Determine what we're pulling
		services := "all services"
		if len(args) > 0 {
			services = fmt.Sprintf("%v", args)
		}

		// Interactive confirmation if not forced
		if !pullForce {
			// Enhanced prompt with colors
			prompt := promptui.Prompt{
				Label: fmt.Sprintf("Pull latest images for %s? Docker will only download if updates are available",
					output.Cyan(services)),
				IsConfirm: true,
				Default:   "n",
			}

			result, err := prompt.Run()
			if err != nil || (result != "y" && result != "Y") {
				output.Info("Pull cancelled")
				return nil
			}

			fmt.Println() // Empty line for better readability
		}

		// Pull images
		if len(args) == 0 {
			output.Info("Checking for image updates (all services)...")
		} else {
			output.Info(fmt.Sprintf("Checking for updates: %v", args))
		}

		if err := executor.Pull(args); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Image pull complete!")
		return nil
	},
}

func init() {
	pullCmd.Flags().BoolVarP(&pullForce, "force", "f", false, "Skip confirmation prompt")
	pullCmd.Flags().BoolVar(&pullCheck, "check", false, "Show configured images without pulling")
}

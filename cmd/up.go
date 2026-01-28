// cmd/up.go
package cmd

import (
	"fmt"

	"github.com/catstackdev/dockman/internal/preset"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var upCmd = &cobra.Command{
	Use:   "up [preset|services...]",
	Short: "Start services or preset",
	Long:  `Start one or more services, or use a preset to start a group of services`,
	Example: `  dockman up dev              # Start 'dev' preset
  dockman up                  # Start all services
  dockman up api postgres     # Start specific services`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create executor (will auto-find docker-compose.yml)
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			output.Warning("Make sure you're in a directory with docker-compose.yml or one of its parent directories")
			return err
		}

		// Load preset manager
		presetMgr, err := preset.NewManager()
		if err != nil {
			output.Warning(fmt.Sprintf("Failed to load presets: %v", err))
			// Continue without presets
		}

		var services []string

		// Check if first arg is a preset
		if len(args) == 1 && presetMgr != nil && presetMgr.Exists(args[0]) {
			presetName := args[0]
			p, err := presetMgr.Get(presetName)
			if err != nil {
				output.Error(err.Error())
				return err
			}

			services = p.Services
			output.Info(fmt.Sprintf("Using preset '%s': %v", presetName, services))
		} else {
			// Use services directly from args
			services = args
		}

		// Start services
		if err := executor.Up(services); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success("Services started successfully!")
		return nil
	},
}

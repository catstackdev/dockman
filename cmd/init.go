package cmd

import (
	"fmt"
	"os"

	"github.com/catstackdev/dockman/internal/config"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dockman config for current project",
	Long:  `Create a .dockman.yml config file in the current project directory`,
	Example: `  dockman init              # Interactive setup
  cd ~/projects/myapp && dockman init`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error("No docker-compose.yml found in current or parent directories")
			return err
		}

		projectDir := executor.ProjectDir
		configPath := fmt.Sprintf("%s/.dockman.yml", projectDir)

		// Check if config already exists
		if _, err := os.Stat(configPath); err == nil {
			output.Warning(fmt.Sprintf("Config already exists at %s", configPath))

			prompt := promptui.Prompt{
				Label:     "Overwrite existing config",
				IsConfirm: true,
				Default:   "n",
			}

			result, err := prompt.Run()
			if err != nil || (result != "y" && result != "Y") {
				output.Info("Init cancelled")
				return nil
			}
		}

		// Interactive config creation
		cfg := &config.ProjectConfig{}

		// Ask for default preset
		presetPrompt := promptui.Prompt{
			Label:   "Default preset (leave empty for none)",
			Default: "",
		}
		preset, _ := presetPrompt.Run()
		if preset != "" {
			cfg.DefaultPreset = preset
		}

		// Ask for auto-pull
		autoPullPrompt := promptui.Prompt{
			Label:     "Auto-pull images on 'up'",
			IsConfirm: true,
			Default:   "n",
		}
		autoPull, _ := autoPullPrompt.Run()
		cfg.AutoPull = (autoPull == "y" || autoPull == "Y")

		// Save config
		if err := config.SaveProjectConfig(projectDir, cfg); err != nil {
			output.Error(err.Error())
			return err
		}

		output.Success(fmt.Sprintf("Created config at %s", configPath))
		fmt.Println("\nExample .dockman.yml:")
		fmt.Println("---")
		fmt.Println("default_preset: dev")
		fmt.Println("auto_pull: false")
		fmt.Println("aliases:")
		fmt.Println("  db: \"up postgres redis\"")
		fmt.Println("  api: \"up api postgres\"")

		return nil
	},
}

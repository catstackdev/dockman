package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/catstackdev/dockman/internal/config"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initForce bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dockman config for current project",
	Long:  `Create a .dockman.yml config file in the current project directory with interactive setup`,
	Example: `  dockman init              # Interactive setup
  dockman init --force      # Overwrite existing config`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.ErrorWithHelp(
				"No docker-compose.yml found in current or parent directories",
				"Make sure you're in a directory with docker-compose.yml",
			)
			return err
		}

		projectDir := executor.ProjectDir
		configPath := fmt.Sprintf("%s/.dockman.yml", projectDir)

		// Check if config already exists
		if !initForce {
			if _, err := os.Stat(configPath); err == nil {
				output.Warning(fmt.Sprintf("Config already exists at %s", configPath))

				prompt := promptui.Prompt{
					Label:     "Overwrite existing config",
					IsConfirm: true,
					Default:   "n",
				}

				result, err := prompt.Run()
				if err != nil || (result != "y" && result != "Y") {
					output.Info("Init cancelled. Use --force to overwrite.")
					return nil
				}
			}
		}

		// Start interactive setup
		output.Info("Let's set up dockman for your project!\n")

		cfg := &config.ProjectConfig{
			Aliases: make(map[string]string),
		}

		// 1. Default preset
		output.Info("Step 1: Default Preset")
		fmt.Println("  Which preset should 'dockman up' use by default?")
		fmt.Println("  (Leave empty to start all services)")

		presetPrompt := promptui.Prompt{
			Label:   "Default preset",
			Default: "",
		}
		preset, _ := presetPrompt.Run()
		if preset != "" {
			cfg.DefaultPreset = preset
		}

		// 2. Auto-pull
		fmt.Println()
		output.Info("Step 2: Auto Pull")
		fmt.Println("  Should images be checked/pulled automatically on 'dockman up'?")

		autoPullPrompt := promptui.Prompt{
			Label:     "Auto-pull images",
			IsConfirm: true,
			Default:   "n",
		}
		autoPull, _ := autoPullPrompt.Run()
		cfg.AutoPull = (autoPull == "y" || autoPull == "Y")

		// 3. Custom compose file
		fmt.Println()
		output.Info("Step 3: Compose File (Optional)")
		fmt.Println("  Using a non-standard compose file name?")
		fmt.Println("  (Default: docker-compose.yml)")

		composePrompt := promptui.Prompt{
			Label:   "Compose file path",
			Default: "",
		}
		composePath, _ := composePrompt.Run()
		if composePath != "" {
			cfg.ComposeFile = composePath
		}

		// 4. Aliases
		fmt.Println()
		output.Info("Step 4: Aliases (Optional)")
		fmt.Println("  Create custom command shortcuts?")
		fmt.Println("  Create custom command shortcuts?")
		fmt.Println("  Examples:")
		fmt.Println("    db   → up postgres redis")
		fmt.Println("    api  → up api postgres")
		fmt.Println("    dev  → up")
		fmt.Println()

		aliasPrompt := promptui.Prompt{
			Label:     "Add aliases",
			IsConfirm: true,
			Default:   "n",
		}
		addAliases, _ := aliasPrompt.Run()

		if addAliases == "y" || addAliases == "Y" {
			for {
				namePrompt := promptui.Prompt{
					Label: "Alias name (empty to finish)",
				}
				name, _ := namePrompt.Run()
				name = strings.TrimSpace(name)

				if name == "" {
					break
				}

				cmdPrompt := promptui.Prompt{
					Label: fmt.Sprintf("Command for '%s'", name),
				}
				command, _ := cmdPrompt.Run()
				command = strings.TrimSpace(command)

				if command != "" {
					cfg.Aliases[name] = command
					output.Success(fmt.Sprintf("Added: dockman %s → dockman %s", name, command))
				}
			}
		}

		// Save config
		if err := config.SaveProjectConfig(projectDir, cfg); err != nil {
			output.Error(err.Error())
			return err
		}

		// Show summary
		fmt.Println()
		output.Box("Configuration Created!", []string{
			fmt.Sprintf("Location: %s", configPath),
			"",
			"Settings:",
			fmt.Sprintf("  Default Preset: %s", valueOrNone(cfg.DefaultPreset)),
			fmt.Sprintf("  Auto Pull:      %v", cfg.AutoPull),
			fmt.Sprintf("  Compose File:   %s", valueOrNone(cfg.ComposeFile)),
			fmt.Sprintf("  Aliases:        %d", len(cfg.Aliases)),
		})

		// Show example usage
		if len(cfg.Aliases) > 0 {
			fmt.Println("Your custom aliases:")
			for name, cmd := range cfg.Aliases {
				fmt.Printf("  dockman %s → dockman %s\n",
					output.Cyan(name),
					output.Gray(cmd))
			}
		}

		fmt.Println()
		output.Info("Edit .dockman.yml anytime to update settings")

		return nil
	},
}

func init() {
	initCmd.Flags().BoolVarP(&initForce, "force", "f", false, "Overwrite existing config without prompting")
}

func valueOrNone(s string) string {
	if s == "" {
		return output.Gray("(none)")
	}
	return s
}

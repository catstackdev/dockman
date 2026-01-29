package cmd

import (
	"fmt"
	"os"

	"github.com/catstackdev/dockman/internal/config"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

var (
	configEdit bool
	configPath bool
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Show or edit project configuration",
	Long:  `Display current .dockman.yml configuration or open it in editor`,
	Example: `  dockman config            # Show current config
  dockman config --edit     # Open in $EDITOR
  dockman config --path     # Show config file path`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		projectDir := executor.ProjectDir
		cfgPath := fmt.Sprintf("%s/.dockman.yml", projectDir)

		// Show path only
		if configPath {
			fmt.Println(cfgPath)
			return nil
		}

		// Edit config
		if configEdit {
			editor := os.Getenv("EDITOR")
			if editor == "" {
				editor = "vim"
			}

			output.Info(fmt.Sprintf("Opening %s in %s...", cfgPath, editor))

			// Check if file exists, create if not
			if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
				output.Warning("Config doesn't exist. Run 'dockman init' first.")
				return nil
			}

			// Open in editor
			executor.runCommand("exec", editor, cfgPath)
			return nil
		}

		// Show config
		cfg, err := config.LoadProjectConfig(projectDir)
		if err != nil {
			output.Error(err.Error())
			return err
		}

		// Check if config exists
		if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
			output.Warning("No .dockman.yml found in project")
			output.Info("Run 'dockman init' to create one")
			return nil
		}

		// Display config
		output.Box("Project Configuration", []string{
			fmt.Sprintf("Location: %s", cfgPath),
			"",
			"Settings:",
			fmt.Sprintf("  Default Preset: %s", valueOrEmpty(cfg.DefaultPreset)),
			fmt.Sprintf("  Auto Pull:      %v", cfg.AutoPull),
			fmt.Sprintf("  Compose File:   %s", valueOrEmpty(cfg.ComposeFile)),
			fmt.Sprintf("  Aliases:        %d", len(cfg.Aliases)),
		})

		if len(cfg.Aliases) > 0 {
			fmt.Println("Aliases:")
			for name, cmd := range cfg.Aliases {
				fmt.Printf("  %s → %s\n",
					output.Cyan(name),
					output.Gray(cmd))
			}
			fmt.Println()
		}

		// Show raw YAML
		data, _ := yaml.Marshal(cfg)
		fmt.Println("Raw YAML:")
		fmt.Println(output.Gray(string(data)))

		return nil
	},
}

func init() {
	configCmd.Flags().BoolVarP(&configEdit, "edit", "e", false, "Open config in $EDITOR")
	configCmd.Flags().BoolVarP(&configPath, "path", "p", false, "Show config file path only")
}

func valueOrEmpty(s string) string {
	if s == "" {
		return output.Gray("(default)")
	}
	return s
}

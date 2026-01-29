package cmd

import (
	"fmt"
	"sort"

	"github.com/catstackdev/dockman/internal/config"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var aliasesCmd = &cobra.Command{
	Use:     "aliases",
	Short:   "List all custom aliases",
	Long:    `Display all custom command aliases defined in .dockman.yml`,
	Example: `  dockman aliases           # List all aliases`,
	RunE: func(cmd *cobra.Command, args []string) error {
		executor, err := getExecutor()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		cfg, err := config.LoadProjectConfig(executor.ProjectDir)
		if err != nil {
			output.Error(err.Error())
			return err
		}

		if len(cfg.Aliases) == 0 {
			output.Warning("No aliases configured")
			output.Info("Run 'dockman init' or edit .dockman.yml to add aliases")
			return nil
		}

		output.Box("Custom Aliases", []string{
			fmt.Sprintf("Found %d alias(es) in .dockman.yml", len(cfg.Aliases)),
		})

		// Sort aliases for consistent output
		names := make([]string, 0, len(cfg.Aliases))
		for name := range cfg.Aliases {
			names = append(names, name)
		}
		sort.Strings(names)

		// Display aliases
		for _, name := range names {
			command := cfg.Aliases[name]
			fmt.Printf("  %s → dockman %s\n",
				output.Cyan(name),
				output.Gray(command))
		}
		fmt.Println()

		return nil
	},
}

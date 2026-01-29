package cmd

import (
	"fmt"
	"sort"

	"github.com/catstackdev/dockman/internal/preset"
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:   "preset",
	Short: "Manage presets",
	Long:  `List and manage service presets`,
}

var presetListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all presets",
	RunE: func(cmd *cobra.Command, args []string) error {
		presetMgr, err := preset.NewManager()
		if err != nil {
			output.Error(err.Error())
			return err
		}

		presets := presetMgr.GetAll()

		if len(presets) == 0 {
			output.Warning("No presets configured")
			return nil
		}

		// Sort preset names
		names := make([]string, 0, len(presets))
		for name := range presets {
			names = append(names, name)
		}
		sort.Strings(names)

		// Print presets
		fmt.Print("\n📋 Available Presets:\n\n")
		for _, name := range names {
			p := presets[name]
			fmt.Printf("  %s %s\n", output.FormatPresetName(name), output.Gray(fmt.Sprintf("(%v)", p.Services)))
			if p.Description != "" {
				fmt.Printf("    %s\n", output.Gray(p.Description))
			}
		}
		fmt.Println()

		return nil
	},
}

func init() {
	presetCmd.AddCommand(presetListCmd)
}

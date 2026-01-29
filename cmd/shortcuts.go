package cmd

import (
	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

// Quick shortcuts for common operations

var stopCmd = &cobra.Command{
	Use:    "stop [services...]",
	Hidden: true, // Hide from main help (it's an alias)
	RunE: func(cmd *cobra.Command, args []string) error {
		return downCmd.RunE(cmd, args)
	},
}

var startCmd = &cobra.Command{
	Use:    "start [services...]",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return upCmd.RunE(cmd, args)
	},
}

var shCmd = &cobra.Command{
	Use:    "sh <service>",
	Hidden: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			output.Error("Service name required")
			return nil
		}
		return execCmd.RunE(cmd, args)
	},
}

package cmd

import (
	"fmt"
	"runtime"

	"github.com/catstackdev/dockman/pkg/output"
	"github.com/spf13/cobra"
)

var (
	Version   = "0.2.4"
	BuildDate = "unknown"
	GitCommit = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run: func(cmd *cobra.Command, args []string) {
		output.Box("Dockman Version", []string{
			fmt.Sprintf("Version:    %s", Version),
			fmt.Sprintf("Build Date: %s", BuildDate),
			fmt.Sprintf("Git Commit: %s", GitCommit),
			fmt.Sprintf("Go Version: %s", runtime.Version()),
			fmt.Sprintf("OS/Arch:    %s/%s", runtime.GOOS, runtime.GOARCH),
		})
	},
}

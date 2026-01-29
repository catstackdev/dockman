package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish]",
	Short: "Generate shell completion script",
	Long: `Generate shell completion script for dockman.

To load completions:

Bash:
  $ source <(dockman completion bash)
  # To load permanently:
  $ dockman completion bash > /usr/local/etc/bash_completion.d/dockman

Zsh:
  $ source <(dockman completion zsh)
  # To load permanently:
  $ dockman completion zsh > "${fpath[1]}/_dockman"

Fish:
  $ dockman completion fish | source
  # To load permanently:
  $ dockman completion fish > ~/.config/fish/completions/dockman.fish
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		}
	},
}

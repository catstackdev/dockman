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
	Args:                  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	RunE: func(cmd *cobra.Command, args []string) error {
		switch args[0] {
		case "bash":
			return cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			return cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			return cmd.Root().GenFishCompletion(os.Stdout, true)
		}
		return nil
	},
}

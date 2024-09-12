package git

import (
	"errors"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/git"
	"github.com/spf13/cobra"
)

func MakeGitPushCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "push",
		Short: "Pushes the trk data repository to the configured git remote",
		Long: `Pushes the trk data repository to the configured git remote.
	
	This is an alias for "trk git push"
	`,
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			if !git.CheckGitExists() {
				return errors.New("unable to locate git on your system")
			}
			return git.InvokeGitCommand("push")
		}),
	}

	return cmd
}

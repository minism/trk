package git

import (
	"errors"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/git"
	"github.com/spf13/cobra"
)

func MakeGitPassthroughCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "git",
		Short: "Invoke a git command on the trk data repository",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			if !git.CheckGitExists() {
				return errors.New("unable to locate git on your system")
			}

			// TODO: Flags need to be passed through here as well.
			return git.InvokeGitCommand(args...)
		}),
	}

	return cmd
}

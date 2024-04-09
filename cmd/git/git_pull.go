package git

import (
	"errors"

	"github.com/minism/trk/cmd/shared"
	"github.com/spf13/cobra"
)

func MakeGitPullCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pull",
		Short: "Pulls the trk data repository from the configured git remote",
		Long: `Pulls the trk data repository from the configured git remote.
	
	This is an alias for "trk git pull"
	`,
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			if !checkGitExists() {
				return errors.New("unable to locate git on your system")
			}
			return invokeGitCommand("pull")
		}),
	}

	return cmd
}

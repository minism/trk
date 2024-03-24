package tracking

import (
	"github.com/minism/trk/cmd/shared"
	"github.com/spf13/cobra"
)

func MakeTrackingSetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set <project> <num_hours>",
		Short: "Sets time tracking hours for a particular day",
		Long: `Sets time tracking hours for a particular day.
	
	This is an alias for "trk add --replace"
	`,
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			flagReplace = true
			return runAddCmd(cmd, args)
		}),
	}

	addSharedArgs(cmd)

	return cmd
}

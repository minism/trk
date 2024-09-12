package tracking

import (
	"github.com/minism/trk/cmd/shared"
	"github.com/spf13/cobra"
)

func MakeTrackingClearCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "clear <project>",
		Short: "Clears time tracking hours for a particular day",
		Long: `Clears time tracking hours for a particular day.
	
	This is an alias for "trk set <project> 0"
	`,
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			flagReplace = true
			return runAddCmd(cmd, []string{args[0], "0"})
		}),
	}

	cmd.Args = cobra.ExactArgs(1)
	addSharedArgs(cmd)

	return cmd
}

package tracking

import (
	"github.com/spf13/cobra"
)

var (
	flagMessage string
	flagDate    string
	flagReplace bool
)

func addSharedArgs(cmd *cobra.Command) {
	// Add flags to both commands, make set an alias with an override.
	cmd.Flags().StringVarP(&flagDate, "date", "d", "", "Update log for the given day, default to today.")
	cmd.Flags().StringVarP(&flagMessage, "message", "m", "", "Provide a message along with the entry.")
	cmd.Flags().BoolVar(&flagReplace, "replace", false, "Replaces all previous entries for that day.")
}

package tracking

import (
	"fmt"
	"log"

	"github.com/minism/trk/internal/git"
	"github.com/minism/trk/internal/storage"
	"github.com/spf13/cobra"
)

var (
	flagMessage string
	flagDate    string
	flagReplace bool
)

func commitIfEnabled(format string, args ...any) error {
	cfg, err := storage.LoadConfig()
	if err != nil {
		return err
	}
	msg := fmt.Sprintf(format, args...)
	if cfg.AutoCommit {
		git.InvokeGitCommand("commit", "-am", msg)
	} else {
		log.Print(msg)
	}
	return nil
}

func addSharedArgs(cmd *cobra.Command) {
	// Add flags to both commands, make set an alias with an override.
	cmd.Args = cobra.ExactArgs(2)
	cmd.Flags().StringVarP(&flagDate, "date", "d", "", "Update log for the given day, default to today.")
	cmd.Flags().StringVarP(&flagMessage, "message", "m", "", "Provide a message along with the entry.")
	cmd.Flags().BoolVar(&flagReplace, "replace", false, "Replaces all previous entries for that day.")
}

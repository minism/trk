package invoice

import (
	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/git"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

func MakeInvoiceDeleteCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete <invoice_id>",
		Short: "Delete invoices",
		Args:  cobra.ExactArgs(1),
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			id := model.ProjectInvoiceId(args[0])
			err := core.DeleteProjectInvoiceById(id)
			if err != nil {
				return err
			}

			// TODO: Rethink what should go to stderr and what should go to stdout for cases like this, use git as a guiding example.
			err = git.CommitIfEnabled("Deleted invoice %s\n", id)
			if err != nil {
				return err
			}

			return err
		}),
	}

	return cmd
}

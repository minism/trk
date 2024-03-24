package invoice

import (
	"fmt"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
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
			// TODO: Rethink what should go to stderr and what should go to stdout for cases like this, use git as a guiding example.
			fmt.Printf("Deleted invoice %s\n", id)
			return err
		}),
	}

	return cmd
}

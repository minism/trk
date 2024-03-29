package invoice

import (
	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

var (
	flagMarkSent bool
	flagMarkPaid bool
)

func MakeInvoiceUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <invoice_id>",
		Short: "Update invoice data",
		Args:  cobra.ExactArgs(1),
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			id := model.ProjectInvoiceId(args[0])
			invoice, err := core.FetchProjectInvoiceById(id)
			if err != nil {
				return err
			}

			// Apply the update.
			if flagMarkSent {
				invoice.IsSent = true
			}
			if flagMarkPaid {
				invoice.IsPaid = true
			}

			err = core.UpdateProjectInvoice(invoice)
			if err != nil {
				return err
			}

			display.PrintProjectInvoicesTable([]model.ProjectInvoice{invoice})

			return nil
		}),
	}

	cmd.Flags().BoolVar(&flagMarkSent, "sent", false, "Mark the invoice as sent")
	cmd.Flags().BoolVar(&flagMarkPaid, "paid", false, "Mark the invoice as paid")

	return cmd
}

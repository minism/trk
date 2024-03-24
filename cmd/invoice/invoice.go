package invoice

import (
	"github.com/spf13/cobra"
)

func MakeInvoiceCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "invoice",
		Aliases: []string{"invoices"},
		Short:   "View and manage invoices",
	}

	cmd.AddCommand(
		MakeInvoiceListCommand(),
		MakeInvoiceGenerateCommand(),
		MakeInvoiceDeleteCommand(),
	)

	return cmd
}

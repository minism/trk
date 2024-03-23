package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

func runSummaryCmd(cmd *cobra.Command, args []string) error {
	invoices, err := core.FetchAllProjectInvoices()
	if err != nil {
		return err
	}

	unpaidInvoices := model.FilterProjectInvoicesByUnpaid(invoices)
	if len(unpaidInvoices) > 0 {
		fmt.Printf("You have %d unpaid invoices totaling %s:\n\n", len(unpaidInvoices), display.ReadableMoney(0))
		display.PrintProjectInvoicesTable(unpaidInvoices)
	}

	return nil
}

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Aliases: []string{"status"},
	Short:   "Show a status overview (work-in-progres)",
	Run:     wrapCommand(runSummaryCmd),
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

func runSummaryCmd(cmd *cobra.Command, args []string) error {
	// Display outstanding/unpaid invoices.
	invoices, err := core.FetchAllProjectInvoices()
	if err != nil {
		return err
	}
	unpaidInvoices := model.FilterProjectInvoicesByUnpaid(invoices)
	if len(unpaidInvoices) > 0 {
		fmt.Printf("You have %d unpaid invoices totaling %s:\n\n", len(unpaidInvoices), display.ReadableMoney(model.SumInvoices(unpaidInvoices)))
		display.PrintProjectInvoicesTable(unpaidInvoices)
		fmt.Println()
	}

	// Display uninvoiced hours.
	uninvoicedEntries := make([]model.LogEntry, 0)
	uninvoicedTotal := 0.0
	projects, err := core.FetchAllProjects()
	if err != nil {
		return err
	}
	for _, project := range projects {
		entries, err := core.FetchLogEntriesForProject(project)
		if err != nil {
			return err
		}
		for _, invoice := range invoices {
			if invoice.Project.Equals(&project) {
				entries = model.ExcludeLogEntriesBetween(entries, invoice.StartDate, invoice.EndDate)
			}
		}
		uninvoicedTotal += model.GetTotalHours(entries) * project.HourlyRate
		uninvoicedEntries = append(uninvoicedEntries, entries...)
	}
	if len(uninvoicedEntries) > 0 {
		uninvoicedHours := model.GetTotalHours(uninvoicedEntries)

		fmt.Printf(
			"You have %s uninvoiced hours totaling %s:\n\n",
			display.FormatFloatMinDecimal(uninvoicedHours),
			display.ReadableMoney(uninvoicedTotal))
		display.PrintLogEntryTable(uninvoicedEntries)
	}

	return nil
}

var summaryCmd = &cobra.Command{
	Use:     "summary",
	Aliases: []string{"sum", "status"},
	Short:   "Show a status overview (work-in-progres)",
	Run:     wrapCommand(runSummaryCmd),
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}

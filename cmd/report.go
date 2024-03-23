package cmd

import (
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/report"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

func runReportCmd(cmd *cobra.Command, args []string) error {
	// For now this just shows an income summary.
	year := util.TrkToday().Year()
	reports, err := report.GenerateMonthlyIncomeForYear(year)
	if err != nil {
		return err
	}

	display.PrintIncomeReportTable(reports)

	return nil
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generate reports (work-in-progres)",
	Run:   wrapCommand(runReportCmd),
}

func init() {
	rootCmd.AddCommand(reportCmd)
}

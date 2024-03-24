package report

import (
	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/report"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

func MakeReportCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "report",
		Short: "Generate reports (work-in-progres)",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			// For now this just shows an income summary.
			year := util.TrkToday().Year()
			reports, err := report.GenerateMonthlyIncomeForYear(year)
			if err != nil {
				return err
			}

			display.PrintIncomeReportTable(reports)

			return nil
		}),
	}

	return cmd
}

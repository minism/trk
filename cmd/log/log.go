/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package log

import (
	"fmt"
	"log"
	"time"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/util"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

var (
	flagDate            string
	flagSince           string
	flagAll             bool
	flagDisplayWeekly   bool
	flagDisplayCombined bool
	flagInvoicePeriod   bool
)

func MakeLogCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "log [project]",
		Short: "Display a summary of the work log",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			projects, err := core.FetchAllProjects()
			if err != nil {
				return err
			}

			// Determine date range.
			// By default we show the last two weeks.
			// TODO: Throw combination incompatibility errors here.
			from := util.MinDate
			to := util.MaxDate
			// TODO: Weekly should work with --since, but snap to beginning of that week.
			if flagAll || flagDisplayWeekly {
				from = util.MinDate
			} else if len(flagDate) > 0 {
				from, err = util.ParseNaturalDate(flagDate)
				if err != nil {
					return err
				}
				to = from.Add(time.Duration(24) * time.Hour)
			} else if flagInvoicePeriod {
				from = util.GetPrevBimonthlyDate(util.TrkToday())
				to = util.GetNextBimonthlyDate(util.TrkToday())

			} else if len(flagSince) > 0 {
				from, err = util.ParseNaturalDate(flagSince)
				if err != nil {
					return err
				}
			}

			// Optionally filter by a single project.
			projectFilter := shared.FlagProject
			if len(args) > 0 {
				projectFilter = args[0]
			}
			if projectFilter != "" {
				project, err := core.FilterProjectsByIdFuzzy(projects, projectFilter)
				if err != nil {
					return err
				}
				projects = []model.Project{project}
			}

			// Fetch entries.
			entries, err := core.FetchAllLogEntries(projects)
			if err != nil {
				return err
			}

			// Apply any other filters.
			entries = model.MergeAndSortLogEntries(entries)
			entries = model.FilterLogEntriesBetween(entries, from, to)

			// Output format.
			log.Printf("Showing logs since %s\n\n", display.ReadableDate(from))
			if flagDisplayWeekly {
				byProject := model.GroupLogEntriesByProject(entries)
				for _, project := range projects {
					log.Printf("Project: %s\n", display.ColorProject(project.Name))
					byWeek := model.GroupLogEntriesByWeekStart(byProject[project.ID()])
					display.PrintWeeklyLogEntryTable(byWeek)
					fmt.Println()
				}
			} else if flagDisplayCombined {
				// TODO: Handle both -c and -w
				combinedEntries := model.CombineLogEntriesByProject(entries)
				display.PrintCombinedLogEntryTable(combinedEntries)
			} else {
				display.PrintLogEntryTable(entries)
			}

			return nil
		}),
	}

	cmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all log history")
	cmd.Flags().StringVarP(&flagDate, "date", "d", "", "Only show logs for the given day")
	cmd.Flags().StringVar(&flagSince, "since", "last week", "Only show logs since the given date")
	cmd.Flags().BoolVarP(&flagInvoicePeriod, "invoice-period", "i", false, "Only show logs for the current invoice period (assumes bimonthly)")
	cmd.Flags().BoolVarP(&flagDisplayWeekly, "weekly", "w", false, "Show weekly aggregated logs")
	cmd.Flags().BoolVarP(&flagDisplayCombined, "combined", "c", false, "Show combined project logs")

	return cmd
}

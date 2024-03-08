/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagSince         string
	flagAll           bool
	flagDisplayWeekly bool
)

func runLogCmd(cmd *cobra.Command, args []string) error {
	projects, err := core.GetAllProjects()
	if err != nil {
		return err
	}

	// Determine date range.
	// By default we show the last two weeks.
	from := time.Now().Add(time.Duration(-24*14) * time.Hour)
	to := util.MaxDate
	if flagAll || flagDisplayWeekly {
		from = util.MinDate
	} else if len(flagSince) > 0 {
		from, err = util.ParseNaturalDate(flagSince)
		if err != nil {
			return err
		}
	}

	// Optionally filter by a single project.
	projectFilter := flagProject
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
	entries, err := core.RetrieveAllLogEntries(projects)
	if err != nil {
		return err
	}

	// Apply any other filters.
	entries = model.MergeAndSortLogEntries(entries)
	entries = model.FilterLogEntriesBetween(entries, from, to)

	// Output format.
	log.Printf("Showing logs since %s\n", display.ReadableDate(from))
	if flagDisplayWeekly {
		byProject := model.GroupLogEntriesByProject(entries)
		for projectId, entries := range byProject {
			byWeek := model.GroupLogEntriesByYearWeek(entries)
			log.Printf("Project: %s\n", display.ColorProject(projectId))
			display.PrintWeeklyLogEntryTable(byWeek)
			fmt.Println()
		}
	} else {
		display.PrintLogEntryTable(entries)
	}

	return nil
}

var logCmd = &cobra.Command{
	Use:   "log [project]",
	Short: "Display a summary of the work log",
	Run:   wrapCommand(runLogCmd),
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVarP(&flagAll, "all", "a", false, "Show all log history.")
	logCmd.Flags().BoolVarP(&flagDisplayWeekly, "weekly", "w", false, "Show weekly as opposed to daily output.")
	logCmd.Flags().StringVar(&flagSince, "since", "", "Only show logs since the given date")
}

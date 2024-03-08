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
	flagDisplayWeekly bool
)

func runLogCmd(cmd *cobra.Command, args []string) error {
	projects, err := core.GetAllProjects()
	if err != nil {
		return err
	}

	// Flag parsing.
	since := time.Unix(0, 0)
	if len(flagSince) > 0 {
		since, err = util.ParseNaturalDate(flagSince)
		if err != nil {
			return err
		}
		log.Printf("Filtering since %s", since)
	}

	// Optionally filter by a single project.
	if flagProject != "" {
		project, err := core.FilterProjectsByIdFuzzy(projects, flagProject)
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
	entries = model.FilterLogEntriesSince(entries, since)

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
	Use:   "log",
	Short: "Display a summary of the work log",
	Run:   wrapCommand(runLogCmd),
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVarP(&flagDisplayWeekly, "weekly", "w", false, "Show weekly as opposed to daily output.")
	logCmd.Flags().StringVar(&flagSince, "since", "", "Only show logs since the given date")
}

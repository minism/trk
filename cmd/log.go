/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"time"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagProject string
	flagSince   string
	flagWeekly  bool
)

func run(cmd *cobra.Command, args []string) {
	projects, err := core.GetAllProjects()
	if err != nil {
		log.Fatal(err)
	}

	// Flag parsing.
	since := time.Unix(0, 0)
	if len(flagSince) > 0 {
		since, err = util.ParseNaturalDate(flagSince)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Filtering since %s", since)
	}

	// Optionally filter by a single project.
	if flagProject != "" {
		project, err := core.FilterProjectsByIdFuzzy(projects, flagProject)
		if err != nil {
			// TODO: Share the error handling which dumps project IDs here.
			log.Fatal(err)
		}
		projects = []model.Project{project}
	}

	// Fetch entries.
	entries, err := core.RetrieveMergedLogEntries(projects)
	if err != nil {
		log.Fatal(err)
	}

	// Apply any other filters.
	entries = core.FilterLogEntriesSince(entries, since)

	display.PrintLogEntryTable(entries)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Display a summary of the work log",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().BoolVarP(&flagWeekly, "weekly", "w", false, "Show weekly as opposed to daily output.")
	logCmd.Flags().StringVarP(&flagProject, "project", "p", "", "Filter by a particular project ID (fuzzy match).")
	logCmd.Flags().StringVar(&flagSince, "since", "", "Only show logs since the given date")
}

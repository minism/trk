/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/spf13/cobra"
)

var (
	flagProject string
	flagSince   string
)

func run(cmd *cobra.Command, args []string) {
	projects, err := core.GetAllProjects()
	if err != nil {
		log.Fatal(err)
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

	entries, err := core.RetrieveMergedLogEntries(projects)
	if err != nil {
		log.Fatal(err)
	}
	display.PrintLogEntryTable(entries)
}

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "Display a summary of the work log",
	Run:   run,
}

func init() {
	rootCmd.AddCommand(logCmd)
	logCmd.Flags().StringVarP(&flagProject, "project", "p", "", "Filter by a particular project ID (fuzzy match).")
	logCmd.Flags().StringVar(&flagProject, "since", "", "Only show logs since the given date")
}

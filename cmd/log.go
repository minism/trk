/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

var (
	flagProject string
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := core.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}

		// Optionally filter.
		if flagProject != "" {
			project, err := core.FilterProjectsByIdFuzzy(projects, flagProject)
			if err != nil {
				// TODO: Share the error handling which dumps project IDs here.
				log.Fatal(err)
			}
			projects = []model.Project{project}
		}

		for _, p := range projects {
			fmt.Printf("Project %s\n", display.ColorProject(p.ID()))
			tbl := table.New("Date", "Hours", "Note")
			tbl.WithFirstColumnFormatter(display.ColorDate)
			entries, err := core.RetrieveLogEntries(p.ID())
			if err != nil {
				log.Fatal(err)
			}
			for _, entry := range entries {
				note := display.ColorNull("(none)")
				if len(entry.Note) > 0 {
					note = entry.Note
				}
				tbl.AddRow(
					display.ReadableDateWithoutColor(entry.Date), entry.Hours, note)
			}
			tbl.Print()
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)

	logCmd.Flags().StringVarP(&flagProject, "project", "p", "", "Filter by a particular project ID (fuzzy match).")
}

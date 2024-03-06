/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   "log",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		// Print all projects for now.
		projects, err := core.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}

		for _, p := range projects {
			fmt.Printf("Project %s\n", display.ColorProject(p.ID()))
			tbl := table.New("Date", "Hours", "Note")
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
					display.ReadableDate(entry.Date), entry.Hours, note)
			}
			tbl.Print()
			fmt.Println()
		}
	},
}

func init() {
	rootCmd.AddCommand(logCmd)
}

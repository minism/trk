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

// projectsCmd represents the projects command
var projectsCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"projects"},
	Short:   "View and manage projects",
	Run: func(cmd *cobra.Command, args []string) {
		projects, err := core.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}

		tbl := table.New("id", "name", "hourly rate")
		tbl.WithFirstColumnFormatter(display.ColorProject)
		for _, p := range projects {
			tbl.AddRow(p.ID(), p.Name, display.ReadableMoney(p.HourlyRate))
		}

		fmt.Printf("All projects:\n\n")
		tbl.Print()
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
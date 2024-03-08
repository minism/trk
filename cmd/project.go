/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func runProjectCmd(cmd *cobra.Command, args []string) error {
	projects, err := core.GetAllProjects()
	if err != nil {
		return err
	}

	tbl := table.New("id", "name", "hourly rate")
	tbl.WithFirstColumnFormatter(display.ColorProject)
	for _, p := range projects {
		tbl.AddRow(p.ID(), p.Name, display.ReadableMoney(p.HourlyRate))
	}

	fmt.Printf("All projects:\n\n")
	tbl.Print()

	return nil
}

var projectsCmd = &cobra.Command{
	Use:     "project",
	Aliases: []string{"projects"},
	Short:   "View and manage projects",
	Run:     wrapCommand(runProjectCmd),
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}

/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package project

import (
	"fmt"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/rodaine/table"
	"github.com/spf13/cobra"
)

func MakeProjectCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "project",
		Aliases: []string{"projects"},
		Short:   "View and manage projects",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			projects, err := core.FetchAllProjects()
			if err != nil {
				return err
			}

			tbl := table.New("id", "name", "hourly rate", "invoice interval")
			tbl.WithFirstColumnFormatter(display.ColorProject)
			for _, p := range projects {
				tbl.AddRow(p.ID(), p.Name, display.ReadableMoney(p.HourlyRate), p.InvoiceInterval)
			}

			fmt.Printf("All projects:\n\n")
			tbl.Print()

			return nil
		}),
	}

	return cmd
}

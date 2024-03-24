/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package tracking

import (
	"fmt"
	"strconv"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/util"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

func MakeTrackingAddCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add <project> <num_hours>",
		Short: "Adds time tracking hours to a project",
		Run:   shared.WrapCommand(runAddCmd),
	}

	addSharedArgs(cmd)

	return cmd
}

func runAddCmd(cmd *cobra.Command, args []string) error {
	var err error

	// Parse date.
	date := util.TrkToday()
	if len(flagDate) > 0 {
		date, err = util.ParseNaturalDate(flagDate)
		if err != nil {
			return err
		}
	}

	// Match project.
	projects, err := core.FetchAllProjects()
	if err != nil {
		return err
	}
	project, err := core.FilterProjectsByIdFuzzy(projects, args[0])
	if err != nil {
		return err
	}

	// Parse hours.
	hours, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return fmt.Errorf("invalid hours value: %w", err)
	}

	// Update log entry.
	entry, err := core.MakeLogEntry(project, date, hours, flagMessage)
	if err != nil {
		return err
	}
	allDayEntries, err := core.AppendLogEntry(entry, flagReplace)
	if err != nil {
		return err
	}

	// Display change and total.
	total := model.GetTotalHours(allDayEntries)
	if total == hours {
		fmt.Printf(
			"Logged %s to project %s for %s\n",
			display.ReadableHours(hours),
			display.ColorProject(project.ID()),
			display.ReadableDate(date))
	} else {
		fmt.Printf(
			"Logged %s to project %s\n",
			display.ReadableHours(hours),
			display.ColorProject(project.ID()))
		fmt.Printf(
			"\nYou have %s total for %s\n",
			display.ReadableHours(total),
			display.ReadableDate(date))
	}

	return nil
}

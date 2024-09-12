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
	"github.com/minism/trk/internal/git"
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

	cmd.Args = cobra.ExactArgs(2)
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
	entry := model.MakeLogEntry(project, date, hours, flagMessage)
	allDayEntries, err := core.AppendLogEntry(entry, flagReplace)
	if err != nil {
		return err
	}

	// Display change and total.
	total := model.GetTotalHours(allDayEntries)
	var msg string
	if total == hours {
		msg = fmt.Sprintf(
			"Logged %s to project %s for %s",
			display.ReadableHours(hours),
			display.ColorProject(project.ID()),
			display.ReadableDate(date))
	} else {
		msg = fmt.Sprintf(
			"Logged %s to project %s\nYou have %s total for %s",
			display.ReadableHours(hours),
			display.ColorProject(project.ID()),
			display.ReadableHours(total),
			display.ReadableDate(date))
	}
	git.CommitIfEnabled(msg)

	return nil
}

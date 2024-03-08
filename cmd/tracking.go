/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagMessage string
	flagDate    string
	flagReplace bool
)

func runAddCmd(cmd *cobra.Command, args []string) error {
	var err error

	// Parse date.
	// TODO: Revisit timezone stuff.
	date := util.UtcToday()
	if len(flagDate) > 0 {
		date, err = util.ParseNaturalDate(flagDate)
		if err != nil {
			return err
		}
	}

	// Match project.
	projects, err := core.GetAllProjects()
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

func runSetCommand(cmd *cobra.Command, args []string) error {
	flagReplace = true
	addCmd.Run(cmd, args)
	return nil
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <project> <num_hours>",
	Short: "Adds time tracking hours to a project",
	Run:   wrapCommand(runAddCmd),
}

// set is just an alias for add with --replace set to true.
var setCmd = &cobra.Command{
	Use:   "set <project> <num_hours>",
	Short: "Sets time tracking hours for a particular day",
	Long: `Sets time tracking hours for a particular day.

This is an alias for "trk add --replace"
`,
	Run: wrapCommand(runSetCommand),
}

func init() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(setCmd)

	// Add flags to both commands, make set an alias with an override.
	for _, cmd := range []*cobra.Command{addCmd, setCmd} {
		cmd.Args = cobra.ExactArgs(2)
		cmd.Flags().StringVarP(&flagDate, "date", "d", "", "Update log for the given day, default to today.")
		cmd.Flags().StringVarP(&flagMessage, "message", "m", "", "Provide a message along with the entry.")
		cmd.Flags().BoolVar(&flagReplace, "replace", false, "Replaces all previous entries for that day.")
	}
}

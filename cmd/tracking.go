/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/spf13/cobra"
)

var (
	flagMessage string
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add <project> <num_hours>",
	Short: "Adds time tracking hours to a project",
	Run: func(cmd *cobra.Command, args []string) {
		// Match project.
		projects, err := core.GetAllProjects()
		if err != nil {
			log.Fatal(err)
		}
		project, err := core.FilterProjectsByIdFuzzy(projects, args[0])
		if err != nil {
			log.Println(err)
			if errors.Is(err, core.ErrMultipleProjectsMatched) {
				printAllProjects(projects)
			}
			os.Exit(1)
		}

		// Parse hours.
		hours, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal("Invalid hours value: ", err)
		}

		// Parse date.
		date := time.Now()

		// Update log entry.
		entry, err := core.MakeLogEntry(project, date, hours, flagMessage)
		if err != nil {
			log.Fatal(err)
		}
		allEntries, err := core.AppendLogEntry(entry)
		if err != nil {
			log.Fatal(err)
		}

		// Display change and total.
		fmt.Printf(
			"Logged %s to project %s\n",
			display.ReadableHours(hours),
			display.ColorProject(project.ID()))
		fmt.Printf(
			"\nYou have %s total for %s\n",
			display.ReadableHours(core.GetTotalHours(allEntries)),
			display.ReadableDate(date))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Args = cobra.ExactArgs(2)
	addCmd.Flags().StringVarP(&flagMessage, "message", "m", "", "Provide a message along with the entry.")
}

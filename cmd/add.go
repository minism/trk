/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
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
		projectId := args[0]
		hours, err := strconv.ParseFloat(args[1], 64)
		if err != nil {
			log.Fatal("Invalid hours value: ", err)
		}

		date := time.Now()
		entry, err := core.MakeValidLogEntry(projectId, date, hours, flagMessage)
		if err != nil {
			log.Fatal(err)
		}
		err = core.AppendLogEntry(entry)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf(
			"%s to project %s on %s\n",
			display.ColorSuccess("Logged %.2f hours", hours),
			display.ColorProject(projectId),
			display.ReadableDate(date))
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Args = cobra.ExactArgs(2)
	addCmd.Flags().StringVarP(&flagMessage, "message", "m", "", "Provide a message along with the entry.")
}

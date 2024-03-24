/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	flagProject string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "trk",
	Short: "trk is a time-tracking and invoicing tool",
	Long:  `trk is a time-tracking and invoicing tool.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&flagProject, "project", "p", "", "Filter by a particular project ID (fuzzy match).")

	// Configure the logger.
	log.SetFlags(0)
}

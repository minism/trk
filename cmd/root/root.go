/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package root

import (
	"log"
	"os"

	"github.com/minism/trk/cmd/git"
	"github.com/minism/trk/cmd/initialize"
	"github.com/minism/trk/cmd/invoice"
	cmdlog "github.com/minism/trk/cmd/log"
	"github.com/minism/trk/cmd/project"
	"github.com/minism/trk/cmd/report"
	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/cmd/summary"
	"github.com/minism/trk/cmd/tracking"
	"github.com/minism/trk/cmd/version"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "trk",
	Short: "trk is a time-tracking and invoicing tool",
	Long:  `trk is a time-tracking and invoicing tool.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.AddCommand(
		initialize.MakeInitCommand(),
		invoice.MakeInvoiceCommand(),
		cmdlog.MakeLogCommand(),
		project.MakeProjectCommand(),
		report.MakeReportCommand(),
		summary.MakeSummaryCommand(),
		tracking.MakeTrackingAddCommand(),
		tracking.MakeTrackingSetCommand(),
		tracking.MakeTrackingClearCommand(),
		version.MakeVersionCommand(),
		git.MakeGitPassthroughCommand(),
		git.MakeGitPushCommand(),
		git.MakeGitPullCommand(),
	)

	RootCmd.PersistentFlags().StringVarP(&shared.FlagProject, "project", "p", "", "Filter by a particular project ID (fuzzy match).")

	// Configure the logger.
	log.SetFlags(0)
}

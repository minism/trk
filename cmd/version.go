package cmd

import (
	"log"

	"github.com/minism/trk/internal/version"
	"github.com/spf13/cobra"
)

func runVersionCmd(cmd *cobra.Command, args []string) error {
	log.Printf("trk version %s\n", version.GetVersion())
	return nil
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show version information",
	Run:   wrapCommand(runVersionCmd),
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

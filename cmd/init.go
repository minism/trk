package cmd

import (
	"fmt"
	"log"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a trk config in your home directory",
	Run: func(cmd *cobra.Command, args []string) {
		created, err := core.InitTrk()
		if err != nil {
			log.Fatal(err)
		}
		configPath := config.GetConfigPath()
		if created {
			fmt.Printf("%s at %s\n", display.ColorSuccess("Initialized trk config"), configPath)
		} else {
			fmt.Println("Already initialized trk config at", configPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

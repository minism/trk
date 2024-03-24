package initialize

import (
	"fmt"
	"os"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/internal/util"
	"github.com/spf13/cobra"
)

var (
	flagForceReset bool
)

func MakeInitCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize a trk config in your home directory",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			if flagForceReset && !util.AskUserForConfirmation("Are you sure you want to reset the config?") {
				os.Exit(0)
			}

			created, err := core.InitTrk(flagForceReset)
			if err != nil {
				return err
			}
			configPath := config.GetConfigPath()
			if created {
				fmt.Printf("%s at %s\n", display.ColorSuccess("Initialized trk config"), configPath)
			} else {
				fmt.Println("Already initialized trk config at", configPath)
			}
			return nil
		}),
	}

	cmd.Flags().BoolVar(&flagForceReset, "force-reset", false, "Forcibly resets the config.")

	return cmd
}

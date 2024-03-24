package version

import (
	"log"

	"github.com/minism/trk/cmd/shared"
	"github.com/minism/trk/internal/version"
	"github.com/spf13/cobra"
)

func MakeVersionCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Show version information",
		Run: shared.WrapCommand(func(cmd *cobra.Command, args []string) error {
			log.Printf("trk version %s\n", version.GetVersion())
			return nil
		}),
	}

	return cmd
}

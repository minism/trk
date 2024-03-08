package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/spf13/cobra"
)

func wrapCommand(handler func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := handler(cmd, args); err != nil {
			log.Println(err)
			if errors.Is(err, core.ErrMultipleProjectsMatched) {
				projects, innerErr := core.GetAllProjects()
				if innerErr != nil {
					panic(innerErr)
				}
				display.PrintProjects(projects)
			}
			os.Exit(1)
		}
	}
}

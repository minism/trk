package shared

import (
	"errors"
	"log"
	"os"

	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/display"
	"github.com/minism/trk/pkg/model"
	"github.com/spf13/cobra"
)

var (
	FlagProject string
)

func WrapCommand(handler func(cmd *cobra.Command, args []string) error) func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {
		if err := handler(cmd, args); err != nil {
			log.Println(err)
			if errors.Is(err, core.ErrMultipleProjectsMatched) {
				projects, innerErr := core.FetchAllProjects()
				if innerErr != nil {
					panic(innerErr)
				}
				display.PrintProjects(projects)
			}
			os.Exit(1)
		}
	}
}

// Get the filtered set of projects for a command based on the global flag.
func GetFilteredProjects() ([]model.Project, error) {
	projects, err := core.FetchAllProjects()
	if err != nil {
		return nil, err
	}

	// Optionally filter by a single project.
	if FlagProject != "" {
		project, err := core.FilterProjectsByIdFuzzy(projects, FlagProject)
		if err != nil {
			return nil, err
		}
		projects = []model.Project{project}
	}

	return projects, nil
}

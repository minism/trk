package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/model"
)

// TODO: Move to display
func printAllProjects(projects []model.Project) {
	fmt.Println("Projects:")
	for _, project := range projects {
		fmt.Println(" - ", project.ID())
	}
}

package cmd

import (
	"fmt"

	"github.com/minism/trk/internal/model"
)

func printAllProjects(projects []model.Project) {
	fmt.Println("Projects:")
	for _, project := range projects {
		fmt.Println(" - ", project.ID())
	}
}

package display

import (
	"fmt"

	"github.com/minism/trk/pkg/model"
)

// TODO: Move to display
func PrintProjects(projects []model.Project) {
	fmt.Println("Projects:")
	for _, project := range projects {
		fmt.Println(" - ", project.ID())
	}
}

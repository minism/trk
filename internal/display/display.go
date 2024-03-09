package display

import (
	"fmt"

	"github.com/minism/trk/pkg/model"
)

func PrintProjects(projects []model.Project) {
	fmt.Println("Projects:")
	for _, project := range projects {
		fmt.Println(" - ", project.ID())
	}
}

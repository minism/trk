package core

import (
	"fmt"
	"strings"

	"github.com/minism/trk/pkg/model"
)

func FilterProjectsByIdFuzzy(projects []model.Project, query string) (model.Project, error) {
	matches := make([]model.Project, 0)
	for _, project := range projects {
		if strings.Contains(project.ID(), strings.ToLower(query)) {
			matches = append(matches, project)
		}
	}
	if len(matches) < 1 {
		return model.Project{}, fmt.Errorf("%w: %s", ErrProjectNotFound, query)
	} else if len(matches) > 1 {
		return model.Project{}, ErrMultipleProjectsMatched
	}
	return matches[0], nil
}

package core

import (
	"fmt"
	"strings"

	"github.com/minism/trk/internal/storage"
	"github.com/minism/trk/pkg/model"
)

func FetchAllProjects() ([]model.Project, error) {
	cfg, err := storage.LoadConfig()
	if err != nil {
		return nil, err
	}

	projects := []model.Project{}
	for _, cfgProject := range cfg.Projects {
		projects = append(projects, model.Project{ProjectConfig: cfgProject})
	}

	return projects, nil
}

func FetchProjectById(id string) (model.Project, error) {
	projects, err := FetchAllProjects()
	if err != nil {
		return model.Project{}, err
	}
	for _, project := range projects {
		if project.ID() == id {
			return project, nil
		}
	}
	return model.Project{}, fmt.Errorf("%w: %s", ErrProjectNotFound, id)
}

func ValidateProjectId(id string) error {
	_, err := FetchProjectById(id)
	return err
}

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

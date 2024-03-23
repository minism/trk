package core

import (
	"fmt"

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

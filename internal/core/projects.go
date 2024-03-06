package core

import (
	"fmt"

	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/storage"
)

func GetAllProjects() ([]model.Project, error) {
	projects := []model.Project{}
	cfg, err := storage.LoadConfig()
	if err != nil {
		return projects, err
	}

	for _, cfgProject := range cfg.Projects {
		projects = append(projects, model.Project{ProjectConfig: cfgProject})
	}

	return projects, nil
}

func GetProjectById(id string) (model.Project, error) {
	projects, err := GetAllProjects()
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

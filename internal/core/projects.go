package core

import (
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/storage"
)

func GetProjects() ([]model.Project, error) {
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

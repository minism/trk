package core

import (
	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/storage"
)

func GetProjects() ([]config.ProjectConfig, error) {
	cfg, err := storage.LoadConfig()
	if err != nil {
		return []config.ProjectConfig{}, err
	}
	return cfg.Projects, nil
}

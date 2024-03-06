package model

import (
	"strings"

	"github.com/minism/trk/internal/config"
)

type Project struct {
	config.ProjectConfig
}

func (p *Project) ID() string {
	id := strings.ToLower(p.ProjectConfig.Name)
	id = strings.Join(strings.Fields(id), "-")
	return id
}

func (p *Project) WorkLogPath() string {
	return config.GetProjectWorkLogPath(p.ID())
}

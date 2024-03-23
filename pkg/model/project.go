package model

import (
	"strings"

	"github.com/minism/trk/internal/config"
)

type Project struct {
	config.ProjectConfig
}

func (p *Project) Equals(other *Project) bool {
	return p.ID() == other.ID()
}

func (p *Project) ID() string {
	id := strings.ToLower(p.ProjectConfig.Name)
	id = strings.Join(strings.Fields(id), "-")
	return id
}

func (p *Project) WorkLogPath() string {
	return config.GetProjectWorkLogPath(p.ID())
}

func (p *Project) InvoicesPath() string {
	return config.GetProjectInvoicesPath(p.ID())
}

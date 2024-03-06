package core

import (
	"time"

	"github.com/minism/trk/internal/model"
)

func MakeValidLogEntry(projectId string, date time.Time, hours float64, note string) (model.LogEntry, error) {
	project, err := GetProjectById(projectId)
	if err != nil {
		return model.LogEntry{}, err
	}
	return model.LogEntry{
		ProjectId: project.ID(),
		Date:      date,
		Hours:     hours,
		Note:      note,
	}, nil
}

func AppendLogEntry(logEntry model.LogEntry) error {
	return nil
}

func UpdateLogEntry(logEntry model.LogEntry) error {
	return nil
}

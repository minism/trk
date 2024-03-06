package core

import (
	"sort"
	"time"

	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/storage"
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

// Retrieve log entries for the given project and date.
func RetrieveLogEntries(projectId string) ([]model.LogEntry, error) {
	project, err := GetProjectById(projectId)
	if err != nil {
		return nil, err
	}

	entries, err := storage.LoadProjectLogEntries(project)
	if err != nil {
		return nil, err
	}

	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Date.Before(entries[j].Date)
	})

	return entries, nil
}

// Appends the given log entry to storage.
func AppendLogEntry(logEntry model.LogEntry) error {
	return nil
}

// Sets the log entry for the given project+day, potentially overwriting any
// previous ones.
func SetDayLogEntry(logEntry model.LogEntry) error {
	return nil
}

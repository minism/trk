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
func AppendLogEntry(entry model.LogEntry) error {
	entries, err := RetrieveLogEntries(entry.ProjectId)
	if err != nil {
		return err
	}

	// Just append, for now we don't care about order.
	entries = append(entries, entry)

	// Write all to storage.
	err = storage.SaveProjectLogEntries(entry.ProjectId, entries)
	if err != nil {
		return err
	}
	return nil
}

// Sets the log entry for the given project+day, potentially overwriting any
// previous ones.
func SetDayLogEntry(entry model.LogEntry) error {
	entries, err := RetrieveLogEntries(entry.ProjectId)
	if err != nil {
		return err
	}

	// Remove entries that match the date of the passed-in entry.
	filteredEntries := make([]model.LogEntry, 0)
	for _, e := range entries {
		if !e.Date.Equal(entry.Date) {
			filteredEntries = append(filteredEntries, e)
		}
	}
	entries = filteredEntries

	// Finally append the new one.
	entries = append(entries, entry)

	// Write all to storage.
	err = storage.SaveProjectLogEntries(entry.ProjectId, entries)
	if err != nil {
		return err
	}

	return nil
}

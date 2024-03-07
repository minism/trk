package core

import (
	"time"

	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/storage"
	"github.com/minism/trk/internal/util"
)

func MakeLogEntry(project model.Project, date time.Time, hours float64, note string) (model.LogEntry, error) {
	return model.LogEntry{
		Project: project,
		Date:    date,
		Hours:   hours,
		Note:    note,
	}, nil
}

// Retrieve sorted log entries for a given project.
func RetrieveLogEntries(project model.Project) ([]model.LogEntry, error) {
	return RetrieveMergedLogEntries([]model.Project{project})
}

// Retrieve merged and sorted log entries for the given projects.
func RetrieveMergedLogEntries(projects []model.Project) ([]model.LogEntry, error) {
	entries := make([]model.LogEntry, 0)

	// Concatenate all entries.
	for _, project := range projects {
		projectEntries, err := storage.LoadProjectLogEntries(project)
		if err != nil {
			return nil, err
		}
		entries = append(entries, projectEntries...)
	}

	// Sort by time.
	SortLogEntries(entries)

	// Merge entries for the same day.
	ret := make([]model.LogEntry, 0)
	last := model.LogEntry{Date: time.Unix(0, 0)}
	for _, entry := range entries {
		if !last.Date.Equal(entry.Date) {
			if last.Hours > 0 {
				ret = append(ret, last)
			}
			last = entry
		} else {
			last.Hours += +entry.Hours
		}
	}
	ret = append(ret, last)

	return ret, nil
}

// Appends the given log entry to storage.
// Returns all entries for the day.
func AppendLogEntry(entry model.LogEntry, erasePreviousForDay bool) ([]model.LogEntry, error) {
	entries, err := RetrieveLogEntries(entry.Project)
	if err != nil {
		return nil, err
	}

	if erasePreviousForDay {
		// If set, remove entries that match the date of the passed-in entry.
		filteredEntries := make([]model.LogEntry, 0)
		for _, e := range entries {
			if !util.IsSameDay(e.Date, entry.Date) {
				filteredEntries = append(filteredEntries, e)
			}
		}
		entries = filteredEntries
	}

	// Validate the hour total.
	if entry.Hours+GetTotalHours(entries) > 24 {
		return nil, ErrHoursExceedsLimit
	}

	// Just append, for now we don't care about order.
	entries = append(entries, entry)

	// Write all to storage.
	err = storage.SaveProjectLogEntries(entry.Project.ID(), entries)
	if err != nil {
		return nil, err
	}
	return FilterLogEntriesByDay(entries, entry.Date), nil
}

package core

import (
	"github.com/minism/trk/internal/storage"
	"github.com/minism/trk/internal/util"
	"github.com/minism/trk/pkg/model"
)

// Retrieve sorted log entries for a given project.
func FetchLogEntriesForProject(project model.Project) ([]model.LogEntry, error) {
	return FetchAllLogEntries([]model.Project{project})
}

// Retrieve merged and sorted log entries for the given projects.
func FetchAllLogEntries(projects []model.Project) ([]model.LogEntry, error) {
	entries := make([]model.LogEntry, 0)
	// Concatenate all entries.
	for _, project := range projects {
		projectEntries, err := storage.LoadProjectLogEntries(project)
		if err != nil {
			return nil, err
		}
		entries = append(entries, projectEntries...)
	}
	return entries, nil
}

// Appends the given log entry to storage.
// Returns all entries for the day.
func AppendLogEntry(entry model.LogEntry, erasePreviousForDay bool) ([]model.LogEntry, error) {
	entries, err := FetchLogEntriesForProject(entry.Project)
	if err != nil {
		return nil, err
	}

	// Partition entries by whether they're for the given day.
	entriesForDay := make([]model.LogEntry, 0)
	otherEntries := make([]model.LogEntry, 0)
	for _, e := range entries {
		if util.IsSameDay(e.Date, entry.Date) {
			entriesForDay = append(entriesForDay, e)
		} else {
			otherEntries = append(otherEntries, e)
		}
	}

	// Validate the hour total.
	if entry.Hours+model.GetTotalHours(entriesForDay) > 24 {
		return nil, ErrHoursExceedsLimit
	}

	// If set, remove entries that match the date of the passed-in entry.
	if erasePreviousForDay {
		entries = otherEntries
	}

	// Just append, for now we don't care about order.
	entries = append(entries, entry)

	// Write all to storage.
	err = storage.SaveProjectLogEntries(entry.Project.ID(), entries)
	if err != nil {
		return nil, err
	}
	return model.FilterLogEntriesByDay(entries, entry.Date), nil
}

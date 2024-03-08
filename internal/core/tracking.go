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
	return RetrieveAllLogEntries([]model.Project{project})
}

// Retrieve merged and sorted log entries for the given projects.
func RetrieveAllLogEntries(projects []model.Project) ([]model.LogEntry, error) {
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
	entries, err := RetrieveLogEntries(entry.Project)
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
	if entry.Hours+GetTotalHours(entriesForDay) > 24 {
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
	return FilterLogEntriesByDay(entries, entry.Date), nil
}

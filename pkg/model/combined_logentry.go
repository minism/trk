package model

import (
	"time"
)

// Represents combined log entries for multiple projects on a given day.
type CombinedLogEntry struct {
	Projects []Project
	Date     time.Time
	Hours    float64
	Note     string
}

func CombineLogEntriesByProject(entries []LogEntry) []CombinedLogEntry {
	byDate := GroupLogEntriesByDate(entries)

	// Create combined log entries for each project
	combinedEntries := make([]CombinedLogEntry, 0)
	for el := byDate.Front(); el != nil; el = el.Next() {
		projects := make([]Project, 0)
		for _, entry := range el.Value {
			projects = append(projects, entry.Project)
		}

		// Create a combined log entry
		combinedEntry := CombinedLogEntry{
			Projects: projects,
			Date:     time.Unix(el.Key, 0),
			Hours:    GetTotalHours(el.Value),

			// TODO: Aggregate notes.
			Note: el.Value[0].Note,
		}

		combinedEntries = append(combinedEntries, combinedEntry)
	}

	return combinedEntries
}

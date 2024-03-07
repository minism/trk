package core

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
)

func FilterProjectsByIdFuzzy(projects []model.Project, query string) (model.Project, error) {
	matches := make([]model.Project, 0)
	for _, project := range projects {
		if strings.Contains(project.ID(), strings.ToLower(query)) {
			matches = append(matches, project)
		}
	}
	if len(matches) < 1 {
		return model.Project{}, fmt.Errorf("%w: %s", ErrProjectNotFound, query)
	} else if len(matches) > 1 {
		return model.Project{}, ErrMultipleProjectsMatched
	}
	return matches[0], nil
}

func FilterLogEntriesByDay(entries []model.LogEntry, date time.Time) []model.LogEntry {
	matches := make([]model.LogEntry, 0)
	for _, entry := range entries {
		if util.IsSameDay(entry.Date, date) {
			matches = append(matches, entry)
		}
	}
	return matches
}

func FilterLogEntriesSince(entries []model.LogEntry, since time.Time) []model.LogEntry {
	matches := make([]model.LogEntry, 0)
	for _, entry := range entries {
		if entry.Date.After(since) || entry.Date.Equal(since) {
			matches = append(matches, entry)
		}
	}
	return matches
}

func SortLogEntries(entries []model.LogEntry) {
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Date.Before(entries[j].Date)
	})
}

func GetTotalHours(entries []model.LogEntry) float64 {
	total := 0.0
	for _, entry := range entries {
		total += entry.Hours
	}
	return total
}

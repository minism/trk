package core

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
	"github.com/snabb/isoweek"
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

func GroupLogEntriesByProject(entries []model.LogEntry) map[string][]model.LogEntry {
	ret := make(map[string][]model.LogEntry)
	for _, entry := range entries {
		id := entry.Project.ID()
		if _, ok := ret[id]; !ok {
			ret[id] = make([]model.LogEntry, 0)
		}
		ret[id] = append(ret[id], entry)
	}
	return ret
}

func GroupLogEntriesByYearWeek(entries []model.LogEntry) *orderedmap.OrderedMap[time.Time, []model.LogEntry] {
	ret := orderedmap.NewOrderedMap[time.Time, []model.LogEntry]()
	for _, entry := range entries {
		year, week := entry.Date.ISOWeek()
		key := isoweek.StartTime(year, week, time.UTC)
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []model.LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

func GroupLogEntriesByBimonthly(entries []model.LogEntry) *orderedmap.OrderedMap[time.Time, []model.LogEntry] {
	ret := orderedmap.NewOrderedMap[time.Time, []model.LogEntry]()
	for _, entry := range entries {
		year, month, day := entry.Date.Date()
		if day > 15 {
			day = 16
		} else {
			day = 1
		}
		key := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []model.LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

func MergeAndSortLogEntries(entries []model.LogEntry) []model.LogEntry {
	// Sort by time.
	SortLogEntries(entries)

	// Merge entries for the same day and project.
	ret := make([]model.LogEntry, 0)
	last := model.LogEntry{Date: time.Unix(0, 0)}
	for _, entry := range entries {
		if !last.Date.Equal(entry.Date) || last.Project != entry.Project {
			if last.Hours > 0 {
				ret = append(ret, last)
			}
			last = entry
		} else {
			last.Hours += +entry.Hours
		}
	}
	ret = append(ret, last)

	return ret
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

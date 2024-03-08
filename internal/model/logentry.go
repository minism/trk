package model

import (
	"sort"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/internal/util"
	"github.com/snabb/isoweek"
)

// Represents hours logged for a particular project for a particular day.
type LogEntry struct {
	Project Project
	Date    time.Time
	Hours   float64
	Note    string
}

func FilterLogEntriesByDay(entries []LogEntry, date time.Time) []LogEntry {
	matches := make([]LogEntry, 0)
	for _, entry := range entries {
		if util.IsSameDay(entry.Date, date) {
			matches = append(matches, entry)
		}
	}
	return matches
}

func FilterLogEntriesSince(entries []LogEntry, since time.Time) []LogEntry {
	matches := make([]LogEntry, 0)
	for _, entry := range entries {
		if entry.Date.After(since) || entry.Date.Equal(since) {
			matches = append(matches, entry)
		}
	}
	return matches
}

func GroupLogEntriesByProject(entries []LogEntry) map[string][]LogEntry {
	ret := make(map[string][]LogEntry)
	for _, entry := range entries {
		id := entry.Project.ID()
		if _, ok := ret[id]; !ok {
			ret[id] = make([]LogEntry, 0)
		}
		ret[id] = append(ret[id], entry)
	}
	return ret
}

func GroupLogEntriesByYearWeek(entries []LogEntry) *orderedmap.OrderedMap[time.Time, []LogEntry] {
	ret := orderedmap.NewOrderedMap[time.Time, []LogEntry]()
	for _, entry := range entries {
		year, week := entry.Date.ISOWeek()
		key := isoweek.StartTime(year, week, time.UTC)
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

func GroupLogEntriesByBimonthly(entries []LogEntry) *orderedmap.OrderedMap[time.Time, []LogEntry] {
	ret := orderedmap.NewOrderedMap[time.Time, []LogEntry]()
	for _, entry := range entries {
		year, month, day := entry.Date.Date()
		if day > 15 {
			day = 16
		} else {
			day = 1
		}
		key := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

func MergeAndSortLogEntries(entries []LogEntry) []LogEntry {
	// Sort by time.
	SortLogEntries(entries)

	// Merge entries for the same day and project.
	ret := make([]LogEntry, 0)
	last := LogEntry{Date: time.Unix(0, 0)}
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

func SortLogEntries(entries []LogEntry) {
	sort.SliceStable(entries, func(i, j int) bool {
		return entries[i].Date.Before(entries[j].Date)
	})
}

func GetTotalHours(entries []LogEntry) float64 {
	total := 0.0
	for _, entry := range entries {
		total += entry.Hours
	}
	return total
}

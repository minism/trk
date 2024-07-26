package model

import (
	"sort"
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/internal/util"
)

// Represents hours logged for a particular project for a particular day.
type LogEntry struct {
	Project Project
	Date    time.Time
	Hours   float64
	Note    string
}

func MakeLogEntry(project Project, date time.Time, hours float64, note string) LogEntry {
	return LogEntry{
		Project: project,
		Date:    date,
		Hours:   hours,
		Note:    note,
	}
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

func FilterLogEntriesBetween(entries []LogEntry, from time.Time, to time.Time) []LogEntry {
	matches := make([]LogEntry, 0)
	for _, entry := range entries {
		if IsLogEntryInPeriod(entry, from, to) {
			matches = append(matches, entry)
		}
	}
	return matches
}

func ExcludeLogEntriesBetween(entries []LogEntry, from time.Time, to time.Time) []LogEntry {
	excluded := make([]LogEntry, 0)
	for _, entry := range entries {
		if !IsLogEntryInPeriod(entry, from, to) {
			excluded = append(excluded, entry)
		}
	}
	return excluded
}

func IsLogEntryInPeriod(entry LogEntry, from time.Time, to time.Time) bool {
	return (entry.Date.After(from) || entry.Date.Equal(from)) && entry.Date.Before(to)
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

// Keyed by Unix seconds.
func GroupLogEntriesByDate(entries []LogEntry) *orderedmap.OrderedMap[int64, []LogEntry] {
	ret := orderedmap.NewOrderedMap[int64, []LogEntry]()
	for _, entry := range entries {
		key := entry.Date.Unix()
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

// Keyed by Unix seconds.
func GroupLogEntriesByWeekStart(entries []LogEntry) *orderedmap.OrderedMap[int64, []LogEntry] {
	ret := orderedmap.NewOrderedMap[int64, []LogEntry]()
	for _, entry := range entries {
		key := util.GetStartOfWeek(entry.Date).Unix()
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

// Keyed by Unix seconds.
func GroupLogEntriesByBimonthly(entries []LogEntry) *orderedmap.OrderedMap[int64, []LogEntry] {
	ret := orderedmap.NewOrderedMap[int64, []LogEntry]()
	for _, entry := range entries {
		key := util.GetPrevBimonthlyDate(entry.Date).Unix()
		if val, ok := ret.Get(key); !ok {
			ret.Set(key, []LogEntry{entry})
		} else {
			ret.Set(key, append(val, entry))
		}
	}
	return ret
}

// Keyed by invoice ID.
// Note - entries are assumed to be sorted already.
func GroupLogEntriesByInvoicePeriods(entries []LogEntry, invoices []Invoice) *orderedmap.OrderedMap[int, []LogEntry] {
	ret := orderedmap.NewOrderedMap[int, []LogEntry]()
	if len(invoices) < 1 {
		return ret
	}

	// Arguable whether sorting both to get N+M is better than N*M to begin with, but it should be.
	// Could profile it if we really wanted to.
	SortLogEntries(entries)
	SortInvoices(invoices)
	i := 0
	for _, entry := range entries {
		for !IsLogEntryInPeriod(entry, invoices[i].StartDate, invoices[i].EndDate) {
			i++
			if i >= len(invoices) {
				return ret
			}
		}

		key := invoices[i].Id
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
	last := LogEntry{Date: util.MinDate}
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
	if last.Hours > 0 {
		ret = append(ret, last)
	}

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

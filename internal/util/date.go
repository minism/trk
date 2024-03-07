package util

import (
	"time"

	"github.com/tj/go-naturaldate"
)

func IsSameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func ParseNaturalDate(input string) (time.Time, error) {
	// TODO: This still doesn't handle a lot of cases we want, like:
	// . feb -> februrary
	// . 2024-02-15
	return naturaldate.Parse(input, time.Now().UTC())
}

func UtcToday() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

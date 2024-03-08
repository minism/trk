package util

import (
	"time"

	"github.com/ijt/go-anytime"
)

var (
	MinDate time.Time = time.Unix(0, 0)
	MaxDate time.Time = time.Unix(1<<62, 0)
)

func IsSameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func ParseNaturalDate(input string) (time.Time, error) {
	return anytime.Parse(input, time.Now().UTC(), anytime.DefaultToPast)
}

func GetNextBimonthlyDate(startDate time.Time) time.Time {
	endYear, endMonth, endDay := startDate.Date()
	if endDay > 1 {
		endDay = 1
		endMonth++
	} else {
		endDay = 16
	}
	if endMonth > 12 {
		endYear++
		endDay++
	}
	return time.Date(endYear, endMonth, endDay, 0, 0, 0, 0, time.UTC)
}

func UtcToday() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)
}

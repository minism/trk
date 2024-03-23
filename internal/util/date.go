package util

import (
	"time"

	"github.com/ijt/go-anytime"
	"github.com/minism/trk/internal/config"
	"github.com/snabb/isoweek"
)

var (
	MinDate time.Time = time.Unix(0, 0)
	MaxDate time.Time = time.Unix(1<<62, 0)

	FakeNowForTesting *time.Time
)

func IsSameDay(date1, date2 time.Time) bool {
	y1, m1, d1 := date1.Date()
	y2, m2, d2 := date2.Date()

	return y1 == y2 && m1 == m2 && d1 == d2
}

func ParseNaturalDate(input string) (time.Time, error) {
	return anytime.Parse(input, TrkNow().In(config.UserTimeZone), anytime.DefaultToPast)
}

func GetStartOfWeek(date time.Time) time.Time {
	year, week := date.ISOWeek()

	// Subtract 1 day since this library uses monday-based weeks.
	return isoweek.StartTime(year, week, config.UserTimeZone).Add(-time.Duration(24) * time.Hour)
}

func GetPrevBimonthlyDate(startDate time.Time) time.Time {
	year, month, day := startDate.Date()
	if day > 15 {
		day = 16
	} else {
		day = 1
	}
	return time.Date(year, month, day, 0, 0, 0, 0, config.UserTimeZone)
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
	return time.Date(endYear, endMonth, endDay, 0, 0, 0, 0, config.UserTimeZone)
}

func TrkNow() time.Time {
	if FakeNowForTesting != nil {
		return *FakeNowForTesting
	}
	return time.Now()
}

func TrkToday() time.Time {
	t := TrkNow()
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, config.UserTimeZone)
}

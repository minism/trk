package display

import (
	"time"

	"github.com/dustin/go-humanize"
)

func ReadableDate(date time.Time) string {
	return ColorDate(date.Format("Monday, January 2"))
}

func ReadableMoney(value float64) string {
	return ColorMoney("$%s", humanize.FormatFloat("#,###.##", value))
}

func ReadableHours(hours float64) string {
	return ColorHours("%.2f hours", hours)
}

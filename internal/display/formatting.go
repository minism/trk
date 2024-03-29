package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func ReadableDate(date time.Time) string {
	return ColorDate(ReadableDateWithoutColor(date))
}

func ReadableDateWithoutColor(date time.Time) string {
	return date.Format("Monday, January 2 2006")
}

func ReadableYearMonth(date time.Time) string {
	return ColorDate(date.Format("January 2006"))
}

func ReadableMoney(value float64) string {
	return ColorMoney("$%s", humanize.FormatFloat("#,###.##", value))
}

func ReadableHours(hours float64) string {
	return ColorHours("%s hours", FormatFloatMinDecimal(hours))
}

func ReadableEmpty() string {
	return ColorNull("(empty)")
}

func FormatFloatMinDecimal(value float64) string {
	return strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", value), "0"), ".")
}

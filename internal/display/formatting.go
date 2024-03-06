package display

import "time"

func ReadableDate(date time.Time) string {
	return ColorDate(date.Format("Monday, January 2"))
}

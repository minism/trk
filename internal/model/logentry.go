package model

import (
	"time"
)

// Represents hours logged for a particular project for a particular day.
type LogEntry struct {
	Project Project
	Date    time.Time
	Hours   float64
	Note    string
}

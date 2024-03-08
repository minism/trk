package model

import "time"

type Invoice struct {
	Project     Project
	StartDate   time.Time
	EndDate     time.Time // Exclusive
	HoursLogged float64
	HoursBilled float64
	HourlyRate  float64
	IsSent      bool
	IsPaid      bool
}

func (i *Invoice) Total() float64 {
	return i.HoursBilled * i.HourlyRate
}

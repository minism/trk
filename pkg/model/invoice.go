package model

import "time"

type Invoice struct {
	Project             Project
	StartDate           time.Time
	EndDate             time.Time // Exclusive
	HoursLogged         float64
	OverrideHoursBilled float64
	OverrideHourlyRate  float64
	IsSent              bool
	IsPaid              bool
}

func (i *Invoice) HoursBilled() float64 {
	if i.OverrideHoursBilled > 0 {
		return i.OverrideHoursBilled
	}
	return i.HoursLogged
}

func (i *Invoice) HourlyRate() float64 {
	if i.OverrideHourlyRate > 0 {
		return i.OverrideHourlyRate
	}
	return i.Project.HourlyRate
}

func (i *Invoice) Total() float64 {
	return i.HoursBilled() * i.HourlyRate()
}

func MaybeFindInvoiceByStartDate(invoices []Invoice, date time.Time) *Invoice {
	for _, invoice := range invoices {
		if invoice.StartDate.Equal(date) {
			return &invoice
		}
	}
	return nil
}

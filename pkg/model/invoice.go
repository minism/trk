package model

import (
	"time"
)

type Invoice struct {
	Id          int       `yaml:"id"`
	StartDate   time.Time `yaml:"start_date"`
	EndDate     time.Time `yaml:"end_date"` // Exclusive
	HoursLogged float64   `yaml:"hours_logged"`
	HoursBilled float64   `yaml:"hours_billed"`
	HourlyRate  float64   `yaml:"hourly_rate"`
	IsSent      bool      `yaml:"is_sent"`
	IsPaid      bool      `yaml:"is_paid"`
}

func (i *Invoice) Status() string {
	if i.IsPaid {
		return "Paid"
	} else if i.IsSent {
		return "Sent"
	} else {
		return ""
	}
}

func (i *Invoice) Total() float64 {
	return i.HoursBilled * i.HourlyRate
}

func MaybeFindInvoiceByStartDate(invoices []Invoice, date time.Time) *Invoice {
	for _, invoice := range invoices {
		if invoice.StartDate.Equal(date) {
			return &invoice
		}
	}
	return nil
}

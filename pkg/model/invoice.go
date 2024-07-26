package model

import (
	"sort"
	"time"
)

type Invoice struct {
	// Fields stored to disk.
	Id          int       `yaml:"id"`
	StartDate   time.Time `yaml:"start_date"`
	EndDate     time.Time `yaml:"end_date"` // Exclusive
	HoursBilled float64   `yaml:"hours_billed"`
	HourlyRate  float64   `yaml:"hourly_rate"`
	IsSent      bool      `yaml:"is_sent"`
	IsPaid      bool      `yaml:"is_paid"`

	// Fields calculated.
	HoursLogged float64
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

func SortInvoices(invoices []Invoice) {
	sort.SliceStable(invoices, func(i, j int) bool {
		return invoices[i].StartDate.Before(invoices[j].StartDate)
	})
}

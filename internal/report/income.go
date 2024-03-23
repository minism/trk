package report

import (
	"time"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/pkg/model"
)

type MonthIncomeReport struct {
	YearMonth     time.Time
	PaidAmount    float64
	PendingAmount float64
}

func GenerateMonthlyIncomeForYear(year int) ([]MonthIncomeReport, error) {
	invoices, err := core.FetchAllProjectInvoices()
	if err != nil {
		return nil, err
	}

	var ret []MonthIncomeReport

	// End at the lastest month we have an invoice for.
	maxMonth := getMaxMonth(invoices)
	for month := time.January; month <= maxMonth; month++ {
		yearMonth := time.Date(year, month, 1, 0, 0, 0, 0, config.UserTimeZone)
		var paidAmount, pendingAmount float64

		for _, invoice := range invoices {
			if invoice.StartDate.Year() == year && invoice.StartDate.Month() == month {
				if invoice.IsPaid {
					paidAmount += invoice.Total()
				} else {
					pendingAmount += invoice.Total()
				}
			}
		}

		ret = append(ret, MonthIncomeReport{
			YearMonth:     yearMonth,
			PaidAmount:    paidAmount,
			PendingAmount: pendingAmount,
		})
	}

	return ret, nil
}

func getMaxMonth(invoices []model.ProjectInvoice) time.Month {
	maxMonth := time.January

	for _, invoice := range invoices {
		month := invoice.StartDate.Month()
		if month > maxMonth {
			maxMonth = month
		}
	}

	return maxMonth
}

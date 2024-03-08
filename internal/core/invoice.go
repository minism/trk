package core

import (
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/util"
)

func GenerateInvoicesForProject(project model.Project) ([]model.Invoice, error) {
	entries, err := RetrieveLogEntries(project)
	if err != nil {
		return nil, err
	}

	// Assume bimonthly now but we need to support other periods.
	byStartDate := model.GroupLogEntriesByBimonthly(entries)
	invoices := make([]model.Invoice, 0)
	for el := byStartDate.Front(); el != nil; el = el.Next() {
		startDate := el.Key
		endDate := util.GetNextBimonthlyDate(startDate)
		totalHours := model.GetTotalHours(el.Value)

		invoices = append(invoices, model.Invoice{
			Project:     project,
			StartDate:   startDate,
			EndDate:     endDate,
			HoursLogged: totalHours,
		})
	}
	return invoices, nil
}

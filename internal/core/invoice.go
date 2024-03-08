package core

import (
	"github.com/minism/trk/internal/model"
	"github.com/minism/trk/internal/storage"
	"github.com/minism/trk/internal/util"
)

func FetchInvoicesForProject(project model.Project) ([]model.Invoice, error) {
	return storage.LoadInvoices(project)
}

func GenerateInvoicesForProject(project model.Project) ([]model.Invoice, error) {
	// Load entries and group by bimonthly.
	// TODO: Support other invoice periods.
	entries, err := RetrieveLogEntries(project)
	if err != nil {
		return nil, err
	}
	entriesByStartDate := model.GroupLogEntriesByBimonthly(entries)

	// Fetch allInvoices for the project and see what's missing.
	allInvoices, err := FetchInvoicesForProject(project)
	if err != nil {
		return nil, err
	}

	// Drop any we've already generated.
	for _, invoice := range allInvoices {
		entriesByStartDate.Delete(invoice.StartDate)
	}

	// Create new invoices for the remainder.
	newInvoices := make([]model.Invoice, 0)
	for el := entriesByStartDate.Front(); el != nil; el = el.Next() {
		startDate := el.Key
		endDate := util.GetNextBimonthlyDate(startDate)
		totalHours := model.GetTotalHours(el.Value)

		invoice := model.Invoice{
			Project:     project,
			StartDate:   startDate,
			EndDate:     endDate,
			HoursLogged: totalHours,
		}
		newInvoices = append(newInvoices, invoice)
		allInvoices = append(allInvoices, invoice)
	}

	// Write back to storage.
	err = storage.SaveInvoices(project, allInvoices)
	if err != nil {
		return nil, err
	}

	return newInvoices, nil
}

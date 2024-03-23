package core

import (
	"time"

	"github.com/minism/trk/internal/storage"
	"github.com/minism/trk/internal/util"
	"github.com/minism/trk/pkg/model"
)

func FetchInvoicesForProject(project model.Project) ([]model.ProjectInvoice, error) {
	invoices, err := storage.LoadInvoices(project)
	if err != nil {
		return nil, err
	}
	return model.MakeProjectInvoices(project, invoices), nil
}

func GenerateNewInvoicesForProject(project model.Project) ([]model.ProjectInvoice, error) {
	// Load entries and group by bimonthly.
	entries, err := RetrieveLogEntries(project)
	if err != nil {
		return nil, err
	}

	// We only want to consider entries for finished invoice periods.
	endDate := util.GetPrevBimonthlyDate(util.TrkToday())
	entries = model.FilterLogEntriesBetween(entries, util.MinDate, endDate)

	// Group invoices bimonthly.
	// TODO: Support other invoice periods.
	entriesByStartDate := model.GroupLogEntriesByBimonthly(entries)

	// Fetch all invoices for the project and see what's missing.
	allInvoices, err := storage.LoadInvoices(project)
	if err != nil {
		return nil, err
	}

	// Drop any we've already generated.
	maxInvoiceId := 0
	for _, invoice := range allInvoices {
		entriesByStartDate.Delete(invoice.StartDate.Unix())

		// Get the highest invoice ID number to generate the next one.
		if invoice.Id > maxInvoiceId {
			maxInvoiceId = invoice.Id
		}
	}

	// Create new invoices for the remainder.
	newInvoices := make([]model.Invoice, 0)
	for el := entriesByStartDate.Front(); el != nil; el = el.Next() {
		startDate := time.Unix(el.Key, 0)
		endDate := util.GetNextBimonthlyDate(startDate)
		totalHours := model.GetTotalHours(el.Value)

		invoice := model.Invoice{
			Id:          maxInvoiceId + 1,
			StartDate:   startDate,
			EndDate:     endDate,
			HoursLogged: totalHours,

			// Billed defaults to logged but can be overridden.
			HoursBilled: totalHours,

			// Rate defaults to project rate but can be overridden.
			HourlyRate: project.HourlyRate,
		}
		newInvoices = append(newInvoices, invoice)
		allInvoices = append(allInvoices, invoice)
		maxInvoiceId++
	}

	// Write back to storage.
	err = storage.SaveInvoices(project, allInvoices)
	if err != nil {
		return nil, err
	}

	return model.MakeProjectInvoices(project, newInvoices), nil
}

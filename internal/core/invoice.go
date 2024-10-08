package core

import (
	"fmt"
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

	// Calculate the hours logged from log entry data.
	entries, err := FetchLogEntriesForProject(project)
	if err != nil {
		return nil, err
	}
	entriesByInvoiceId := model.GroupLogEntriesByInvoicePeriods(entries, invoices)
	for i, invoice := range invoices {
		if entries, ok := entriesByInvoiceId.Get(invoice.Id); !ok {
			invoices[i].HoursLogged = 0
		} else {
			invoices[i].HoursLogged = model.GetTotalHours(entries)
		}
	}

	return model.MakeProjectInvoices(project, invoices), nil
}

func FetchAllProjectInvoices() ([]model.ProjectInvoice, error) {
	projects, err := FetchAllProjects()
	if err != nil {
		return nil, err
	}

	ret := make([]model.ProjectInvoice, 0)
	for _, project := range projects {
		invoices, err := FetchInvoicesForProject(project)
		if err != nil {
			return nil, err
		}
		ret = append(ret, invoices...)
	}

	return ret, nil
}

func FetchProjectInvoiceById(id model.ProjectInvoiceId) (model.ProjectInvoice, error) {
	projectId, invoiceId, err := model.ParseProjectInvoiceId(id)
	if err != nil {
		return model.ProjectInvoice{}, err
	}

	project, err := FetchProjectById(projectId)
	if err != nil {
		return model.ProjectInvoice{}, err
	}

	invoices, err := FetchInvoicesForProject(project)
	if err != nil {
		return model.ProjectInvoice{}, err
	}

	for _, pi := range invoices {
		if pi.Invoice.Id == invoiceId {
			return pi, nil
		}
	}

	return model.ProjectInvoice{}, fmt.Errorf("%w: %s", ErrInvoiceNotFound, id)
}

func DeleteProjectInvoiceById(id model.ProjectInvoiceId) error {
	pi, err := FetchProjectInvoiceById(id)
	if err != nil {
		return err
	}

	// Rewrite updatedInvoices.
	updatedInvoices, err := storage.LoadInvoices(pi.Project)
	if err != nil {
		return err
	}
	for i, invoice := range updatedInvoices {
		if pi.Invoice.Id == invoice.Id {
			updatedInvoices = append(updatedInvoices[:i], updatedInvoices[i+1:]...)
			break
		}
	}
	return storage.SaveInvoices(pi.Project, updatedInvoices)
}

func UpdateProjectInvoice(updatedInvoice model.ProjectInvoice) error {
	invoices, err := storage.LoadInvoices(updatedInvoice.Project)
	if err != nil {
		return err
	}

	// Update the invoice in the list of all invoices.
	updated := false
	for i, invoice := range invoices {
		if invoice.Id == updatedInvoice.Invoice.Id {
			invoices[i] = updatedInvoice.Invoice
			updated = true
			break
		}
	}
	if !updated {
		return fmt.Errorf("%w: %s", ErrInvoiceNotFound, updatedInvoice.Id())
	}

	// Save it back to storage.
	err = storage.SaveInvoices(updatedInvoice.Project, invoices)
	if err != nil {
		return err
	}

	return nil
}

func GenerateNewInvoicesForProject(project model.Project, upToDate time.Time) ([]model.ProjectInvoice, error) {
	// Load entries and group by bimonthly.
	entries, err := FetchLogEntriesForProject(project)
	if err != nil {
		return nil, err
	}

	// We only want to consider entries for finished invoice periods.
	endDate := util.GetPrevBimonthlyDate(upToDate)
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

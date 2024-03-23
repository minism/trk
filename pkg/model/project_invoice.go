package model

import "fmt"

// Simple composite structure that referse to an invoice for a particular project.
type ProjectInvoice struct {
	Project Project
	Invoice Invoice
}

func MakeProjectInvoices(project Project, invoices []Invoice) []ProjectInvoice {
	ret := make([]ProjectInvoice, 0)
	for _, invoice := range invoices {
		ret = append(ret, ProjectInvoice{
			Project: project,
			Invoice: invoice,
		})
	}
	return ret
}

func (pi *ProjectInvoice) GlobalId() string {
	return fmt.Sprintf("%s-%d", pi.Project.ID(), pi.Invoice.Id)
}

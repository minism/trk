package model

import (
	"fmt"
	"strconv"

	"github.com/minism/trk/internal/util"
)

// Simple composite structure that referse to an invoice for a particular project.
type ProjectInvoice struct {
	Invoice
	Project Project
}

type ProjectInvoiceId string

func (pi *ProjectInvoice) Id() ProjectInvoiceId {
	return ProjectInvoiceId(fmt.Sprintf("%s-%d", pi.Project.ID(), pi.Invoice.Id))
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

func ParseProjectInvoiceId(id ProjectInvoiceId) (string, int, error) {
	projectId, invoiceIdStr := util.SplitStringUpToLastHyphen(string(id))
	invoiceId, err := strconv.Atoi(invoiceIdStr)
	if err != nil {
		return "", 0, err
	}
	return projectId, invoiceId, nil
}

func FilterProjectInvoicesByUnpaid(invoices []ProjectInvoice) []ProjectInvoice {
	var ret []ProjectInvoice
	for _, invoice := range invoices {
		if !invoice.IsPaid {
			ret = append(ret, invoice)
		}
	}
	return ret
}

func SumInvoices(invoices []ProjectInvoice) float64 {
	total := 0.0
	for _, invoice := range invoices {
		total += invoice.Total()
	}
	return total
}

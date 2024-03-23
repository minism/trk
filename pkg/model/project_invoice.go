package model

import (
	"fmt"
	"strconv"

	"github.com/minism/trk/internal/util"
)

// Simple composite structure that referse to an invoice for a particular project.
type ProjectInvoice struct {
	Project Project
	Invoice Invoice
}

type ProjectInvoiceId string

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

func (pi *ProjectInvoice) Id() ProjectInvoiceId {
	return ProjectInvoiceId(fmt.Sprintf("%s-%d", pi.Project.ID(), pi.Invoice.Id))
}

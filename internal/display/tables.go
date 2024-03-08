package display

import (
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/pkg/model"
	"github.com/rodaine/table"
)

func init() {
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return ColorTableHeader(format, vals...)
	}
	table.DefaultPadding = 12
}

func PrintLogEntryTable(entries []model.LogEntry) {
	tbl := table.New("Date", "Project", "Hours", "Note")
	tbl.WithFirstColumnFormatter(ColorDate)
	for _, entry := range entries {
		tbl.AddRow(
			entry.Date.Format("Mon 1/2"), ColorProject(entry.Project.Name), entry.Hours, entry.Note)
	}
	tbl.Print()
}

func PrintWeeklyLogEntryTable(byWeek *orderedmap.OrderedMap[time.Time, []model.LogEntry]) {
	tbl := table.New("Week", "Total Hours")
	tbl.WithFirstColumnFormatter(ColorDate)
	for el := byWeek.Front(); el != nil; el = el.Next() {
		total := model.GetTotalHours(el.Value)
		tbl.AddRow(el.Key.Format("Week of 1/2/2006"), total)
	}
	tbl.Print()
}

func PrintInvoicePeriodLogEntryTable(byInvoiceDate *orderedmap.OrderedMap[time.Time, []model.LogEntry], invoices []model.Invoice) {
	tbl := table.New("Start Date", "Total Hours", "Associated invoice")

	for el := byInvoiceDate.Front(); el != nil; el = el.Next() {
		total := model.GetTotalHours(el.Value)
		invoice := model.MaybeFindInvoiceByStartDate(invoices, el.Key)
		tbl.AddRow(el.Key, total, invoice)
	}
	tbl.Print()
}

func PrintInvoicesTable(invoices []model.Invoice) {
	tbl := table.New("Invoice Date", "Hours Billed", "Rate", "Total", "Sent?", "Paid?")
	tbl.WithFirstColumnFormatter(ColorDate)
	for _, invoice := range invoices {
		tbl.AddRow(invoice.StartDate, invoice.HoursBilled(), invoice.HourlyRate(), ReadableMoney(invoice.Total()), invoice.IsSent, invoice.IsPaid)
	}
	tbl.Print()
}

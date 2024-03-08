package display

import (
	"fmt"
	"time"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/pkg/model"
	"github.com/rodaine/table"
)

func init() {
	table.DefaultPadding = 8
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return ColorTableHeader(format, vals...)
	}
	table.DefaultWidthFunc = func(s string) int {
		return utf8.RuneCountInString(stripansi.Strip(s))
	}
}

func PrintLogEntryTable(entries []model.LogEntry) {
	tbl := table.New("Date", "Project", "Hours", "Note")
	tbl.WithFirstColumnFormatter(ColorDate)
	for _, entry := range entries {
		tbl.AddRow(
			entry.Date.Format("Mon 1/2"), ColorProject(entry.Project.ID()), entry.Hours, entry.Note)
	}
	tbl.Print()
}

func PrintWeeklyLogEntryTable(byWeek *orderedmap.OrderedMap[time.Time, []model.LogEntry]) {
	tbl := table.New("Week", "Total Hours")
	for el := byWeek.Front(); el != nil; el = el.Next() {
		total := model.GetTotalHours(el.Value)
		tbl.AddRow(fmt.Sprintf("Week of %s", ColorDate(el.Key.Format("1/2/2006"))), total)
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
	tbl := table.New("Invoice Date", "Hours Billed", "Rate", "Total", "Status")
	tbl.WithFirstColumnFormatter(ColorDate)
	for _, invoice := range invoices {
		tbl.AddRow(invoice.StartDate.Format("2006-01-02"), invoice.HoursBilled(), invoice.HourlyRate(), ReadableMoney(invoice.Total()), invoice.Status())
	}
	tbl.Print()
}

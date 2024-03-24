package display

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/acarl005/stripansi"
	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/internal/report"
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

func PrintCombinedLogEntryTable(entries []model.CombinedLogEntry) {
	tbl := table.New("Date", "Projects", "Hours", "Notes")
	tbl.WithFirstColumnFormatter(ColorDate)
	for _, entry := range entries {
		projectIds := make([]string, 0)
		for _, project := range entry.Projects {
			projectIds = append(projectIds, project.ID())
		}
		tbl.AddRow(
			entry.Date.Format("Mon 1/2"), ColorProject(strings.Join(projectIds, ", ")), entry.Hours, entry.Note)
	}
	tbl.Print()
}

func PrintWeeklyLogEntryTable(byWeek *orderedmap.OrderedMap[int64, []model.LogEntry]) {
	tbl := table.New("Week", "Total Hours")
	for el := byWeek.Front(); el != nil; el = el.Next() {
		total := model.GetTotalHours(el.Value)
		weekStart := time.Unix(el.Key, 0)
		tbl.AddRow(fmt.Sprintf("Week of %s", ColorDate(weekStart.Format("1/2/2006"))), total)
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

func PrintProjectInvoicesTable(invoices []model.ProjectInvoice) {
	tbl := table.New("Invoice ID", "Invoice Date", "Hours Billed", "Rate", "Total", "Status")
	for _, invoice := range invoices {
		hours := FormatFloatMinDecimal(invoice.HoursBilled)
		rate := FormatFloatMinDecimal(invoice.HourlyRate)

		// If hours billed differed from logs, show slightly more information.
		if invoice.HoursBilled != invoice.HoursLogged {
			hours = fmt.Sprintf("%s (%s logged)", hours, FormatFloatMinDecimal(invoice.HoursLogged))
			rate = fmt.Sprintf("%s (%.f)", rate, invoice.HoursBilled/invoice.HoursLogged*invoice.HourlyRate)
		}

		tbl.AddRow(
			ColorIdentifier(string(invoice.Id())),
			ColorDate(invoice.StartDate.Format("2006-01-02")),
			hours,
			rate,
			ReadableMoney(invoice.Total()),
			invoice.Status())
	}
	tbl.Print()
}

func PrintIncomeReportTable(reports []report.MonthIncomeReport) {
	tbl := table.New("Month", "Paid amount", "Pending amount")
	totalPaid := 0.0
	totalPending := 0.0

	for _, r := range reports {
		tbl.AddRow(ReadableYearMonth(r.YearMonth), ReadableMoney(r.PaidAmount), ReadableMoney(r.PendingAmount))
		totalPaid += r.PaidAmount
		totalPending += r.PendingAmount
	}

	// TODO: Ideally we could support separators within a table.
	tbl.AddRow(ColorMoney("TOTAL"), ReadableMoney(totalPaid), ReadableMoney(totalPending))
	tbl.Print()
}

package display

import (
	"time"

	"github.com/elliotchance/orderedmap/v2"
	"github.com/minism/trk/internal/core"
	"github.com/minism/trk/internal/model"
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
		total := core.GetTotalHours(el.Value)
		tbl.AddRow(el.Key.Format("Week of 1/2/2006"), total)
	}
	tbl.Print()
}

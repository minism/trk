package display

import (
	"fmt"

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
	tbl := table.New("Project", "Date", "Hours", "Note")
	tbl.WithFirstColumnFormatter(ColorProject)
	for _, entry := range entries {
		tbl.AddRow(
			entry.Project.Name, (entry.Date.Format("2006-01-02")), entry.Hours, entry.Note)
	}
	tbl.Print()
	fmt.Println()
}

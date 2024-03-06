package display

import (
	"github.com/rodaine/table"
)

func init() {
	table.DefaultHeaderFormatter = func(format string, vals ...interface{}) string {
		return ColorTableHeader(format, vals...)
	}
	table.DefaultPadding = 12
}

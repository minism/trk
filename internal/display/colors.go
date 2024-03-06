package display

import (
	"github.com/fatih/color"
)

var (
	ColorSuccess     = color.New(color.FgGreen).SprintfFunc()
	ColorProject     = color.New(color.FgGreen).SprintfFunc()
	ColorTableHeader = color.New(color.FgCyan, color.Underline).SprintfFunc()
	ColorError       = color.New(color.FgRed).SprintfFunc()
)

package display

import (
	"github.com/fatih/color"
)

var (
	ColorSuccess     = color.New(color.FgGreen).SprintfFunc()
	ColorError       = color.New(color.FgRed).SprintfFunc()
	ColorDate        = color.New(color.FgCyan).SprintfFunc()
	ColorMoney       = color.New(color.FgGreen).SprintfFunc()
	ColorHours       = color.New(color.FgGreen).SprintfFunc()
	ColorProject     = color.New(color.FgYellow).SprintfFunc()
	ColorTableHeader = color.New(color.FgHiBlack, color.Underline).SprintfFunc()
	ColorNull        = color.New(color.FgHiBlack).SprintfFunc()
)

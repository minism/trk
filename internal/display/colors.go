package display

import (
	"github.com/fatih/color"
)

var (
	ColorSuccess     = color.New(color.FgGreen).SprintfFunc()
	ColorError       = color.New(color.FgRed).SprintfFunc()
	ColorDate        = color.New(color.FgCyan).SprintfFunc()
	ColorProject     = color.New(color.FgYellow).SprintfFunc()
	ColorTableHeader = color.New(color.FgCyan, color.Underline).SprintfFunc()
)

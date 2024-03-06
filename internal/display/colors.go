package display

import (
	"github.com/fatih/color"
)

var (
	ColorSuccess = color.New(color.FgGreen).SprintfFunc()
	ColorError   = color.New(color.FgRed).SprintfFunc()
)

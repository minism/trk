package tests

import (
	"os"
	"testing"

	"github.com/minism/trk/cmd"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"trk": func() int {
			cmd.Execute()
			return 0
		},
	}))
}

// For more info see https://bitfieldconsulting.com/golang/cli-testing
func TestIntegration(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "scripts",
	})
}

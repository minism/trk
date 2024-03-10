package tests

import (
	"os"
	"testing"
	"time"

	"github.com/minism/trk/cmd"
	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/util"
	"github.com/rogpeppe/go-internal/testscript"
)

func TestMain(m *testing.M) {
	// Setup a fake "today" for testing.
	// Wednesday, March 6th, 2024.
	// We use wednesday to test mid-week behavior.
	t := time.Date(2024, 3, 6, 0, 0, 0, 0, config.UserTimeZone)
	util.FakeNowForTesting = &t

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

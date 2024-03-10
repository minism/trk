package core

import (
	"os"
	"time"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/storage"
)

// Tries to initialize trk for the current user.
// Returns true if successful, or false if it already exists.
// Returns an error if trk couldn't be initialized for some reason.
func InitTrk(forceReset bool) (bool, error) {
	initialized := false

	// Ensure the app+logs directory exists.
	worklogDir := config.GetWorkLogDir()
	_, err := os.Stat(worklogDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(worklogDir, os.ModePerm)
		if err != nil {
			return initialized, err
		}
	} else if err != nil {
		return initialized, err
	}

	// Ensure the config exists.
	configPath := config.GetConfigPath()
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) || forceReset {
		err = writeDefaultConfig()
		if err != nil {
			return initialized, err
		}
		initialized = true
	} else if err != nil {
		return initialized, err
	}

	return initialized, nil
}

func writeDefaultConfig() error {
	loc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		return err
	}

	return storage.SaveConfig(config.TrkConfig{
		TimeZone: loc,
		Projects: []config.ProjectConfig{
			{
				Name:       "Example project",
				HourlyRate: 50,
			},
		},
	})
}

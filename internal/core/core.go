package core

import (
	"os"

	"github.com/minism/trk/internal/config"
	"github.com/minism/trk/internal/storage"
)

// Tries to initialize trk for the current user.
// Returns true if successful, or false if it already exists.
// Returns an error if trk couldn't be initialized for some reason.
func InitTrk() (bool, error) {
	initialized := false

	// Ensure the app directory exists.
	appDir := config.GetUserAppDir()
	_, err := os.Stat(appDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(appDir, os.ModePerm)
		if err != nil {
			return initialized, err
		}
	} else if err != nil {
		return initialized, err
	}

	// Ensure the config exists.
	configPath := config.GetConfigPath()
	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
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
	// TODO: When should we use pointers for these?
	return storage.SaveConfig(config.TrkConfig{})
}

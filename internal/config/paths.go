package config

import (
	"os"
	"path"
)

func GetConfigPath() string {
	return path.Join(GetUserAppDir(), "config.yaml")
}

// The directory all trk's application data is located for the current user.
func GetUserAppDir() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return path.Join(home, ".trk")
}

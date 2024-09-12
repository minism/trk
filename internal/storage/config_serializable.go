package storage

import (
	"time"

	"github.com/minism/trk/internal/config"
)

type trkConfigSerializable struct {
	AutoCommit bool                   `yaml:"auto_commit"`
	TimeZone   string                 `yaml:"time_zone"`
	Projects   []config.ProjectConfig `yaml:"projects"`
}

// TODO: Can we just use a custom yaml marshal function rather than these
// transformations?
func configToSerializable(config config.TrkConfig) (trkConfigSerializable, error) {
	return trkConfigSerializable{
		AutoCommit: config.AutoCommit,
		TimeZone:   config.TimeZone.String(),
		Projects:   config.Projects,
	}, nil
}

func configFromSerializable(sConfig trkConfigSerializable) (config.TrkConfig, error) {
	loc, err := time.LoadLocation(sConfig.TimeZone)
	if err != nil {
		return config.TrkConfig{}, err
	}
	return config.TrkConfig{
		AutoCommit: sConfig.AutoCommit,
		TimeZone:   loc,
		Projects:   sConfig.Projects,
	}, nil
}
